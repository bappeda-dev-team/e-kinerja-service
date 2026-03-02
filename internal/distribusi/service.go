package distribusi

func GetDistribusiServices() ([]Distribusi, error) {
	return GetAllDistribusi()
}

func GetDistribusiServicesID(id string) (Distribusi, error) {
	return GetId(id)
}