package penilaian

import (
	"aplikasi-internal/internal/exception"
	"time"
)

func hitungKetepatanWaktu(tanggalSelesai time.Time, deadline string) (string, error) {
	dl, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		// coba format date saja
		dl, err = time.Parse("2006-01-02T15:04:05Z07:00", deadline)
		if err != nil {
			dl, err = time.Parse("2006-01-02 15:04:05+00", deadline)
			if err != nil {
				return "tepat_waktu", nil
			}
		}
	}
	if tanggalSelesai.Before(dl) {
		return "lebih_awal", nil
	} else if tanggalSelesai.Equal(dl) || tanggalSelesai.Before(dl.Add(24*time.Hour)) {
		return "tepat_waktu", nil
	}
	return "terlambat", nil
}

func CreatePenilaianService(req CreatePenilaianRequest, penilaiID string) (*PenilaianResponse, error) {
	if req.TingkatKeberhasilan < 0 || req.TingkatKeberhasilan > 100 {
		return nil, exception.BadRequest("tingkat_keberhasilan harus antara 0 dan 100")
	}

	distribusiExists, err := CheckDistribusiExists(req.DistribusiID)
	if err != nil {
		return nil, err
	}
	if !distribusiExists {
		return nil, exception.ResourceNotFound("distribusi tidak ditemukan")
	}

	tanggalSelesai, err := time.Parse(time.RFC3339, req.TanggalSelesai)
	if err != nil {
		tanggalSelesai, err = time.Parse("2006-01-02", req.TanggalSelesai)
		if err != nil {
			return nil, exception.BadRequest("format tanggal_selesai tidak valid, gunakan ISO 8601")
		}
	}

	deadline, err := GetDeadlineByDistribusiID(req.DistribusiID)
	if err != nil {
		return nil, err
	}

	ketepatanWaktu, err := hitungKetepatanWaktu(tanggalSelesai, deadline)
	if err != nil {
		return nil, err
	}

	p := &Penilaian{
		DistribusiID:        req.DistribusiID,
		PenilaiID:           penilaiID,
		TingkatKeberhasilan: req.TingkatKeberhasilan,
		KetepatanWaktu:      ketepatanWaktu,
		Komentar:            req.Komentar,
		TanggalSelesai:      tanggalSelesai,
	}

	if err := Create(p); err != nil {
		return nil, err
	}

	return GetByID(p.ID)
}

func UpdatePenilaianService(id string, req UpdatePenilaianRequest) (*PenilaianResponse, error) {
	existing, err := GetRawByID(id)
	if err != nil {
		return nil, exception.ResourceNotFound("penilaian tidak ditemukan")
	}

	if req.TingkatKeberhasilan != nil {
		existing.TingkatKeberhasilan = *req.TingkatKeberhasilan
	}
	if req.Komentar != "" {
		existing.Komentar = req.Komentar
	}
	if req.TanggalSelesai != "" {
		ts, err := time.Parse(time.RFC3339, req.TanggalSelesai)
		if err != nil {
			ts, err = time.Parse("2006-01-02", req.TanggalSelesai)
			if err != nil {
				return nil, exception.BadRequest("format tanggal_selesai tidak valid")
			}
		}
		existing.TanggalSelesai = ts

		deadline, err := GetDeadlineByDistribusiID(existing.DistribusiID)
		if err != nil {
			return nil, err
		}
		existing.KetepatanWaktu, _ = hitungKetepatanWaktu(ts, deadline)
	}

	if err := Update(id, existing); err != nil {
		return nil, err
	}

	return GetByID(id)
}

func GetAllPenilaianService() ([]PenilaianResponse, error) {
	return GetAll()
}

func GetByDistribusiIDService(distribusiID string) (*PenilaianResponse, error) {
	return GetByDistribusiID(distribusiID)
}
