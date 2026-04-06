package permintaan

import (
	"fmt"
	"time"
)

func GetPermintaanDetailServices() ([]PermintaanDetailResponse, error) {
	return GetAllDetail()
}

func GetPermintaanDetailServicesID(id string) (PermintaanDetailResponse, error) {
	return GetByIdDetail(id)
}

func parseDate(s string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("format tanggal tidak valid '%s', gunakan YYYY-MM-DD", s)
	}
	return t, nil
}

func CreatePermintaanServices(pemdaID string, aplikasiID string, userID string, req PermintaanRequest, lampiran StringArray) (*PermintaanDetailResponse, error) {

	tanggalPesanan, err := parseDate(req.TanggalPesanan)
	if err != nil {
		return nil, err
	}
	tanggalDeadline, err := parseDate(req.TanggalDeadline)
	if err != nil {
		return nil, err
	}

	data := &Permintaan{
		PemdaID:           pemdaID,
		AplikasiID:        aplikasiID,
		Menu:              req.Menu,
		KondisiAwal:       req.KondisiAwal,
		KondisiDiharapkan: req.KondisiDiharapkan,
		TanggalPesanan:    tanggalPesanan,
		TanggalDeadline:   tanggalDeadline,
		Lampiran:          lampiran,
		CreatedBy:         userID,
	}

	err = Create(data)
	if err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(data.ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func UpdatePermintaanServices(id string, pemdaID string, aplikasiID string, userID string, req PermintaanRequest, lampiran StringArray) (*PermintaanDetailResponse, error) {

	tanggalPesanan, err := parseDate(req.TanggalPesanan)
	if err != nil {
		return nil, err
	}
	tanggalDeadline, err := parseDate(req.TanggalDeadline)
	if err != nil {
		return nil, err
	}

	data := &Permintaan{
		PemdaID:           pemdaID,
		AplikasiID:        aplikasiID,
		Menu:              req.Menu,
		KondisiAwal:       req.KondisiAwal,
		KondisiDiharapkan: req.KondisiDiharapkan,
		TanggalPesanan:    tanggalPesanan,
		TanggalDeadline:   tanggalDeadline,
		Lampiran:          lampiran,
		CreatedBy:         userID,
	}

	err = Update(id, data)
	if err != nil {
		return nil, err
	}

	result, err := GetByIdDetail(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func DeletePermintaanServices(id string) error {
	return Delete(id)
}

func UpdateStatusPermintaanServices(id string, status string) error {
	validStatus := map[string]bool{"proses": true, "selesai": true, "revisi": true}
	if !validStatus[status] {
		return fmt.Errorf("status tidak valid, harus: proses, selesai, atau revisi")
	}
	return UpdateStatus(id, status)
}

func UpdateLampiranServices(id string, urls []string) error {
	return UpdateLampiran(id, StringArray(urls))
}
