package laporan

func GetLaporanServices() ([]Laporan, error) {
	return GetAll()
}

func GetLaporanServicesID(id string) (Laporan, error) {
	return GetId(id)
}

func CreateLaporanServices(req LaporanRequest) (*Laporan, error) {

	data := &Laporan{
		PermintaanID:    req.PermintaanID,
		ProgrammerID:    req.ProgrammerID,
		LaporanProgress: req.LaporanProgress,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdateLaporanServices(id string, req LaporanRequest) (*Laporan, error) {

	data := &Laporan{
		PermintaanID:    req.PermintaanID,
		ProgrammerID:    req.ProgrammerID,
		LaporanProgress: req.LaporanProgress,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	data.ID = id

	return data, nil
}

func DeleteLaporanServices(id string) error {
	return Delete(id)
}