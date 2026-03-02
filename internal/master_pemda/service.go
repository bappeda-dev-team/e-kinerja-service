package master_pemda

func GetMasterPemdaServices() ([]MasterPemda, error) {
	return GetAllMasterPemda()
}

func GetMasterPemdaServicesID(id string) (MasterPemda, error) {
	return GetMasterPemdaId(id)
}