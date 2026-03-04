package distribusi

func GetDistribusiServices() ([]Distribusi, error) {
	return GetAll()
}

func GetDistribusiServicesID(id string) (Distribusi, error) {
	return GetById(id)
}
func GetDistribusiNamaServices() ([]DistribusiByNama, error) {
	return GetAllByNama()
}

func GetDistribusiNamaServicesID(id string) (DistribusiByNama, error) {
	return GetByNamaId(id)
}

func CreateDistribusiServices(req DistribusiRequest) (*Distribusi, error) {

	data := &Distribusi{
		PermintaanID: req.PermintaanID,
		AdminID:      req.AdminID,
		Komentar:     req.Komentar,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdateDistribusiServices(id string, req DistribusiRequest) (*Distribusi, error) {

	data := &Distribusi{
		PermintaanID: req.PermintaanID,
		AdminID:      req.AdminID,
		Komentar:     req.Komentar,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	data.ID = id

	return data, nil
}

func DeleteDistribusiServices(id string) error {
	return Delete(id)
}