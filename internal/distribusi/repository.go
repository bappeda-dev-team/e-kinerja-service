package distribusi

import (
	"aplikasi-internal/config"
	"database/sql"
	"fmt"
	"strings"
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

// getPelaksanaByDistribusiIDs mengambil pelaksana untuk beberapa distribusi_id sekaligus.
// Mengembalikan map[distribusi_id][]PelaksanaInfo
func getPelaksanaByDistribusiIDs(ids []string) (map[string][]PelaksanaInfo, error) {
	result := make(map[string][]PelaksanaInfo)
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
		SELECT dp.distribusi_id, u.id, u.username, u.full_name
		FROM distribusi_pelaksana dp
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.distribusi_id IN (%s)
	`, strings.Join(placeholders, ","))

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var distribusiID string
		var p PelaksanaInfo
		if err := rows.Scan(&distribusiID, &p.ID, &p.Username, &p.FullName); err != nil {
			return nil, err
		}
		result[distribusiID] = append(result[distribusiID], p)
	}
	return result, nil
}

func GetAllDetail() ([]DistribusiDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			d.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
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
			&data.Permintaan.ID, &data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo, 
			&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Admin.ID, &data.Admin.Username, &data.Admin.FullName, &data.Admin.ProfilePicture,
			&data.Komentar, &data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		data.Pelaksana = []PelaksanaInfo{}
		distribusi = append(distribusi, data)
	}

	// batch fetch pelaksana
	ids := make([]string, len(distribusi))
	for i, d := range distribusi {
		ids[i] = d.ID
	}
	pelaksanaMap, err := getPelaksanaByDistribusiIDs(ids)
	if err != nil {
		return nil, err
	}
	for i, d := range distribusi {
		if p, ok := pelaksanaMap[d.ID]; ok {
			distribusi[i].Pelaksana = p
		}
	}

	return distribusi, nil
}

func GetByIdDetail(id string) (DistribusiDetailResponse, error) {
	var data DistribusiDetailResponse
	err := config.DB.QueryRow(`
		SELECT
			d.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
			d.komentar, d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON d.admin_id = u.id
		WHERE d.id = $1
	`, id).Scan(
		&data.ID,
		&data.Permintaan.ID, &data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo, 
		&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Admin.ID, &data.Admin.Username, &data.Admin.FullName, &data.Admin.ProfilePicture,
		&data.Komentar, &data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return DistribusiDetailResponse{}, err
	}

	pelaksanaMap, err := getPelaksanaByDistribusiIDs([]string{data.ID})
	if err != nil {
		return DistribusiDetailResponse{}, err
	}
	if p, ok := pelaksanaMap[data.ID]; ok {
		data.Pelaksana = p
	} else {
		data.Pelaksana = []PelaksanaInfo{}
	}

	return data, nil
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