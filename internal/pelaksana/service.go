package pelaksana

func GetPelaksanaServices() ([]Pelaksana, error) {
	return GetAllPelaksana()
}

func GetPelaksanaServicesID(id string) (Pelaksana, error) {
	return GetId(id)
}