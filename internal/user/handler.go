package user

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/internal/storage"
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var req LoginRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.ErrorResponse(400, "Validasi gagal",
				helpers.FormatValidationError(err)))
	}

	token, err := LoginService(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized,
			helpers.ErrorResponse(401, err.Error(), nil))
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "login berhasil", map[string]string{
			"token": token,
		}))
}

func GetAllUser(c echo.Context) error {
	result, err := GetUserServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func GetUserID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetUserServicesID(id)

	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func Create(c echo.Context) error {
	var req RegisterRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	user, err := CreateUserService(req)
	if err != nil {

		if err.Error() == "role tidak ditemukan" ||
			err.Error() == "username sudah digunakan" {
			return exception.BadRequest(err.Error())
		}

		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Registrasi berhasil", user))
}

func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Logout berhasil", nil))
}

// PresignProfilePicUpload godoc
// @Summary Buat pre-signed URL untuk upload foto profil user
// @Description Menghasilkan pre-signed PUT URL ke S3. Client upload file langsung ke URL tersebut, lalu panggil PATCH /:id/profile-picture dengan key yang dikembalikan.
// @Tags Users
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Param ext query string false "Ekstensi file, misal .png atau .jpg (default: .jpg)"
// @Success 200 {object} helpers.APIResponse{data=map[string]string}
// @Failure 400 {object} helpers.APIResponse
// @Router /users/{id}/profile-picture/presign [post]
func PresignProfilePicUpload(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	ext := c.QueryParam("ext")
	if ext == "" {
		ext = ".jpg"
	}

	presignURL, key, err := storage.PresignUpload(storage.FolderProfilePic, ext, 15*time.Minute)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Presign URL berhasil dibuat", map[string]string{
		"presign_url": presignURL,
		"key":         key,
	}))
}

// ConfirmProfilePicUpload godoc
// @Summary Simpan key foto profil user setelah upload selesai
// @Description Setelah client upload ke pre-signed URL, kirim key S3 yang diterima untuk disimpan ke database.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Param request body ConfirmProfilePicRequest true "Key S3 hasil upload"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /users/{id}/profile-picture [patch]
func ConfirmProfilePicUpload(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req ConfirmProfilePicRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	pictureURL := storage.PublicURL(req.Key)
	if err := UpdateProfilePictureService(id, pictureURL); err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("User tidak ditemukan")
		}
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Foto profil berhasil disimpan", map[string]string{
		"profile_picture": pictureURL,
	}))
}
