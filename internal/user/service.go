package user

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/helpers"
	refreshtoken "aplikasi-internal/internal/refresh_token"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetUserServices() ([]UserResponseDetail, error) {
	return GetAll()
}

func GetUserServicesID(id string) (UserResponseDetail, error) {
	return GetId(id)
}

func CreateUserService(req RegisterRequest, pictureURL string) (*UserResponse, error) {

	// cek role valid
	roleExists, err := IsRoleExists(req.RoleID)
	if err != nil {
		return nil, err
	}
	if !roleExists {
		return nil, exception.ResourceNotFound("role tidak ditemukan")
	}

	// cek username duplicate
	usernameExists, err := IsUsernameExists(req.Username)
	if err != nil {
		return nil, err
	}
	if usernameExists {
		return nil, exception.Conflict("username sudah digunakan")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &User{
		RoleID:         req.RoleID,
		Username:       req.Username,
		FullName:       req.FullName,
		ProfilePicture: pictureURL,
	}

	err = CreateUser(user, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        user.ID,
		Role:      RoleResponse{ID: user.RoleID},
		Username:  user.Username,
		FullName:  user.FullName,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func UpdateUserService(id string, req UpdateUserRequest) (*UserResponseDetail, error) {
	fields := map[string]interface{}{}

	if req.RoleID != "" {
		roleExists, err := IsRoleExists(req.RoleID)
		if err != nil {
			return nil, err
		}
		if !roleExists {
			return nil, exception.ResourceNotFound("role tidak ditemukan")
		}
		fields["role_id"] = req.RoleID
	}
	if req.Username != "" {
		usernameExists, err := IsUsernameExists(req.Username)
		if err != nil {
			return nil, err
		}
		if usernameExists {
			return nil, exception.Conflict("username sudah digunakan")
		}
		fields["username"] = req.Username
	}
	if req.FullName != "" {
		fields["full_name"] = req.FullName
	}
	if req.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		fields["password"] = string(hashed)
	}
	if req.IsActive != nil {
		fields["is_active"] = *req.IsActive
	}

	if len(fields) == 0 {
		return nil, exception.BadRequest("tidak ada field yang diupdate")
	}

	if err := UpdateUser(id, fields); err != nil {
		return nil, err
	}

	result, err := GetId(id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateProfilePictureService(id string, pictureURL string) error {
	return UpdateProfilePicture(id, pictureURL)
}

func LoginService(req LoginRequest) (*LoginResponse, error) {

	user, err := GetUserByUsername(req.Username)
	if err != nil {
		return nil, exception.Unauthorized("username tidak ditemukan")
	}

	if !user.IsActive {
		return nil, exception.Unauthorized("user tidak aktif")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, exception.Unauthorized("password salah")
	}

	accessToken, err := helpers.GenerateToken(
		user.ID,
		user.Username,
		user.RoleID,
		user.RoleName,
	)
	if err != nil {
		return nil, err
	}

	rawRefresh, err := helpers.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	if err := refreshtoken.Save(user.ID, rawRefresh, time.Now().Add(7*24*time.Hour)); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: rawRefresh,
	}, nil
}

func DeactivateUserService(id string) error {

	err := DeactivateUser(id)
	if err != nil {
		return err
	}

	return nil
}
