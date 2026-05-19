package distribusi

func GetDistribusiDetailServices(sort string) ([]DistribusiFullResponse, error) {
	return GetAll(sort)
}

func GetDistribusiDetailServicesID(id string) (DistribusiFullResponse, error) {
	return GetById(id)
}

func CreateDistribusiServices(permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {
	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	if err := Create(data); err != nil {
		return nil, err
	}

	if err := InsertPelaksana(data.ID, req.Pelaksana); err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateKomentarServices(distribusiID string, userID string, req KomentarDistribusiRequest) (*KomentarResponse, error) {

	data := &KomentarDistribusi{
		DistribusiID: distribusiID,
		UserID:       userID,
		Komentars:    req.Komentars,
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

func UpdateDistribusiServices(id string, permintaanID string, adminID string, req DistribusiRequest) (*DistribusiDetailResponse, error) {
	data := &Distribusi{
		PermintaanID: permintaanID,
		AdminID:      adminID,
		Komentar:     req.Komentar,
	}

	if err := Update(id, data); err != nil {
		return nil, err
	}

	if err := ReplacePelaksana(id, req.Pelaksana); err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteDistribusiServices(id string) error {
	return Delete(id)
}
