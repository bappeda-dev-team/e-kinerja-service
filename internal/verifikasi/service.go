package verifikasi

func GetVerifikasiDetailServices() ([]VerifikasiDetailResponse, error) {
	return GetAllDetail()
}

func GetVerifikasiDetailServicesID(id string) (VerifikasiDetailResponse, error) {
	return GetByIdDetail(id)
}

func CreateVerifikasiServices(laporanID string, userID string, req VerifikasiRequest) (*VerifikasiDetailResponse, error) {

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

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateVerifikasiServices(id string, laporanID string, userID string, req VerifikasiRequest) (*VerifikasiDetailResponse, error) {

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

	result, err := GetByIdDetail(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteVerifikasiServices(id string) error {
	return Delete(id)
}