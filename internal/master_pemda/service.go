package master_pemda

func GetMasterPemdaServices() ([]MasterPemda, error) {
	return GetAllMasterPemda()
}

func GetMasterPemdaServicesID(id string) (MasterPemda, error) {
	return GetMasterPemdaId(id)
}

func CreateMasterPemdaServices(req MasterPemdaRequest) (*MasterPemda, error) {

	data := &MasterPemda{
		Name: req.Name,
	}

	err := CreateMasterPemda(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdateMasterPemdaServices(id string, req MasterPemdaRequest) (*MasterPemda, error) {

	data := &MasterPemda{
		Name: req.Name,
	}

	err := UpdateMasterPemda(id, data)
	if err != nil {
		return nil, err
	}

	data.ID = id

	return data, nil
}

func DeleteMasterPemdaServices(id string) error {
	return DeleteMasterPemda(id)
}