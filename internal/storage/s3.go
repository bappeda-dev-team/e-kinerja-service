package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	presignedS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

// Prefix folder per jenis file — ubah sesuai kebutuhan
const (
	FolderPemda      = "pemda/logo"
	FolderAplikasi   = "aplikasi/logo"
	FolderProfilePic = "users/profile"
	FolderLampiran   = "permintaan/lampiran"
)

// allowedMIME maps each folder to the set of permitted content types.
var allowedMIME = map[string]map[string]bool{
	FolderPemda:      {"image/jpeg": true, "image/png": true, "image/webp": true},
	FolderAplikasi:   {"image/jpeg": true, "image/png": true, "image/webp": true},
	FolderProfilePic: {"image/jpeg": true, "image/png": true, "image/webp": true},
	FolderLampiran:   {"image/jpeg": true, "image/png": true, "image/webp": true, "application/pdf": true},
}

func detectAndValidateMIME(file multipart.File, folder string) (string, error) {
	buf := make([]byte, 512)
	n, err := io.ReadFull(file, buf)
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", fmt.Errorf("gagal membaca file: %w", err)
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("gagal reset file: %w", err)
	}

	mime := http.DetectContentType(buf[:n])
	// DetectContentType may append "; charset=..." — strip it.
	mime = strings.SplitN(mime, ";", 2)[0]

	allowed, ok := allowedMIME[folder]
	if !ok || !allowed[mime] {
		return "", fmt.Errorf("tipe file tidak diizinkan: %s", mime)
	}
	return mime, nil
}

var (
	s3Client      *s3.Client
	presignClient *presignedS3.PresignClient
	bucketName    string
)

func Init() error {
	bucketName = os.Getenv("S3_BUCKET")
	region := os.Getenv("S3_REGION")
	endpoint := os.Getenv("S3_ENDPOINT")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return fmt.Errorf("gagal load S3 config: %w", err)
	}

	opts := []func(*s3.Options){}
	if endpoint != "" {
		opts = append(opts, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = true
		})
	}

	s3Client = s3.NewFromConfig(cfg, opts...)
	presignClient = s3.NewPresignClient(s3Client)
	return nil
}

// UploadFile mengupload file ke S3 dan mengembalikan public URL.
// Gunakan konstanta Folder* sebagai argumen folder.
func UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	mime, err := detectAndValidateMIME(file, folder)
	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	key := buildKey(folder, ext)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(mime),
	})
	if err != nil {
		return "", fmt.Errorf("gagal upload ke S3: %w", err)
	}

	return publicURL(key), nil
}

// PresignUpload membuat pre-signed URL untuk upload langsung dari client (PUT).
// expires: durasi validitas URL, misal 15 * time.Minute.
func PresignUpload(folder, ext string, expires time.Duration) (presignURL string, key string, err error) {
	key = buildKey(folder, ext)

	req, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))
	if err != nil {
		return "", "", fmt.Errorf("gagal buat presign upload URL: %w", err)
	}

	return req.URL, key, nil
}

// PresignDownload membuat pre-signed URL untuk download file private (GET).
// Gunakan ini jika bucket tidak public.
func PresignDownload(key string, expires time.Duration) (string, error) {
	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))
	if err != nil {
		return "", fmt.Errorf("gagal buat presign download URL: %w", err)
	}

	return req.URL, nil
}

// DeleteFile menghapus file dari S3 berdasarkan key-nya (bukan full URL).
func DeleteFile(key string) error {
	if key == "" {
		return nil
	}

	_, err := s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	return err
}

// KeyFromURL mengekstrak object key dari public URL.
// Berguna sebelum memanggil DeleteFile.
func KeyFromURL(fileURL string) string {
	if fileURL == "" {
		return ""
	}

	endpoint := os.Getenv("S3_ENDPOINT")
	if endpoint != "" {
		prefix := fmt.Sprintf("%s/%s/", strings.TrimRight(endpoint, "/"), bucketName)
		return strings.TrimPrefix(fileURL, prefix)
	}

	region := os.Getenv("S3_REGION")
	prefix := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", bucketName, region)
	return strings.TrimPrefix(fileURL, prefix)
}

// --- internal helpers ---

func buildKey(folder, ext string) string {
	return fmt.Sprintf("%s/%s_%d%s",
		strings.Trim(folder, "/"),
		uuid.NewString(),
		time.Now().UnixMilli(),
		ext,
	)
}

// PublicURL mengkonversi object key menjadi public URL bucket.
func PublicURL(key string) string {
	return publicURL(key)
}

func publicURL(key string) string {
	endpoint := os.Getenv("S3_ENDPOINT")
	if endpoint != "" {
		return fmt.Sprintf("%s/%s/%s", strings.TrimRight(endpoint, "/"), bucketName, key)
	}
	region := os.Getenv("S3_REGION")
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, key)
}
