package pelaksana

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]PelaksanaResponse, error) {
	rows, err := config.DB.Query(
		`SELECT id, distribusi_id, Programmer_id, created_at, updated_at FROM distribusi_pelaksana`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pelaksana []PelaksanaResponse

	for rows.Next() {
		var data PelaksanaResponse
		err := rows.Scan(&data.ID, &data.DistribusiID, &data.ProgrammerID, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		pelaksana = append(pelaksana, data)
	}

	return pelaksana, err

}

func GetId(id string) (PelaksanaResponse, error) {
	var data PelaksanaResponse
	err := config.DB.QueryRow(`SELECT id, distribusi_id, Programmer_id, created_at, updated_at FROM distribusi_pelaksana WHERE id=$1`, id).
		Scan(&data.ID, &data.DistribusiID, &data.ProgrammerID, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return PelaksanaResponse{}, err
	}

	return data, err
}

func GetAllDetail() ([]PelaksanaDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			dp.id,
			d.id, mp.name, ma.name, d.komentar,
			u.id, u.username, u.full_name, u.profile_picture,
			dp.created_at, dp.updated_at
		FROM distribusi_pelaksana dp
		LEFT JOIN distribusi d ON dp.distribusi_id = d.id
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON dp.programmer_id = u.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pelaksana []PelaksanaDetailResponse
	for rows.Next() {
		var data PelaksanaDetailResponse
		err := rows.Scan(
			&data.ID,
			&data.Distribusi.ID, &data.Distribusi.Pemda, &data.Distribusi.Aplikasi, &data.Distribusi.Komentar,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		pelaksana = append(pelaksana, data)
	}
	return pelaksana, err
}

func GetByIdDetail(id string) (PelaksanaDetailResponse, error) {
	var data PelaksanaDetailResponse
	err := config.DB.QueryRow(`
		SELECT
			dp.id,
			d.id, mp.name, ma.name, d.komentar,
			u.id, u.username, u.full_name, u.profile_picture,
			dp.created_at, dp.updated_at
		FROM distribusi_pelaksana dp
		LEFT JOIN distribusi d ON dp.distribusi_id = d.id
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.id = $1
	`, id).Scan(
		&data.ID,
		&data.Distribusi.ID, &data.Distribusi.Pemda, &data.Distribusi.Aplikasi, &data.Distribusi.Komentar,
		&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return PelaksanaDetailResponse{}, err
	}
	return data, err
}

func IsDistribusiIdExists(distribusi_id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM distribusi_pelaksana WHERE distribusi_id = $1)`
	err := config.DB.QueryRow(query, distribusi_id).Scan(&exists)
	return exists, err
}

func IsProgrammerIdExists(programmer_id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM distribusi_pelaksana WHERE programmer_id = $1)`
	err := config.DB.QueryRow(query, programmer_id).Scan(&exists)
	return exists, err
}

func Create(data *Pelaksana) error {
	query := `
		INSERT INTO distribusi_pelaksana (distribusi_id, programmer_id)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.DistribusiID,
		data.ProgrammerID,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func Update(id string, data *Pelaksana) error {
	query := `
		UPDATE distribusi_pelaksana
		SET distribusi_id = $1,
			programmer_id = $2,
		    updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.DistribusiID,
		data.ProgrammerID,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func Delete(id string) error {
	query := `
		DELETE FROM distribusi_pelaksana
		WHERE id = $1
	`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
