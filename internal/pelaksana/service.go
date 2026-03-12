package pelaksana

import "errors"

func GetPelaksanaServices() ([]PelaksanaResponse, error) {
	return GetAll()
}

func GetPelaksanaServicesID(id string) (PelaksanaResponse, error) {
	return GetId(id)
}

func GetPelaksanaNamaServices() ([]PelaksanaDetailResponse, error) {
	return GetAllByNama()
}

func GetPelaksanaNamaServicesID(id string) (PelaksanaDetailResponse, error) {
	return GetByNamaId(id)
}

func CreatePelaksanaServices(distribusiID string, programmerID string) (*PelaksanaResponse, error) {

	distribusiidExists, err := IsDistribusiIdExists(distribusiID)
	if err != nil {
		return nil, err
	}
	if distribusiidExists {
		return nil, errors.New("distribusi_id sudah digunakan")
	}

	programmeridExists, err := IsProgrammerIdExists(programmerID)
	if err != nil {
		return nil, err
	}
	if programmeridExists {
		return nil, errors.New("programmer_id sudah digunakan")
	}

	data := &Pelaksana{
		DistribusiID: distribusiID,
		ProgrammerID: programmerID,
	}

	err = Create(data)
	if err != nil {
		return nil, err
	}

	return &PelaksanaResponse{
		ID:           data.ID,
		DistribusiID: data.DistribusiID,
		ProgrammerID: data.ProgrammerID,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}, nil
}

func UpdatePelaksanaServices(id string, distribusiID string, programmerID string) (*PelaksanaResponse, error) {

	distribusiidExists, err := IsDistribusiIdExists(distribusiID)
	if err != nil {
		return nil, err
	}
	if distribusiidExists {
		return nil, errors.New("distribusi_id sudah digunakan")
	}

	programmeridExists, err := IsProgrammerIdExists(programmerID)
	if err != nil {
		return nil, err
	}
	if programmeridExists {
		return nil, errors.New("programmer_id sudah digunakan")
	}

	data := &Pelaksana{
		DistribusiID: distribusiID,
		ProgrammerID: programmerID,
	}

	err = Update(id, data)
	if err != nil {
		return nil, err
	}

	return &PelaksanaResponse{
		ID:           id,
		DistribusiID: data.DistribusiID,
		ProgrammerID: data.ProgrammerID,
		UpdatedAt:    data.UpdatedAt,
	}, nil
}

func DeletePelaksanaServices(id string) error {
	return Delete(id)
}