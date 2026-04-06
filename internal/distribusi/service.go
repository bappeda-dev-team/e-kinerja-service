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

func UpdateDistribusiServices(id string, permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {

	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
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

func DeleteDistribusiServices(id string) error {
	return Delete(id)
}
