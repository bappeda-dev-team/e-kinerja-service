package user

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/internal/storage"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func handleDBError(err error) error {
	log.Printf("[handleDBError] type=%T msg=%s\n", err, err.Error())
	if err == sql.ErrNoRows {
		return exception.ResourceNotFound("Data tidak ditemukan")
	}
	return exception.InternalServer("Terjadi kesalahan pada server")
}

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

	var pictureURL string
	if fh, err := c.FormFile("file"); err == nil {
		f, err := fh.Open()
		if err != nil {
			return exception.BadRequest("Gagal membuka file")
		}
		defer f.Close()
		pictureURL, err = storage.UploadFile(f, fh, storage.FolderProfilePic)
		if err != nil {
			return err
		}
	}

	user, err := CreateUserService(req, pictureURL)
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

// UploadProfilePic godoc
// @Summary Upload foto profil user
// @Description Upload foto profil langsung ke S3 via multipart form. Gunakan field "file".
// @Tags Users
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "User ID (UUID)"
// @Param file formData file true "File foto profil"
// @Success 200 {object} helpers.APIResponse{data=map[string]string}
// @Failure 400 {object} helpers.APIResponse
// @Failure 404 {object} helpers.APIResponse
// @Router /users/{id}/profile-picture [patch]
func UploadProfilePic(c echo.Context) error {
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

	pictureURL, err := storage.UploadFile(file, fileHeader, storage.FolderProfilePic)
	if err != nil {
		return err
	}

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

func DeleteUser(c echo.Context) error {

	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return exception.BadRequest("UUID tidak valid")
	}

	err := DeactivateUserService(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return exception.ResourceNotFound("Data tidak ditemukan")
		}
		return err
	}

	return c.JSON(http.StatusOK,
		helpers.SuccessResponse(200, "Akun berhasil dinonaktifkan", nil))
}
