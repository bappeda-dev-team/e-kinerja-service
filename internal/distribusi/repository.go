package distribusi

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]DistribusiResponse, error) {
	rows, err := config.DB.Query(`SELECT id, permintaan_id, admin_id, komentar, created_at, updated_at FROM distribusi`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribusi []DistribusiResponse

	for rows.Next() {
		var data DistribusiResponse
		err := rows.Scan(&data.ID, &data.PermintaanID, &data.AdminID, &data.Komentar, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		distribusi = append(distribusi, data)
	}

	return distribusi, err

}

func GetById(id string) (DistribusiResponse, error) {
	var data DistribusiResponse
	err := config.DB.QueryRow(`SELECT id, permintaan_id, admin_id, komentar, created_at, updated_at FROM distribusi WHERE id=$1`, id).
		Scan(&data.ID, &data.PermintaanID, &data.AdminID, &data.Komentar, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return DistribusiResponse{}, err
	}

	return data, err
}

func GetAllDetail() ([]DistribusiDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			d.id,
			p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name,
			d.komentar, d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON d.admin_id = u.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribusi []DistribusiDetailResponse
	for rows.Next() {
		var data DistribusiDetailResponse
		err := rows.Scan(
			&data.ID,
			&data.Permintaan.ID, &data.Permintaan.Pemda, &data.Permintaan.Aplikasi,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Admin.ID, &data.Admin.Username, &data.Admin.FullName,
			&data.Komentar, &data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		distribusi = append(distribusi, data)
	}
	return distribusi, err
}

func GetByIdDetail(id string) (DistribusiDetailResponse, error) {
	var data DistribusiDetailResponse
	err := config.DB.QueryRow(`
		SELECT
			d.id,
			p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name,
			d.komentar, d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON d.admin_id = u.id
		WHERE d.id = $1
	`, id).Scan(
		&data.ID,
		&data.Permintaan.ID, &data.Permintaan.Pemda, &data.Permintaan.Aplikasi,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Admin.ID, &data.Admin.Username, &data.Admin.FullName,
		&data.Komentar, &data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return DistribusiDetailResponse{}, err
	}
	return data, err
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