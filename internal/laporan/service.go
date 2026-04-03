package laporan

func GetLaporanServices() ([]LaporanFullResponse, error) {
	return GetAll()
}

func GetLaporanServicesID(id string) (LaporanFullResponse, error) {
	return GetId(id)
}

func GetLaporanDetailServices() ([]LaporanDetailResponse, error) {
	return GetAllDetail()
}

func GetLaporanDetailServicesID(id string) (LaporanDetailResponse, error) {
	return GetByIdDetail(id)
}

func CreateLaporanServices(permintaanID string, userID string, req LaporanRequest) (*LaporanDetailResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateVerifikasiService(userID string, LaporanID string) (*VerifikasiResponse, error) {

	data := &Verifikasi{
		LaporanID:    LaporanID,
		ProgrammerID: userID,
	}

	err := CreateVerifikasi(data)
	if err != nil {
		return nil, err
	}

	result, err := GetVerifId(data.ID)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateLaporanServices(id string, permintaanID string, userID string, req LaporanRequest) (*LaporanDetailResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteLaporanServices(id string) error {
	return Delete(id)
}