package verifikasi

func GetVerifikasiServices() ([]Verifikasi, error) {
	return GetAll()
}

func GetVerifikasiServicesID(id string) (Verifikasi, error) {
	return GetId(id)
}

func CreateVerifikasiServices(laporanID string, userID string, req VerifikasiRequest) (*Verifikasi, error) {

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

	return data, nil
}

func UpdateVerifikasiServices(id string, laporanID string, userID string, req VerifikasiRequest) (*Verifikasi, error) {

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

	data.ID = id

	return data, nil
}

func DeleteVerifikasiServices(id string) error {
	return Delete(id)
}