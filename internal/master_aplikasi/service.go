package master_aplikasi

func GetMasterAplikasiServices() ([]MasterAplikasi, error) {
	return GetAllMasterAplikasi()
}

func GetMasterAplikasiServicesID(id string) (MasterAplikasi, error) {
	return GetMasterAplikasiId(id)
}

func CreateMasterAplikasiServices(req CreateMasterAplikasiRequest) (*MasterAplikasi, error) {

	data := &MasterAplikasi{
		Name: req.Name,
	}

	err := CreateMasterAplikasi(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdateMasterAplikasiServices(id string, req CreateMasterAplikasiRequest) (*MasterAplikasi, error) {

	data := &MasterAplikasi{
		Name: req.Name,
	}

	err := UpdateMasterAplikasi(id, data)
	if err != nil {
		return nil, err
	}

	data.ID = id

	return data, nil
}

func DeleteMasterAplikasiServices(id string) error {
	return DeleteMasterAplikasi(id)
}
