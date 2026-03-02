package master_aplikasi

func GetMasterAplikasiServices() ([]MasterAplikasi, error) {
	return GetAllMasterAplikasi()
}

func GetMasterAplikasiServicesID(id string) (MasterAplikasi, error) {
	return GetMasterAplikasiId(id)
}