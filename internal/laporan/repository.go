package laporan

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAll() ([]LaporanResponse, error) {
	rows, err := config.DB.Query(
		`SELECT id, permintaan_id, Programmer_id, laporan_progress, created_at, updated_at FROM laporan_kinerja`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanResponse

	for rows.Next() {
		var data LaporanResponse
		err := rows.Scan(&data.ID, &data.PermintaanID, &data.ProgrammerID, &data.LaporanProgress, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		laporan = append(laporan, data)
	}

	return laporan, err

}

func GetId(id string) (LaporanResponse, error) {
	var data LaporanResponse
	err := config.DB.QueryRow(`SELECT id, permintaan_id, Programmer_id, laporan_progress, created_at, updated_at FROM laporan_kinerja WHERE id=$1`, id).
		Scan(&data.ID, &data.PermintaanID, &data.ProgrammerID, &data.LaporanProgress, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return LaporanResponse{}, err
	}

	return data, err
}

func GetAllDetail() ([]LaporanDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			l.id,
			p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN permintaan p ON l.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON l.programmer_id = u.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanDetailResponse
	for rows.Next() {
		var data LaporanDetailResponse
		err := rows.Scan(
			&data.ID,
			&data.Permintaan.ID, &data.Permintaan.Pemda, &data.Permintaan.Aplikasi,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName,
			&data.LaporanProgress, &data.Status,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		laporan = append(laporan, data)
	}
	return laporan, err
}

func GetByIdDetail(id string) (LaporanDetailResponse, error) {
	var data LaporanDetailResponse
	err := config.DB.QueryRow(`
		SELECT
			l.id,
			p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN permintaan p ON l.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON l.programmer_id = u.id
		WHERE l.id = $1
	`, id).Scan(
		&data.ID,
		&data.Permintaan.ID, &data.Permintaan.Pemda, &data.Permintaan.Aplikasi,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName,
		&data.LaporanProgress, &data.Status,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return LaporanDetailResponse{}, err
	}
	return data, err
}

func Create(data *Laporan) error {
	query := `
		INSERT INTO laporan_kinerja (permintaan_id, programmer_id, laporan_progress)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.ProgrammerID,
		data.LaporanProgress,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}

func Update(id string, data *Laporan) error {
	query := `
		UPDATE laporan_kinerja
		SET permintaan_id = $1,
			programmer_id = $2,
			laporan_progress = $3,
		    updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.ProgrammerID,
		data.LaporanProgress,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func Delete(id string) error {
	query := `
		DELETE FROM laporan_kinerja
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