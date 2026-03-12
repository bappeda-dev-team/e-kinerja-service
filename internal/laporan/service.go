package laporan

func GetLaporanServices() ([]LaporanResponse, error) {
	return GetAll()
}

func GetLaporanServicesID(id string) (LaporanResponse, error) {
	return GetId(id)
}

func CreateLaporanServices(permintaanID string, userID string, req LaporanRequest) (*LaporanResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	return &LaporanResponse{
		ID:              data.ID,
		PermintaanID:    data.PermintaanID,
		ProgrammerID:    data.ProgrammerID,
		LaporanProgress: data.LaporanProgress,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}, nil
}

func UpdateLaporanServices(id string, permintaanID string, userID string, req LaporanRequest) (*LaporanResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	return &LaporanResponse{
		ID:              id,
		PermintaanID:    data.PermintaanID,
		ProgrammerID:    data.ProgrammerID,
		LaporanProgress: data.LaporanProgress,
		UpdatedAt:       data.UpdatedAt,
	}, nil
}

func DeleteLaporanServices(id string) error {
	return Delete(id)
}