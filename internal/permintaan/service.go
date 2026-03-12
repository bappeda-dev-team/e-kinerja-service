package permintaan

func GetPermintaanServices() ([]PermintaanResponse, error) {
	return GetAll()
}

func GetPermintaanServicesID(id string) (PermintaanResponse, error) {
	return GetById(id)
}

func GetPermintaanNamaServices() ([]PermintaanDetailResponse, error) {
	return GetAllByNama()
}

func GetPermintaanNamaServicesID(id string) (PermintaanDetailResponse, error) {
	return GetByNamaId(id)
}

func CreatePermintaanServices(pemdaID string, aplikasiID string, userID string, req PermintaanRequest) (*PermintaanResponse, error) {

	data := &Permintaan{
		PemdaID:           pemdaID,
		AplikasiID:        aplikasiID,
		Menu:              req.Menu,
		KondisiAwal:       req.KondisiAwal,
		KondisiDiharapkan: req.KondisiDiharapkan,
		TanggalPesanan:    req.TanggalPesanan,
		TanggalDeadline:   req.TanggalDeadline,
		CreatedBy:         userID,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	resp := &PermintaanResponse{
		ID:                data.ID,
		PemdaID:           data.PemdaID,
		AplikasiID:        data.AplikasiID,
		Menu:              data.Menu,
		KondisiAwal:       data.KondisiAwal,
		KondisiDiharapkan: data.KondisiDiharapkan,
		TanggalPesanan:    data.TanggalPesanan,
		TanggalDeadline:   data.TanggalDeadline,
		CreatedBy:         data.CreatedBy,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}

	return resp, nil
}

func UpdatePermintaanServices(id string, pemdaID string, aplikasiID string, userID string, req PermintaanRequest) (*PermintaanResponse, error) {

	data := &Permintaan{
		PemdaID:           pemdaID,
		AplikasiID:        aplikasiID,
		Menu:              req.Menu,
		KondisiAwal:       req.KondisiAwal,
		KondisiDiharapkan: req.KondisiDiharapkan,
		TanggalPesanan:    req.TanggalPesanan,
		TanggalDeadline:   req.TanggalDeadline,
		CreatedBy:         userID,
	}

	err := Update(id, data)
	if err != nil {
		return nil, err
	}

	resp := &PermintaanResponse{
		ID:                id,
		PemdaID:           data.PemdaID,
		AplikasiID:        data.AplikasiID,
		Menu:              data.Menu,
		KondisiAwal:       data.KondisiAwal,
		KondisiDiharapkan: data.KondisiDiharapkan,
		TanggalPesanan:    data.TanggalPesanan,
		TanggalDeadline:   data.TanggalDeadline,
		CreatedBy:         data.CreatedBy,
		UpdatedAt:         data.UpdatedAt,
	}

	return resp, nil
}

func DeletePermintaanServices(id string) error {
	return Delete(id)
}

func UpdateLampiranServices(id string, urls []string) error {
	return UpdateLampiran(id, StringArray(urls))
}