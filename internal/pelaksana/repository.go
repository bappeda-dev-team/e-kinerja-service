package pelaksana

import (
	"aplikasi-internal/config"
	"database/sql"
	"fmt"
	"strings"
)

func GetAll() ([]PelaksanaResponse, error) {
	rows, err := config.DB.Query(
		`SELECT id, distribusi_id, programmer_id, is_read, created_at, updated_at FROM distribusi_pelaksana`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pelaksana []PelaksanaResponse

	for rows.Next() {
		var data PelaksanaResponse
		err := rows.Scan(&data.ID, &data.DistribusiID, &data.ProgrammerID, &data.IsRead, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		pelaksana = append(pelaksana, data)
	}

	return pelaksana, err

}

func GetId(id string) (PelaksanaResponse, error) {
	var data PelaksanaResponse
	err := config.DB.QueryRow(`SELECT id, distribusi_id, programmer_id, is_read, created_at, updated_at FROM distribusi_pelaksana WHERE id=$1`, id).
		Scan(&data.ID, &data.DistribusiID, &data.ProgrammerID, &data.IsRead, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return PelaksanaResponse{}, err
	}

	return data, err
}

func GetAllDetail(userID string) ([]PelaksanaDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			dp.id,
			d.id, d.permintaan_id, mp.name, ma.name,
			u.id, u.username, u.full_name, u.profile_picture,
			dp.is_read,
			dp.created_at, dp.updated_at
		FROM distribusi_pelaksana dp
		LEFT JOIN distribusi d ON dp.distribusi_id = d.id
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.programmer_id = $1
		ORDER BY dp.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pelaksana []PelaksanaDetailResponse
	for rows.Next() {
		var data PelaksanaDetailResponse
		err := rows.Scan(
			&data.ID,
			&data.Distribusi.ID, &data.Distribusi.PermintaanID, &data.Distribusi.Pemda, &data.Distribusi.Aplikasi,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
			&data.IsRead,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		data.Distribusi.Komentars = []KomentarInfo{}
		pelaksana = append(pelaksana, data)
	}

	ids := make([]string, len(pelaksana))
	for i, d := range pelaksana {
		ids[i] = d.Distribusi.ID
	}

	komentarMap, err := getKomentarByDistribusiIDs(ids)
	if err != nil {
		return nil, err
	}
	for i, d := range pelaksana {
		if k, ok := komentarMap[d.Distribusi.ID]; ok {
			pelaksana[i].Distribusi.Komentars = k
		}
	}

	return pelaksana, err
}

func GetByIdDetail(id string) (PelaksanaDetailResponse, error) {
	var data PelaksanaDetailResponse
	err := config.DB.QueryRow(`
		SELECT
			dp.id,
			d.id, d.permintaan_id, mp.name, ma.name,
			u.id, u.username, u.full_name, u.profile_picture,
			dp.is_read,
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
		&data.Distribusi.ID, &data.Distribusi.PermintaanID, &data.Distribusi.Pemda, &data.Distribusi.Aplikasi,
		&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
		&data.IsRead,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return PelaksanaDetailResponse{}, err
	}

	komentarMap, err := getKomentarByDistribusiIDs([]string{data.Distribusi.ID})

	if err != nil {
		return PelaksanaDetailResponse{}, err
	}

	if k, ok := komentarMap[data.Distribusi.ID]; ok {
		data.Distribusi.Komentars = k
	} else {
		data.Distribusi.Komentars = []KomentarInfo{}
	}

	return data, err
}

func getKomentarByDistribusiIDs(ids []string) (map[string][]KomentarInfo, error) {
	result := make(map[string][]KomentarInfo)
	if len(ids) == 0 {
		return result, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT kd.distribusi_id, kd.id, u.full_name, kd.komentars, kd.created_at
		FROM komentar_distribusi kd
		LEFT JOIN users u ON kd.user_id = u.id
		WHERE kd.distribusi_id IN (%s)
		ORDER BY kd.created_at ASC
	`, strings.Join(placeholders, ","))

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var distribusiID string
		var p KomentarInfo
		if err := rows.Scan(&distribusiID, &p.ID, &p.FullName, &p.Komentar, &p.CreatedAt); err != nil {
			return nil, err
		}
		result[distribusiID] = append(result[distribusiID], p)
	}
	return result, nil
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
		INSERT INTO distribusi_pelaksana (distribusi_id, programmer_id, is_read)
		VALUES ($1, $2, $3)
		RETURNING id, is_read, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.DistribusiID,
		data.ProgrammerID,
		false,
	).Scan(
		&data.ID,
		&data.IsRead,
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
			is_read = $3,
		    updated_at = NOW()
		WHERE id = $4
		RETURNING is_read, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.DistribusiID,
		data.ProgrammerID,
		false,
		id,
	).Scan(&data.IsRead, &data.UpdatedAt)

	return err
}

func MarkAllReadByProgrammerID(programmerID string) error {
	_, err := config.DB.Exec(`
		UPDATE distribusi_pelaksana
		SET is_read = TRUE,
			updated_at = NOW()
		WHERE programmer_id = $1
			AND is_read = FALSE
	`, programmerID)
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
