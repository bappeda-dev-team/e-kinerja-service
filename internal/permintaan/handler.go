package permintaan

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/internal/storage"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetPermintaan godoc
// @Summary Ambil semua permintaan
// @Description Mendapatkan daftar permintaan. Gunakan ?expand=names untuk menampilkan nama lengkap pemda, aplikasi, dan pembuat.
// @Tags Permintaan
// @Produce json
// @Param expand query string false "Gunakan 'names' untuk join nama"
// @Success 200 {object} helpers.APIResponse{data=[]PermintaanResponse}
// @Failure 500 {object} helpers.APIResponse
// @Router /permintaan [get]
func GetPermintaan(c echo.Context) error {
	if c.QueryParam("expand") == "names" {
		result, err := GetPermintaanNamaServices()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
	}

	result, err := GetPermintaanServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPermintaanId godoc
// @Summary Ambil permintaan berdasarkan ID
// @Description Mendapatkan data permintaan berdasarkan UUID. Gunakan ?expand=names untuk menampilkan nama lengkap.
// @Tags Permintaan
// @Produce json
// @Param id path string true "Permintaan ID (UUID)"
// @Param expand query string false "Gunakan 'names' untuk join nama"
// @Success 200 {object} helpers.APIResponse{data=[]PermintaanResponse}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /permintaan/{id} [get]
func GetPermintaanId(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	if c.QueryParam("expand") == "names" {
		result, err := GetPermintaanNamaServicesID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return exception.ResourceNotFound("Data tidak ditemukan")
			}
			return err
		}
		return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
	}

	result, err := GetPermintaanServicesID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
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
// @Success 201 {object} helpers.APIResponse{data=PermintaanResponse}
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

	result, err := CreatePermintaanServices(req.PemdaID, req.AplikasiID, userID, req)
	if err != nil {
		return err
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
// @Success 200 {object} helpers.APIResponse{data=PermintaanResponse}
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

	result, err := UpdatePermintaanServices(id, req.PemdaID, req.AplikasiID, userID, req)
	if err != nil {
		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
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
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

// PresignLampiranUpload godoc
// @Summary Buat pre-signed URL untuk upload satu lampiran permintaan
// @Description Menghasilkan pre-signed PUT URL ke S3 untuk satu file. Panggil endpoint ini per file (max 3x), lalu konfirmasi semua key via PATCH /:id/lampiran.
// @Tags Permintaan
// @Produce json
// @Param id path string true "ID Permintaan"
// @Param ext query string false "Ekstensi file, misal .pdf, .png, .jpg (default: .pdf)"
// @Success 200 {object} helpers.APIResponse{data=map[string]string}
// @Failure 400 {object} helpers.APIResponse
// @Router /permintaan/{id}/lampiran/presign [post]
func PresignLampiranUpload(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	ext := c.QueryParam("ext")
	if ext == "" {
		ext = ".pdf"
	}

	presignURL, key, err := storage.PresignUpload(storage.FolderLampiran, ext, 15*time.Minute)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Presign URL berhasil dibuat", map[string]string{
		"presign_url": presignURL,
		"key":         key,
	}))
}

// ConfirmLampiranUpload godoc
// @Summary Simpan key lampiran permintaan setelah semua upload selesai
// @Description Kirim semua key S3 (max 3) setelah client selesai upload. Key akan dikonversi ke public URL dan disimpan ke kolom lampiran.
// @Tags Permintaan
// @Accept json
// @Produce json
// @Param id path string true "ID Permintaan"
// @Param request body ConfirmLampiranRequest true "Keys S3 hasil upload (max 3)"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /permintaan/{id}/lampiran [patch]
func ConfirmLampiranUpload(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req ConfirmLampiranRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	if len(req.Keys) > 3 {
		return exception.BadRequest(fmt.Sprintf("maksimal 3 lampiran, dikirim %d", len(req.Keys)))
	}

	urls := make([]string, len(req.Keys))
	for i, key := range req.Keys {
		urls[i] = storage.PublicURL(key)
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
