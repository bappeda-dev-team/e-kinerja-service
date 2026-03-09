package pelaksana

import "errors"

func GetPelaksanaServices() ([]Pelaksana, error) {
	return GetAll()
}

func GetPelaksanaServicesID(id string) (Pelaksana, error) {
	return GetId(id)
}
func GetPelaksanaNamaServices() ([]PelaksanaNama, error) {
	return GetAllByNama()
}

func GetPelaksanaNamaServicesID(id string) (PelaksanaNama, error) {
	return GetByNamaId(id)
}

func CreatePelaksanaServices(distribusiID string, programmerID string) (*Pelaksana, error) {

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

	return data, nil
}

func UpdatePelaksanaServices(id string, distribusiID string, programmerID string) (*Pelaksana, error) {

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

	data.ID = id

	return data, nil
}

func DeletePelaksanaServices(id string) error {
	return Delete(id)
}