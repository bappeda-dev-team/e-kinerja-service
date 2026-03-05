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

func CreatePelaksanaServices(req PelaksanaRequest) (*Pelaksana, error) {

	distribusiidExists, err := IsDistribusiIdExists(req.DistribusiID)
	if err != nil {
		return nil, err
	}
	if distribusiidExists {
		return nil, errors.New("programmer sudah ditugaskan di distribusi ini")
	}

	programmeridExists, err := IsProgrammerIdExists(req.ProgrammerID)
	if err != nil {
		return nil, err
	}
	if programmeridExists {
		return nil, errors.New("programmer sudah ditugaskan di distribusi ini")
	}

	data := &Pelaksana{
		DistribusiID: req.DistribusiID,
		ProgrammerID: req.ProgrammerID,
	}

	err = Create(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdatePelaksanaServices(id string, req PelaksanaRequest) (*Pelaksana, error) {

	data := &Pelaksana{
		DistribusiID: req.DistribusiID,
		ProgrammerID: req.ProgrammerID,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	data.ID = id

	return data, nil
}

func DeletePelaksanaServices(id string) error {
	return Delete(id)
}