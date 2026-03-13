package master_aplikasi

func GetMasterAplikasiServices() ([]MasterAplikasi, error) {
	return GetAllMasterAplikasi()
}

func GetMasterAplikasiServicesID(id string) (MasterAplikasi, error) {
	return GetMasterAplikasiId(id)
}

func CreateMasterAplikasiServices(req CreateMasterAplikasiRequest, logoURL string) (*MasterAplikasi, error) {

	data := &MasterAplikasi{
		Name: req.Name,
		Logo: logoURL,
	}

	err := CreateMasterAplikasi(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdateMasterAplikasiServices(id string, req CreateMasterAplikasiRequest, logoURL string) (*MasterAplikasi, error) {

	data := &MasterAplikasi{
		Name: req.Name,
		Logo: logoURL,
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

func UpdateLogoServices(id string, logoURL string) error {
	return UpdateLogoMasterAplikasi(id, logoURL)
}
