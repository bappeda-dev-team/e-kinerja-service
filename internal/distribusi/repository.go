package distribusi

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]Distribusi, error) {
	rows, err := config.DB.Query(`SELECT id, permintaan_id, admin_id, komentar, created_at, updated_at FROM distribusi`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permintaan []Distribusi

	for rows.Next() {
		var data Distribusi
		err := rows.Scan(&data.ID, &data.PermintaanID, &data.AdminID, &data.Komentar, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permintaan = append(permintaan, data)
	}

	return permintaan, err

}

func GetById(id string) (Distribusi, error){
	var distribusi Distribusi
	err := config.DB.QueryRow(`SELECT id, permintaan_id, admin_id, komentar, created_at, updated_at FROM distribusi WHERE id=$1`, id).
		Scan(&distribusi.ID, &distribusi.PermintaanID, &distribusi.AdminID, &distribusi.Komentar, &distribusi.CreatedAt, &distribusi.UpdatedAt)

	if err != nil {
		return Distribusi{}, err
	}

	return distribusi, err
}

func GetAllByNama() ([]DistribusiByNama, error) {
	rows, err := config.DB.Query(
		`SELECT d.id, mp.name, ma.name, u.full_name, d.komentar, d.created_at, d.updated_at FROM distribusi d 
		LEFT JOIN permintaan p ON d.permintaan_id = p.id 
		LEFT JOIN users u ON d.admin_id = u.id 
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id 
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribusi []DistribusiByNama

	for rows.Next() {
		var data DistribusiByNama
		err := rows.Scan(&data.ID, &data.Pemda, &data.Aplikasi, &data.Admin, &data.Komentar, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		distribusi = append(distribusi, data)
	}

	return distribusi, err

}

func GetByNamaId(id string) (DistribusiByNama, error){
	var distribusi DistribusiByNama
	err := config.DB.QueryRow(`SELECT d.id, mp.name, ma.name, u.full_name, d.komentar, d.created_at, d.updated_at FROM distribusi d 
	LEFT JOIN permintaan p ON d.permintaan_id = p.id 
	LEFT JOIN users u ON d.admin_id = u.id 
	LEFT JOIN master_pemda mp ON p.pemda_id = mp.id 
	LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id WHERE d.id=$1`, id).
		Scan(&distribusi.ID, &distribusi.Pemda, &distribusi.Aplikasi, &distribusi.Admin, &distribusi.Komentar, &distribusi.CreatedAt, &distribusi.UpdatedAt)

	if err != nil {
		return DistribusiByNama{}, err
	}

	return distribusi, err
}

func Create(data *Distribusi) error {
	query := `
		INSERT INTO distribusi (permintaan_id, admin_id, komentar)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.AdminID,
		data.Komentar,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func Update(id string, data *Distribusi) error {
	query := `
		UPDATE distribusi
		SET permintaan_id = $1,
			admin_id = $2,
			komentar = $3,
		    updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.AdminID,
		data.Komentar,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func Delete(id string) error {
	query := `
		DELETE FROM distribusi
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