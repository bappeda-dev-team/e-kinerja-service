package permintaan

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
			return exception.BadRequest("pemda_id atau aplikasi_id tidak ditemukan")
		case "23505":
			return exception.Conflict("Data permintaan sudah ada")
		}
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

// GetPermintaan godoc
// @Summary Ambil semua permintaan
// @Description Mendapatkan daftar permintaan
// @Tags Permintaan
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]PermintaanDetailResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /permintaan [get]
func GetPermintaan(c echo.Context) error {
	result, err := GetPermintaanDetailServices()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPermintaanId godoc
// @Summary Ambil permintaan berdasarkan ID
// @Description Mendapatkan data permintaan berdasarkan UUID
// @Tags Permintaan
// @Produce json
// @Param id path string true "Permintaan ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=PermintaanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan/{id} [get]
func GetPermintaanId(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetPermintaanDetailServicesID(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreatePermintaan godoc
// @Summary Membuat permintaan baru
// @Description Menambahkan data permintaan
// @Tags Permintaan
// @Accept json
// @Produce json
// @Param request body PermintaanRequest true "Data permintaan"
// @Success 201 {object} helpers.APIResponse{data=PermintaanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan [post]
func CreatePermintaan(c echo.Context) error {
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	var req PermintaanRequest

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

	result, err := CreatePermintaanServices(req.PemdaID, req.AplikasiID, userID, req, lampiran)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdatePermintaan godoc
// @Summary Update permintaan
// @Description Mengupdate data permintaan
// @Tags Permintaan
// @Accept json
// @Produce json
// @Param id path string true "ID Permintaan"
// @Param request body PermintaanRequest true "Data permintaan"
// @Success 200 {object} helpers.APIResponse{data=PermintaanDetailResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan/{id} [put]
func UpdatePermintaan(c echo.Context) error {
	id := c.Param("id")
	userIDInterface := c.Get("user_id")

	if userIDInterface == nil {
		return c.JSON(401, "user tidak ditemukan di token")
	}

	userID := userIDInterface.(string)

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req PermintaanRequest

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
		existing, err := GetById(id)
		if err == nil {
			lampiran = existing.Lampiran
		}
	}

	result, err := UpdatePermintaanServices(id, req.PemdaID, req.AplikasiID, userID, req, lampiran)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil mengupdate data", result))
}

// DeletePermintaan godoc
// @Summary Hapus permintaan
// @Description Menghapus data permintaan berdasarkan ID
// @Tags Permintaan
// @Produce json
// @Param id path string true "ID Permintaan"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /permintaan/{id} [delete]
func DeletePermintaan(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeletePermintaanServices(id)
	if err != nil {
		return handleDBError(err)
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

// UpdateStatusPermintaan godoc
// @Summary Update status permintaan
// @Description Mengubah status permintaan: proses, selesai, atau revisi
// @Tags Permintaan
// @Accept json
// @Produce json
// @Param id path string true "ID Permintaan"
// @Param request body UpdateStatusRequest true "Status baru"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /permintaan/{id}/status [patch]
func UpdateStatusPermintaan(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req UpdateStatusRequest
	if err := c.Bind(&req); err != nil || req.Status == "" {
		return exception.BadRequest("Field 'status' harus diisi")
	}

	if err := UpdateStatusPermintaanServices(id, req.Status); err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}
		return exception.BadRequest(err.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Status berhasil diupdate", nil))
}

// UploadLampiran godoc
// @Summary Upload lampiran permintaan
// @Description Upload hingga 3 file lampiran langsung ke S3 via multipart form. Gunakan field "files" untuk setiap file.
// @Tags Permintaan
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID Permintaan"
// @Param files formData file true "File lampiran (max 3)"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /permintaan/{id}/lampiran [patch]
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
