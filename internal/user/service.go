package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetUserServices() ([]User, error) {
	return GetAll()
}

func GetUserServicesID(id string) (User, error) {
	return GetId(id)
}

func RegisterUserService(req RegisterRequest) (*User, error) {

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
		RoleID:   req.RoleID,
		Username: req.Username,
		FullName: req.FullName,
	}

	err = CreateUser(user, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return user, nil
}
