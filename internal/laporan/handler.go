package laporan

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/internal/storage"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func handleDBError(err error) error {
	log.Printf("[handleDBError] type=%T msg=%s\n", err, err.Error())
	if err == sql.ErrNoRows {
		return exception.ResourceNotFound("Data tidak ditemukan")
	}
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23503":
			return exception.BadRequest("permintaan_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Data laporan untuk permintaan ini sudah ada")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

// GetLaporan godoc
// @Summary Ambil laporan milik programmer yang login
// @Description Mendapatkan daftar laporan berdasarkan programmer yang sedang login
// @Tags Laporan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]LaporanFullResponse}
// @Failure 401 {object} helpers.APIResponse
// @Failure 500 {object} helpers.APIResponse
// @Router /laporan [get]
func GetLaporan(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}
	userID := userIDInterface.(string)

	roleInterface := c.Get("name")
	if roleInterface == nil {
		return exception.Unauthorized("role tidak ditemukan")
	}

	roleName := roleInterface.(string)
	var result []LaporanFullResponse
	var err error

	switch roleName {
	case "super_admin", "verifikator":
		result, err = GetLaporanServices()
	case "programmer":
		result, err = GetLaporanByProgrammerServices(userID)
	default:
		return exception.AccessDenied("akses ditolak")
	}

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetLaporanID godoc
// @Summary Ambil laporan berdasarkan ID
// @Description Mendapatkan data laporan berdasarkan UUID
// @Tags Laporan
// @Produce json
// @Param id path string true "Laporan ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=LaporanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan/{id} [get]
func GetLaporanID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetLaporanServicesID(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

func GetHistory(c echo.Context) error {
	result, err := GetHistoryServices()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreateLaporan godoc
// @Summary Membuat laporan baru
// @Description Menambahkan data laporan
// @Tags Laporan
// @Accept json
// @Produce json
// @Param request body LaporanRequest true "Data Laporan"
// @Success 201 {object} helpers.APIResponse{data=LaporanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan [post]
func CreateLaporan(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	var req LaporanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	var lampiran StringArray
	if form, err := c.MultipartForm(); err == nil {
		if fhs := form.File["files"]; len(fhs) > 0 {
			if len(fhs) > 3 {
				return exception.BadRequest(fmt.Sprintf("maksimal 3 lampiran, dikirim %d", len(fhs)))
			}
			for _, fh := range fhs {
				f, err := fh.Open()
				if err != nil {
					return exception.BadRequest("Gagal membuka file")
				}
				url, err := storage.UploadFile(f, fh, storage.FolderLampiran)
				f.Close()
				if err != nil {
					return err
				}
				lampiran = append(lampiran, url)
			}
		}
	}

	result, err := CreateLaporanServices(req.PermintaanID, userID, req, lampiran)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}
func CreateVerif(c echo.Context) error {
	userIDInterface := c.Get("user_id")
	LaporanID := c.Param("laporan_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(LaporanID); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := CreateVerifikasiService(userID, LaporanID)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdateLaporan godoc
// @Summary Update laporan
// @Description Mengupdate data laporan
// @Tags Laporan
// @Accept json
// @Produce json
// @Param id path string true "ID Laporan"
// @Param request body LaporanRequest true "Data laporan"
// @Success 200 {object} helpers.APIResponse{data=LaporanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /laporan/{id} [put]
func UpdateLaporan(c echo.Context) error {
	id := c.Param("id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req LaporanUpdateRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	var lampiran StringArray
	if form, err := c.MultipartForm(); err == nil {
		if fhs := form.File["files"]; len(fhs) > 0 {
			if len(fhs) > 3 {
				return exception.BadRequest(fmt.Sprintf("maksimal 3 lampiran, dikirim %d", len(fhs)))
			}
			for _, fh := range fhs {
				f, err := fh.Open()
				if err != nil {
					return exception.BadRequest("Gagal membuka file")
				}
				url, err := storage.UploadFile(f, fh, storage.FolderLampiran)
				f.Close()
				if err != nil {
					return err
				}
				lampiran = append(lampiran, url)
			}
		}
	}
	if lampiran == nil {
		existing, err := GetId(id)
		if err == nil {
			lampiran = existing.Lampiran
		}
	}

	result, err := UpdateLaporanServices(id, req.PermintaanID, userID, req, lampiran)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengupdate data", result))
}

// DeleteLaporan godoc
// @Summary Hapus laporan
// @Description Menghapus data laporan berdasarkan ID
// @Tags Laporan
// @Produce json
// @Param id path string true "ID Laporan"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /laporan/{id} [delete]
func DeleteLaporan(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteLaporanServices(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

func UploadLampiran(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return exception.BadRequest("Gagal membaca multipart form")
	}

	fileHeaders := form.File["files"]
	if len(fileHeaders) == 0 {
		return exception.BadRequest("File tidak ditemukan, gunakan field 'files'")
	}
	if len(fileHeaders) > 3 {
		return exception.BadRequest(fmt.Sprintf("maksimal 3 lampiran, dikirim %d", len(fileHeaders)))
	}

	urls := make([]string, 0, len(fileHeaders))
	for _, fh := range fileHeaders {
		file, err := fh.Open()
		if err != nil {
			return exception.BadRequest("Gagal membuka file")
		}
		url, err := storage.UploadFile(file, fh, storage.FolderLampiran)
		file.Close()
		if err != nil {
			return err
		}
		urls = append(urls, url)
	}

	if err := UpdateLampiranServices(id, urls); err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Lampiran berhasil disimpan", map[string]interface{}{
		"lampiran": urls,
	}))
}

func CreateKomentarLaporan(c echo.Context) error {
	laporanID := c.Param("laporan_id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return exception.Unauthorized("user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(laporanID); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req KomentarLaporanRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}


	result, err := CreateKomentarServices(laporanID, userID, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23503":
				return exception.BadRequest("laporan_id tidak ditemukan")
			}
		}

		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

