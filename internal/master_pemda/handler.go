package master_pemda

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/internal/storage"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetPemda godoc
// @Summary Ambil semua master pemda
// @Description Mendapatkan daftar master pemda
// @Tags Master Pemda
// @Produce json
// @Success 200 {object} helpers.APIResponse{data=[]MasterPemda}
// @Failure 500 {object} helpers.APIResponse
// @Router /master-pemda [get]
func GetPemda(c echo.Context) error {
	result, err := GetMasterPemdaServices()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// GetPemdaID godoc
// @Summary Ambil master pemda berdasarkan ID
// @Description Mendapatkan data master pemda berdasarkan UUID
// @Tags Master Pemda
// @Produce json
// @Param id path string true "Master Pemda ID (UUID)"
// @Success 200 {object} helpers.APIResponse{data=[]MasterPemda}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /master-pemda/{id} [get]
func GetPemdaID(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	result, err := GetMasterPemdaServicesID(id)

	if err != nil {

		// kalau ID tidak ditemukan
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound,
				helpers.ErrorResponse(404, "Data tidak ditemukan", nil))
		}

		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Berhasil mengambil data", result))
}

// CreatePemda godoc
// @Summary Membuat master pemda baru
// @Description Menambahkan data master pemda
// @Tags Master Pemda
// @Accept json
// @Produce json
// @Param request body MasterPemdaRequest true "Data master pemda"
// @Success 201 {object} helpers.APIResponse{data=MasterPemda}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Router /master-pemda [post]
func CreatePemda(c echo.Context) error {
	var req MasterPemdaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	var logoURL string
	if fh, err := c.FormFile("file"); err == nil {
		f, err := fh.Open()
		if err != nil {
			return exception.BadRequest("Gagal membuka file")
		}
		defer f.Close()
		logoURL, err = storage.UploadFile(f, fh, storage.FolderPemda)
		if err != nil {
			return err
		}
	}

	result, err := CreateMasterPemdaServices(req, logoURL)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated,
		helpers.SuccessResponse(201, "Berhasil membuat data", result))
}

// UpdatePemda godoc
// @Summary Update master pemda
// @Description Mengupdate data master pemda
// @Tags Master Pemda
// @Accept json
// @Produce json
// @Param id path string true "ID Master Pemda"
// @Param request body MasterPemdaRequest true "Data master pemda"
// @Success 200 {object} helpers.APIResponse{data=MasterPemda}
// @Failure 400 {object} helpers.APIResponse{errors=[]string}
// @Failure 404 {object} helpers.APIResponse{errors=[]string}
// @Router /master-pemda/{id} [put]
func UpdatePemda(c echo.Context) error {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	var req MasterPemdaRequest

	if err := helpers.BindAndValidate(c, &req); err != nil {
		return exception.BadRequest("Validasi gagal")
	}

	var logoURL string
	if fh, err := c.FormFile("file"); err == nil {
		f, err := fh.Open()
		if err != nil {
			return exception.BadRequest("Gagal membuka file")
		}
		defer f.Close()
		logoURL, err = storage.UploadFile(f, fh, storage.FolderPemda)
		if err != nil {
			return err
		}
	}

	result, err := UpdateMasterPemdaServices(id, req, logoURL)
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

// DeletePemda godoc
// @Summary Hapus master pemda
// @Description Menghapus data master pemda berdasarkan ID
// @Tags Master Pemda
// @Produce json
// @Param id path string true "ID Master Pemda"
// @Success 200 {object} helpers.APIResponse
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /master-pemda/{id} [delete]
func DeletePemda(c echo.Context) error {
	id := c.Param("id")

	// ✅ Validasi UUID
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeleteMasterPemdaServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}

		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Berhasil menghapus data", nil))
}

// UploadLogo godoc
// @Summary Upload logo pemda
// @Description Upload logo langsung ke S3 via multipart form. Gunakan field "file".
// @Tags Master Pemda
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID Master Pemda"
// @Param file formData file true "File logo"
// @Success 200 {object} helpers.APIResponse{data=map[string]string}
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /master-pemda/{id}/logo [patch]
func UploadLogo(c echo.Context) error {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return exception.BadRequest("File tidak ditemukan, gunakan field 'file'")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return exception.BadRequest("Gagal membuka file")
	}
	defer file.Close()

	logoURL, err := storage.UploadFile(file, fileHeader, storage.FolderPemda)
	if err != nil {
		return err
	}

	if err := UpdateLogoServices(id, logoURL); err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse(200, "Logo berhasil disimpan", map[string]string{
		"logo": logoURL,
	}))
}
