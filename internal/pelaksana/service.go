package pelaksana

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/user"
)

func validateProgrammerRole(programmerID string) error {
	roleName, err := user.GetUserRoleNameByID(programmerID)
	if err != nil {
		return exception.ResourceNotFound("programmer tidak ditemukan")
	}
	if roleName != "programmer" {
		return exception.BadRequest("user bukan programmer")
	}
	return nil
}

func GetPelaksanaDetailServices(userID string) ([]PelaksanaDetailResponse, error) {
	return GetAllDetail(userID)
}

func GetPelaksanaDetailServicesID(id string) (PelaksanaDetailResponse, error) {
	return GetByIdDetail(id)
}

func CreatePelaksanaServices(distribusiID string, programmerID string) (*PelaksanaDetailResponse, error) {
	if err := validateProgrammerRole(programmerID); err != nil {
		return nil, err
	}

	data := &Pelaksana{
		DistribusiID: distribusiID,
		ProgrammerID: programmerID,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdatePelaksanaServices(id string, distribusiID string, programmerID string) (*PelaksanaDetailResponse, error) {
	if err := validateProgrammerRole(programmerID); err != nil {
		return nil, err
	}

	data := &Pelaksana{
		DistribusiID: distribusiID,
		ProgrammerID: programmerID,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeletePelaksanaServices(id string) error {
	return Delete(id)
}

func MarkAllReadPelaksanaServices(programmerID string) error {
	return MarkAllReadByProgrammerID(programmerID)
}
