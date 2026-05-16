package penugasan

func GetAllService(pelaksanaID, distribusiID, programmerID string) ([]PenugasanResponse, error) {
	if pelaksanaID != "" {
		return GetAllByPelaksanaID(pelaksanaID)
	}
	if distribusiID != "" {
		return GetAllByDistribusiID(distribusiID)
	}
	if programmerID != "" {
		return GetAllByProgrammerID(programmerID)
	}
	return []PenugasanResponse{}, nil
}

func GetByIDService(id string) (PenugasanResponse, error) {
	return GetByID(id)
}

func CreateService(req CreatePenugasanRequest) (*PenugasanResponse, error) {
	deadline, err := parseDeadline(req.Deadline)
	if err != nil {
		return nil, err
	}

	urutan := req.Urutan
	if urutan == 0 {
		urutan = 1
	}

	data := &Penugasan{
		DistribusiPelaksanaID: req.DistribusiPelaksanaID,
		Judul:                 req.Judul,
		Deskripsi:             req.Deskripsi,
		Deadline:              deadline,
		Prioritas:             req.Prioritas,
		EstimasiHari:          req.EstimasiHari,
		Urutan:                urutan,
		Status:                "belum_mulai",
	}

	if err := Create(data); err != nil {
		return nil, err
	}

	result, err := GetByID(data.ID)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateService(id string, req UpdatePenugasanRequest) (*PenugasanResponse, error) {
	deadline, err := parseDeadline(req.Deadline)
	if err != nil {
		return nil, err
	}

	urutan := req.Urutan
	if urutan == 0 {
		urutan = 1
	}

	data := &Penugasan{
		Judul:        req.Judul,
		Deskripsi:    req.Deskripsi,
		Deadline:     deadline,
		Prioritas:    req.Prioritas,
		EstimasiHari: req.EstimasiHari,
		Urutan:       urutan,
	}

	if err := Update(id, data); err != nil {
		return nil, err
	}

	result, err := GetByID(id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateStatusService(id, status string) (*PenugasanResponse, error) {
	if err := UpdateStatus(id, status); err != nil {
		return nil, err
	}
	result, err := GetByID(id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ReassignService(id, distribusiPelaksanaID string) (*PenugasanResponse, error) {
	if err := Reassign(id, distribusiPelaksanaID); err != nil {
		return nil, err
	}
	result, err := GetByID(id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteService(id string) error {
	return Delete(id)
}
