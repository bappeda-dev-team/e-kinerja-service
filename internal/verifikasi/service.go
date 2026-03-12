package verifikasi

func GetVerifikasiServices() ([]VerifikasiResponse, error) {
	return GetAll()
}

func GetVerifikasiServicesID(id string) (VerifikasiResponse, error) {
	return GetId(id)
}

func CreateVerifikasiServices(laporanID string, userID string, req VerifikasiRequest) (*VerifikasiResponse, error) {

	data := &Verifikasi{
		LaporanID:      laporanID,
		VerifikatorID:  userID,
		Komentar:       req.Komentar,
		StatusVerified: req.StatusVerified,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	return &VerifikasiResponse{
		ID:             data.ID,
		LaporanID:      data.LaporanID,
		VerifikatorID:  data.VerifikatorID,
		Komentar:       data.Komentar,
		StatusVerified: data.StatusVerified,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}, nil
}

func UpdateVerifikasiServices(id string, laporanID string, userID string, req VerifikasiRequest) (*VerifikasiResponse, error) {

	data := &Verifikasi{
		LaporanID:      laporanID,
		VerifikatorID:  userID,
		Komentar:       req.Komentar,
		StatusVerified: req.StatusVerified,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	return &VerifikasiResponse{
		ID:             id,
		LaporanID:      data.LaporanID,
		VerifikatorID:  data.VerifikatorID,
		Komentar:       data.Komentar,
		StatusVerified: data.StatusVerified,
		UpdatedAt:      data.UpdatedAt,
	}, nil
}

func DeleteVerifikasiServices(id string) error {
	return Delete(id)
}