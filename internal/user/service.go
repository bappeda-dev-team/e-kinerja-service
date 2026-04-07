package user

import (
	"aplikasi-internal/internal/helpers"
	"errors"

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
		return nil, errors.New("role tidak ditemukan")
	}

	// cek username duplicate
	usernameExists, err := IsUsernameExists(req.Username)
	if err != nil {
		return nil, err
	}
	if usernameExists {
		return nil, errors.New("username sudah digunakan")
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
		RoleID:    user.RoleID,
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
			return nil, errors.New("role tidak ditemukan")
		}
		fields["role_id"] = req.RoleID
	}
	if req.Username != "" {
		usernameExists, err := IsUsernameExists(req.Username)
		if err != nil {
			return nil, err
		}
		if usernameExists {
			return nil, errors.New("username sudah digunakan")
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
		return nil, errors.New("tidak ada field yang diupdate")
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

func LoginService(req LoginRequest) (string, error) {

	user, err := GetUserByUsername(req.Username)
	if err != nil {
		return "", errors.New("username tidak ditemukan")
	}

	if !user.IsActive {
		return "", errors.New("user tidak aktif")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("password salah")
	}

	token, err := helpers.GenerateToken(
		user.ID,
		user.Username,
		user.RoleID,
		user.RoleName,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func DeactivateUserService(id string) error {

	err := DeactivateUser(id)
	if err != nil {
		return err
	}

	return nil
}
