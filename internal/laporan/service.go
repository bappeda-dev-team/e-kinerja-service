package laporan

func GetLaporanServices() ([]LaporanFullResponse, error) {
	return GetAll()
}

func GetLaporanByProgrammerServices(programmerID string) ([]LaporanFullResponse, error) {
	return GetAllByProgrammer(programmerID)
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

func CreateLaporanServices(permintaanID string, userID string, req LaporanRequest) (*LaporanFullResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
		Status:          req.Status,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	result, err := GetId(data.ID)
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

func UpdateLaporanServices(id string, permintaanID string, userID string, req LaporanUpdateRequest) (*LaporanFullResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
		Status:          req.Status,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	if req.VerifikasiID != "" && req.StatusVerified != "" {
		err = UpdateStatusVerified(req.VerifikasiID, req.StatusVerified, req.IsSubmittedToVerified)
		if err != nil {
			return nil, err
		}
	}

	result, err := GetId(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteLaporanServices(id string) error {
	return Delete(id)
}
