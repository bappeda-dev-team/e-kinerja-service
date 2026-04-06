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
		Role:      RoleResponse{ID: user.RoleID},
		Username:  user.Username,
		FullName:  user.FullName,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
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
