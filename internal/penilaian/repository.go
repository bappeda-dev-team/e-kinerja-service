package penilaian

import (
	"aplikasi-internal/config"
	"database/sql"
)

func Create(p *Penilaian) error {
	return config.DB.QueryRow(`
		INSERT INTO penilaian (distribusi_id, penilai_id, tingkat_keberhasilan, ketepatan_waktu, komentar, tanggal_selesai)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`, p.DistribusiID, p.PenilaiID, p.TingkatKeberhasilan, p.KetepatanWaktu, p.Komentar, p.TanggalSelesai).
		Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func Update(id string, p *Penilaian) error {
	return config.DB.QueryRow(`
		UPDATE penilaian
		SET tingkat_keberhasilan = $1,
		    ketepatan_waktu = $2,
		    komentar = $3,
		    tanggal_selesai = $4,
		    updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`, p.TingkatKeberhasilan, p.KetepatanWaktu, p.Komentar, p.TanggalSelesai, id).
		Scan(&p.UpdatedAt)
}

func GetAll() ([]PenilaianResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			pn.id,
			d.id, d.permintaan_id, p.menu,
			u.id, u.username, u.full_name,
			pn.tingkat_keberhasilan, pn.ketepatan_waktu, pn.komentar,
			pn.tanggal_selesai, pn.created_at, pn.updated_at
		FROM penilaian pn
		JOIN distribusi d ON pn.distribusi_id = d.id
		JOIN permintaan p ON d.permintaan_id = p.id
		JOIN users u ON pn.penilai_id = u.id
		ORDER BY pn.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []PenilaianResponse
	for rows.Next() {
		var r PenilaianResponse
		if err := rows.Scan(
			&r.ID,
			&r.Distribusi.ID, &r.Distribusi.PermintaanID, &r.Distribusi.Menu,
			&r.Penilai.ID, &r.Penilai.Username, &r.Penilai.FullName,
			&r.TingkatKeberhasilan, &r.KetepatanWaktu, &r.Komentar,
			&r.TanggalSelesai, &r.CreatedAt, &r.UpdatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func GetByDistribusiID(distribusiID string) (*PenilaianResponse, error) {
	var r PenilaianResponse
	err := config.DB.QueryRow(`
		SELECT
			pn.id,
			d.id, d.permintaan_id, p.menu,
			u.id, u.username, u.full_name,
			pn.tingkat_keberhasilan, pn.ketepatan_waktu, pn.komentar,
			pn.tanggal_selesai, pn.created_at, pn.updated_at
		FROM penilaian pn
		JOIN distribusi d ON pn.distribusi_id = d.id
		JOIN permintaan p ON d.permintaan_id = p.id
		JOIN users u ON pn.penilai_id = u.id
		WHERE pn.distribusi_id = $1
	`, distribusiID).Scan(
		&r.ID,
		&r.Distribusi.ID, &r.Distribusi.PermintaanID, &r.Distribusi.Menu,
		&r.Penilai.ID, &r.Penilai.Username, &r.Penilai.FullName,
		&r.TingkatKeberhasilan, &r.KetepatanWaktu, &r.Komentar,
		&r.TanggalSelesai, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func GetByID(id string) (*PenilaianResponse, error) {
	var r PenilaianResponse
	err := config.DB.QueryRow(`
		SELECT
			pn.id,
			d.id, d.permintaan_id, p.menu,
			u.id, u.username, u.full_name,
			pn.tingkat_keberhasilan, pn.ketepatan_waktu, pn.komentar,
			pn.tanggal_selesai, pn.created_at, pn.updated_at
		FROM penilaian pn
		JOIN distribusi d ON pn.distribusi_id = d.id
		JOIN permintaan p ON d.permintaan_id = p.id
		JOIN users u ON pn.penilai_id = u.id
		WHERE pn.id = $1
	`, id).Scan(
		&r.ID,
		&r.Distribusi.ID, &r.Distribusi.PermintaanID, &r.Distribusi.Menu,
		&r.Penilai.ID, &r.Penilai.Username, &r.Penilai.FullName,
		&r.TingkatKeberhasilan, &r.KetepatanWaktu, &r.Komentar,
		&r.TanggalSelesai, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &r, nil
}


func ExistsByID(id string) (bool, error) {
	var exists bool
	err := config.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM penilaian WHERE id = $1)`, id).Scan(&exists)
	return exists, err
}

func GetRawByID(id string) (*Penilaian, error) {
	var p Penilaian
	err := config.DB.QueryRow(`
		SELECT id, distribusi_id, penilai_id, tingkat_keberhasilan, ketepatan_waktu, komentar, tanggal_selesai, created_at, updated_at
		FROM penilaian WHERE id = $1
	`, id).Scan(
		&p.ID, &p.DistribusiID, &p.PenilaiID,
		&p.TingkatKeberhasilan, &p.KetepatanWaktu, &p.Komentar,
		&p.TanggalSelesai, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func CheckDistribusiExists(distribusiID string) (bool, error) {
	var exists bool
	err := config.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM distribusi WHERE id = $1)`, distribusiID).Scan(&exists)
	return exists, err
}

func GetNullableKomentar(id string) (*sql.NullString, error) {
	var k sql.NullString
	err := config.DB.QueryRow(`SELECT komentar FROM penilaian WHERE id = $1`, id).Scan(&k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}
