package permintaan

func GetPermintaanServices() ([]Permintaan, error) {
	return GetAll()
}

func GetPermintaanServicesID(id string) (Permintaan, error) {
	return GetById(id)
}
func GetPermintaanNamaServices() ([]PermintaanByNama, error) {
	return GetAllByNama()
}

func GetPermintaanNamaServicesID(id string) (PermintaanByNama, error) {
	return GetByNamaId(id)
}

func CreatePermintaanServices(pemdaID string, aplikasiID string, userID string, req PermintaanRequest) (*Permintaan, error) {

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

	return data, nil
}

func UpdatePermintaanServices(id string, pemdaID string, aplikasiID string, userID string, req PermintaanRequest) (*Permintaan, error) {

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

	data.ID = id

	return data, nil
}

func DeletePermintaanServices(id string) error {
	return Delete(id)
}