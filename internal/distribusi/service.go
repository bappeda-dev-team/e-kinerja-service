package distribusi

import (
	"aplikasi-internal/internal/exception"
	"aplikasi-internal/internal/user"
)

func validateAllProgrammer(ids []string) error {
	for _, id := range ids {
		roleName, err := user.GetUserRoleNameByID(id)
		if err != nil {
			return exception.ResourceNotFound("programmer tidak ditemukan: " + id)
		}
		if roleName != "programmer" {
			return exception.BadRequest("user bukan programmer: " + id)
		}
	}
	return nil
}

func GetDistribusiDetailServices(sort string) ([]DistribusiFullResponse, error) {
	return GetAll(sort)
}

func GetDistribusiDetailServicesID(id string) (DistribusiFullResponse, error) {
	return GetById(id)
}

func CreateDistribusiServices(permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {
	if err := validateAllProgrammer(req.Pelaksana); err != nil {
		return nil, err
	}

	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	tx, err := beginTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if err := createTx(tx, data); err != nil {
		return nil, err
	}
	if err := insertPelaksanaTx(tx, data.ID, req.Pelaksana); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateKomentarServices(distribusiID string, userID string, req KomentarDistribusiRequest) (*KomentarResponse, error) {

	data := &KomentarDistribusi{
		DistribusiID: distribusiID,
		UserID:       userID,
		Komentars:    req.Komentars,
	}

	err := CreateKomentar(data)
	if err != nil {
		return nil, err
	}

	result, err := GetKomentarById(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateDistribusiServices(id string, permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {
	if err := validateAllProgrammer(req.Pelaksana); err != nil {
		return nil, err
	}

	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	tx, err := beginTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if err := updateTx(tx, id, data); err != nil {
		return nil, err
	}
	if err := replacePelaksanaTx(tx, id, req.Pelaksana); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteDistribusiServices(id string) error {
	return Delete(id)
}
