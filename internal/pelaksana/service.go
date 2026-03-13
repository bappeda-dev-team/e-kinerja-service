package pelaksana

func GetPelaksanaDetailServices() ([]PelaksanaDetailResponse, error) {
	return GetAllDetail()
}

func GetPelaksanaDetailServicesID(id string) (PelaksanaDetailResponse, error) {
	return GetByIdDetail(id)
}

func CreatePelaksanaServices(distribusiID string, programmerID string) (*PelaksanaDetailResponse, error) {
	data := &Pelaksana{
		DistribusiID: distribusiID,
		ProgrammerID: programmerID,
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

func UpdatePelaksanaServices(id string, distribusiID string, programmerID string) (*PelaksanaDetailResponse, error) {
	data := &Pelaksana{
		DistribusiID: distribusiID,
		ProgrammerID: programmerID,
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

func DeletePelaksanaServices(id string) error {
	return Delete(id)
}