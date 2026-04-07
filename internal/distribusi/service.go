package distribusi

func GetDistribusiDetailServices() ([]DistribusiFullResponse, error) {
	return GetAll()
}

func GetDistribusiDetailServicesID(id string) (DistribusiFullResponse, error) {
	return GetById(id)
}

func CreateDistribusiServices(permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {
	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	if err := Create(data); err != nil {
		return nil, err
	}

	if err := InsertPelaksana(data.ID, req.ProgrammerIDs); err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateDistribusiServices(id string, permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {
	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	if err := Update(id, data); err != nil {
		return nil, err
	}

	if err := ReplacePelaksana(id, req.ProgrammerIDs); err != nil {
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
