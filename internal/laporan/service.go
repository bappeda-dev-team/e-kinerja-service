package laporan

func GetLaporanServices() ([]LaporanFullResponse, error) {
	return GetAll()
}

func GetLaporanByProgrammerServices(userID string) ([]LaporanFullResponse, error) {
	return GetAllByProgrammer(userID)
}

func GetLaporanServicesID(id string) (LaporanFullResponse, error) {
	return GetId(id)
}

func GetHistoryServices() ([]HistoryResponse, error) {
	return getAllHistory()
}

func GetLaporanDetailServices() ([]LaporanDetailResponse, error) {
	return GetAllDetail()
}

func GetLaporanDetailServicesID(id string) (LaporanDetailResponse, error) {
	return GetByIdDetail(id)
}

func CreateLaporanServices(permintaanID string, userID string, req LaporanRequest, lampiran StringArray) (*LaporanFullResponse, error) {

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
		Status:          req.Status,
		Lampiran:        lampiran,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}
	datahis := &History{
		LaporanID:    data.ID,
		ProgrammerID: userID,
		OldStatus:    data.Status,
		OldProgress:  data.LaporanProgress,
	}

	err = CreateHistory(datahis)
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

func CreateKomentarServices(laporanID string, userID string, req KomentarLaporanRequest) (*KomentarResponse, error) {

	data := &KomentarLaporan{
		LaporanID: laporanID,
		UserID:    userID,
		Komentar:  req.Komentar,
	}

	err := CreateKomentar(data)
	if err != nil {
		return nil, err
	}

	result, err := GetKomentarById(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdateLaporanServices(id string, permintaanID string, userID string, req LaporanUpdateRequest, lampiran StringArray) (*LaporanFullResponse, error) {

	existing, err := GetId(id)
	if err != nil {
		return nil, err
	}

	data := &Laporan{
		PermintaanID:    permintaanID,
		ProgrammerID:    userID,
		LaporanProgress: req.LaporanProgress,
		Status:          req.Status,
		Lampiran:        lampiran,
	}

	err = Update(id, data)
	if err != nil {
		return nil, err
	}

	if req.VerifikasiID != "" && req.StatusVerified != "" {
		err = UpdateStatusVerified(req.VerifikasiID, req.StatusVerified, req.IsSubmittedToVerified)
		if err != nil {
			return nil, err
		}
	}

	datahis := &History{
		LaporanID:    existing.ID,
		ProgrammerID: userID,
		OldStatus:    existing.Status,
		NewStatus:    data.Status,
		OldProgress:  existing.LaporanProgress,
		NewProgress:  data.LaporanProgress,
	}

	err = CreateHistory(datahis)
	if err != nil {
		return nil, err
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

func UpdateLampiranServices(id string, urls []string) error {
	return UpdateLampiran(id, StringArray(urls))
}
