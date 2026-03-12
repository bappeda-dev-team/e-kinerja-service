package permintaan

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]PermintaanResponse, error) {
	rows, err := config.DB.Query(`SELECT id, pemda_id, aplikasi_id, menu, kondisi_awal, kondisi_diharapkan,
	tanggal_pesanan, tanggal_deadline, lampiran, created_by, created_at, updated_at FROM permintaan`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permintaan []PermintaanResponse

	for rows.Next() {
		var data PermintaanResponse
		err := rows.Scan(&data.ID, &data.PemdaID, &data.AplikasiID, &data.Menu, &data.KondisiAwal, &data.KondisiDiharapkan,
			&data.TanggalPesanan, &data.TanggalDeadline, &data.Lampiran, &data.CreatedBy, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permintaan = append(permintaan, data)
	}

	return permintaan, err

}

func GetById(id string) (PermintaanResponse, error) {
	var data PermintaanResponse
	err := config.DB.QueryRow(`SELECT id, pemda_id, aplikasi_id, menu, kondisi_awal, kondisi_diharapkan,
	tanggal_pesanan, tanggal_deadline, lampiran, created_by, created_at, updated_at FROM permintaan WHERE id=$1`, id).
		Scan(&data.ID, &data.PemdaID, &data.AplikasiID, &data.Menu, &data.KondisiAwal, &data.KondisiDiharapkan,
			&data.TanggalPesanan, &data.TanggalDeadline, &data.Lampiran, &data.CreatedBy, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return PermintaanResponse{}, err
	}

	return data, err
}

func GetAllByNama() ([]PermintaanDetailResponse, error) {
	rows, err := config.DB.Query(
		`SELECT p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan, p.tanggal_pesanan, p.tanggal_deadline, p.lampiran, u.full_name, p.created_at, p.updated_at FROM permintaan p
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON p.created_by = u.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permintaan []PermintaanDetailResponse

	for rows.Next() {
		var data PermintaanDetailResponse
		err := rows.Scan(&data.ID, &data.Pemda, &data.Aplikasi, &data.Menu, &data.KondisiAwal, &data.KondisiDiharapkan, &data.TanggalPesanan, &data.TanggalDeadline, &data.Lampiran, &data.Pembuat, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permintaan = append(permintaan, data)
	}

	return permintaan, err

}

func GetByNamaId(id string) (PermintaanDetailResponse, error) {
	var data PermintaanDetailResponse
	err := config.DB.QueryRow(`SELECT p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan, p.tanggal_pesanan, p.tanggal_deadline, p.lampiran, u.full_name, p.created_at, p.updated_at FROM permintaan p
	LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
	LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
	LEFT JOIN users u ON p.created_by = u.id WHERE p.id=$1`, id).
		Scan(&data.ID, &data.Pemda, &data.Aplikasi, &data.Menu, &data.KondisiAwal, &data.KondisiDiharapkan, &data.TanggalPesanan, &data.TanggalDeadline, &data.Lampiran, &data.Pembuat, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return PermintaanDetailResponse{}, err
	}

	return data, err
}

func Create(data *Permintaan) error {
	query := `
		INSERT INTO permintaan (pemda_id, aplikasi_id, menu, kondisi_awal, kondisi_diharapkan, tanggal_pesanan, tanggal_deadline, lampiran, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PemdaID,
		data.AplikasiID,
		data.Menu,
		data.KondisiAwal,
		data.KondisiDiharapkan,
		data.TanggalPesanan,
		data.TanggalDeadline,
		data.Lampiran,
		data.CreatedBy,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func Update(id string, data *Permintaan) error {
	query := `
		UPDATE permintaan
		SET pemda_id = $1,
			aplikasi_id = $2,
			menu = $3,
			kondisi_awal = $4,
			kondisi_diharapkan = $5,
			tanggal_pesanan = $6,
			tanggal_deadline = $7,
			lampiran = $8,
			created_by = $9,
		    updated_at = NOW()
		WHERE id = $10
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PemdaID,
		data.AplikasiID,
		data.Menu,
		data.KondisiAwal,
		data.KondisiDiharapkan,
		data.TanggalPesanan,
		data.TanggalDeadline,
		data.Lampiran,
		data.CreatedBy,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func UpdateLampiran(id string, lampiran StringArray) error {
	result, err := config.DB.Exec(`UPDATE permintaan SET lampiran = $1, updated_at = NOW() WHERE id = $2`, lampiran, id)
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
	query := `
		DELETE FROM permintaan
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
