package penugasan

import (
	"aplikasi-internal/config"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func GetAllByPelaksanaID(pelaksanaID string) ([]PenugasanResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			p.id, p.distribusi_pelaksana_id,
			dp.id, dp.programmer_id, u.full_name,
			p.judul, p.deskripsi, p.deadline, p.prioritas,
			p.estimasi_hari, p.urutan, p.status,
			p.created_at, p.updated_at
		FROM penugasan p
		LEFT JOIN distribusi_pelaksana dp ON p.distribusi_pelaksana_id = dp.id
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE p.distribusi_pelaksana_id = $1
		ORDER BY p.urutan ASC, p.created_at ASC
	`, pelaksanaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PenugasanResponse
	for rows.Next() {
		var data PenugasanResponse
		var pelaksana PelaksanaInfo
		if err := rows.Scan(
			&data.ID, &data.DistribusiPelaksanaID,
			&pelaksana.ID, &pelaksana.ProgrammerID, &pelaksana.FullName,
			&data.Judul, &data.Deskripsi, &data.Deadline, &data.Prioritas,
			&data.EstimasiHari, &data.Urutan, &data.Status,
			&data.CreatedAt, &data.UpdatedAt,
		); err != nil {
			return nil, err
		}
		data.Pelaksana = &pelaksana
		results = append(results, data)
	}
	return results, rows.Err()
}

func GetAllByDistribusiID(distribusiID string) ([]PenugasanResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			p.id, p.distribusi_pelaksana_id,
			dp.id, dp.programmer_id, u.full_name,
			p.judul, p.deskripsi, p.deadline, p.prioritas,
			p.estimasi_hari, p.urutan, p.status,
			p.created_at, p.updated_at
		FROM penugasan p
		LEFT JOIN distribusi_pelaksana dp ON p.distribusi_pelaksana_id = dp.id
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.distribusi_id = $1
		ORDER BY p.urutan ASC, p.created_at ASC
	`, distribusiID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PenugasanResponse
	for rows.Next() {
		var data PenugasanResponse
		var pelaksana PelaksanaInfo
		if err := rows.Scan(
			&data.ID, &data.DistribusiPelaksanaID,
			&pelaksana.ID, &pelaksana.ProgrammerID, &pelaksana.FullName,
			&data.Judul, &data.Deskripsi, &data.Deadline, &data.Prioritas,
			&data.EstimasiHari, &data.Urutan, &data.Status,
			&data.CreatedAt, &data.UpdatedAt,
		); err != nil {
			return nil, err
		}
		data.Pelaksana = &pelaksana
		results = append(results, data)
	}
	return results, rows.Err()
}

func GetAllByProgrammerID(programmerID string) ([]PenugasanResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			p.id, p.distribusi_pelaksana_id,
			dp.id, dp.programmer_id, u.full_name,
			p.judul, p.deskripsi, p.deadline, p.prioritas,
			p.estimasi_hari, p.urutan, p.status,
			p.created_at, p.updated_at
		FROM penugasan p
		LEFT JOIN distribusi_pelaksana dp ON p.distribusi_pelaksana_id = dp.id
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.programmer_id = $1
		ORDER BY p.urutan ASC, p.created_at ASC
	`, programmerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PenugasanResponse
	for rows.Next() {
		var data PenugasanResponse
		var pelaksana PelaksanaInfo
		if err := rows.Scan(
			&data.ID, &data.DistribusiPelaksanaID,
			&pelaksana.ID, &pelaksana.ProgrammerID, &pelaksana.FullName,
			&data.Judul, &data.Deskripsi, &data.Deadline, &data.Prioritas,
			&data.EstimasiHari, &data.Urutan, &data.Status,
			&data.CreatedAt, &data.UpdatedAt,
		); err != nil {
			return nil, err
		}
		data.Pelaksana = &pelaksana
		results = append(results, data)
	}
	return results, rows.Err()
}

func GetByID(id string) (PenugasanResponse, error) {
	var data PenugasanResponse
	var pelaksana PelaksanaInfo
	err := config.DB.QueryRow(`
		SELECT
			p.id, p.distribusi_pelaksana_id,
			dp.id, dp.programmer_id, u.full_name,
			p.judul, p.deskripsi, p.deadline, p.prioritas,
			p.estimasi_hari, p.urutan, p.status,
			p.created_at, p.updated_at
		FROM penugasan p
		LEFT JOIN distribusi_pelaksana dp ON p.distribusi_pelaksana_id = dp.id
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE p.id = $1
	`, id).Scan(
		&data.ID, &data.DistribusiPelaksanaID,
		&pelaksana.ID, &pelaksana.ProgrammerID, &pelaksana.FullName,
		&data.Judul, &data.Deskripsi, &data.Deadline, &data.Prioritas,
		&data.EstimasiHari, &data.Urutan, &data.Status,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return PenugasanResponse{}, err
	}
	data.Pelaksana = &pelaksana
	return data, nil
}

func Create(data *Penugasan) error {
	return config.DB.QueryRow(`
		INSERT INTO penugasan
			(distribusi_pelaksana_id, judul, deskripsi, deadline, prioritas, estimasi_hari, urutan, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`,
		data.DistribusiPelaksanaID,
		data.Judul,
		data.Deskripsi,
		data.Deadline,
		data.Prioritas,
		data.EstimasiHari,
		data.Urutan,
		data.Status,
	).Scan(&data.ID, &data.CreatedAt, &data.UpdatedAt)
}

func Update(id string, data *Penugasan) error {
	return config.DB.QueryRow(`
		UPDATE penugasan
		SET judul = $1, deskripsi = $2, deadline = $3, prioritas = $4,
		    estimasi_hari = $5, urutan = $6, updated_at = NOW()
		WHERE id = $7
		RETURNING updated_at
	`,
		data.Judul,
		data.Deskripsi,
		data.Deadline,
		data.Prioritas,
		data.EstimasiHari,
		data.Urutan,
		id,
	).Scan(&data.UpdatedAt)
}

func UpdateStatus(id, status string) error {
	result, err := config.DB.Exec(`
		UPDATE penugasan SET status = $1, updated_at = NOW() WHERE id = $2
	`, status, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func Reassign(id, distribusiPelaksanaID string) error {
	result, err := config.DB.Exec(`
		UPDATE penugasan SET distribusi_pelaksana_id = $1, updated_at = NOW() WHERE id = $2
	`, distribusiPelaksanaID, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func Delete(id string) error {
	result, err := config.DB.Exec(`DELETE FROM penugasan WHERE id = $1`, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func IsPelaksanaExists(id string) bool {
	var exists bool
	config.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM distribusi_pelaksana WHERE id = $1)`, id).Scan(&exists)
	return exists
}

func parseDeadline(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	formats := []string{time.RFC3339, "2006-01-02"}
	for _, f := range formats {
		t, err := time.Parse(f, s)
		if err == nil {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("format deadline tidak valid, gunakan RFC3339 atau YYYY-MM-DD")
}

func buildPlaceholders(n int) string {
	ph := make([]string, n)
	for i := range ph {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(ph, ",")
}
