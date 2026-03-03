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

func CreatePermintaanServices(req PermintaanRequest) (*Permintaan, error) {

	data := &Permintaan{
		PemdaID:           req.PemdaID,
		AplikasiID:        req.AplikasiID,
		Menu:              req.Menu,
		KondisiAwal:       req.KondisiAwal,
		KondisiDiharapkan: req.KondisiDiharapkan,
		TanggalPesanan:    req.TanggalPesanan,
		TanggalDeadline:   req.TanggalDeadline,
		CreatedBy:         req.CreatedBy,
	}

	err := Create(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UpdatePermintaanServices(id string, req PermintaanRequest) (*Permintaan, error) {

	data := &Permintaan{
		PemdaID:           req.PemdaID,
		AplikasiID:        req.AplikasiID,
		Menu:              req.Menu,
		KondisiAwal:       req.KondisiAwal,
		KondisiDiharapkan: req.KondisiDiharapkan,
		TanggalPesanan:    req.TanggalPesanan,
		TanggalDeadline:   req.TanggalDeadline,
		CreatedBy:         req.CreatedBy,
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