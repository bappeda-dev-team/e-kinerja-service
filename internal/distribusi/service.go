package distribusi

func GetDistribusiServices() ([]DistribusiResponse, error) {
	return GetAll()
}

func GetDistribusiServicesID(id string) (DistribusiResponse, error) {
	return GetById(id)
}

func GetDistribusiNamaServices() ([]DistribusiDetailResponse, error) {
	return GetAllByNama()
}

func GetDistribusiNamaServicesID(id string) (DistribusiDetailResponse, error) {
	return GetByNamaId(id)
}

func CreateDistribusiServices(permintaanID string, adminID string, req DistribusiRequest) (*DistribusiResponse, error) {

	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	return &DistribusiResponse{
		ID:           data.ID,
		PermintaanID: data.PermintaanID,
		AdminID:      data.AdminID,
		Komentar:     data.Komentar,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}, nil
}

func UpdateDistribusiServices(id string, permintaanID string, adminID string, req DistribusiRequest) (*DistribusiResponse, error) {

	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	return &DistribusiResponse{
		ID:           id,
		PermintaanID: data.PermintaanID,
		AdminID:      data.AdminID,
		Komentar:     data.Komentar,
		UpdatedAt:    data.UpdatedAt,
	}, nil
}

func DeleteDistribusiServices(id string) error {
	return Delete(id)
}