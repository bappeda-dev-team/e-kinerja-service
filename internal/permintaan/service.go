package permintaan

func GetPermintaanServices() ([]Permintaan, error) {
	return GetAllPermintaan()
}

func GetPermintaanServicesID(id string) (Permintaan, error) {
	return GetId(id)
}