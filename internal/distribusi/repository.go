package distribusi

import (
	"aplikasi-internal/config"
	"database/sql"
	"fmt"
	"strings"
)

func GetAll() ([]DistribusiFullResponse, error) {
	rows, err := config.DB.Query(`
		SELECT DISTINCT ON (d.id)
			d.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
			v.id, v.status_verified, v.komentar, 
			d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON d.admin_id = u.id
		LEFT JOIN laporan_kinerja l ON p.id = l.permintaan_id
		LEFT JOIN verifikasi v ON l.id = v.laporan_id
		ORDER BY d.id, v.updated_at DESC NULLS LAST
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribusi []DistribusiFullResponse
	for rows.Next() {
		var data DistribusiFullResponse
		var vID, vStatus, vKomentar sql.NullString
		err := rows.Scan(
			&data.ID,
			&data.Permintaan.ID, &data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
			&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Admin.ID, &data.Admin.Username, &data.Admin.FullName, &data.Admin.ProfilePicture,
			&vID, &vStatus, &vKomentar,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		data.Verifikasi = VerifikasiInfo{
			ID:             vID.String,
			StatusVerified: vStatus.String,
			Komentar:       vKomentar.String,
		}
		data.Pelaksana = []PelaksanaInfo{}
		data.Komentar = []KomentarInfo{}
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

	komentarMap, err := getKomentarByDistribusiIDs(ids)
	if err != nil {
		return nil, err
	}
	for i, d := range distribusi {
		if k, ok := komentarMap[d.ID]; ok {
			distribusi[i].Komentar = k
		}
	}

	return distribusi, nil

}

func GetById(id string) (DistribusiFullResponse, error) {
	var data DistribusiFullResponse
	var vID, vStatus, vKomentar sql.NullString
	err := config.DB.QueryRow(`
		SELECT DISTINCT ON (d.id)
			d.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
			v.id, v.status_verified, v.komentar,
			d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN permintaan p ON d.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON d.admin_id = u.id
		LEFT JOIN laporan_kinerja l ON p.id = l.permintaan_id
		LEFT JOIN verifikasi v ON l.id = v.laporan_id
		WHERE d.id = $1
		ORDER BY d.id, v.updated_at DESC NULLS LAST
	`, id).Scan(
		&data.ID,
		&data.Permintaan.ID, &data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
		&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Admin.ID, &data.Admin.Username, &data.Admin.FullName, &data.Admin.ProfilePicture,
		&vID, &vStatus, &vKomentar,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return DistribusiFullResponse{}, err
	}
	data.Verifikasi = VerifikasiInfo{
		ID:             vID.String,
		StatusVerified: vStatus.String,
		Komentar:       vKomentar.String,
	}

	pelaksanaMap, err := getPelaksanaByDistribusiIDs([]string{data.ID})
	komentarMap, err := getKomentarByDistribusiIDs([]string{data.ID})
	if err != nil {
		return DistribusiFullResponse{}, err
	}
	if p, ok := pelaksanaMap[data.ID]; ok {
		data.Pelaksana = p
	} else {
		data.Pelaksana = []PelaksanaInfo{}
	}

	if k, ok := komentarMap[data.ID]; ok {
		data.Komentar = k
	} else {
		data.Komentar = []KomentarInfo{}
	}

	return data, nil
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
		if err := rows.Scan(&distribusiID, &p.ID, &p.FullName, &p.Komentars, &p.CreatedAt); err != nil {
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

func InsertPelaksana(distribusiID string, programmerIDs []string) error {
	for _, programmerID := range programmerIDs {
		_, err := config.DB.Exec(`
			INSERT INTO distribusi_pelaksana (distribusi_id, programmer_id, is_read)
			VALUES ($1, $2, $3)
			ON CONFLICT (distribusi_id, programmer_id) DO NOTHING
		`, distribusiID, programmerID, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReplacePelaksana(distribusiID string, programmerIDs []string) error {
	_, err := config.DB.Exec(`DELETE FROM distribusi_pelaksana WHERE distribusi_id = $1`, distribusiID)
	if err != nil {
		return err
	}
	return InsertPelaksana(distribusiID, programmerIDs)
}

func GetKomentarById(id string) (KomentarResponse, error) {
	var data KomentarResponse
	err := config.DB.QueryRow(`
		SELECT
			kd.id,
			u.full_name, kd.komentars,
			kd.created_at, kd.updated_at
		FROM komentar_distribusi kd
		LEFT JOIN users u ON kd.user_id = u.id
		WHERE kd.id = $1
	`, id).Scan(
		&data.ID, &data.FullName, &data.Komentars,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return KomentarResponse{}, err
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
func CreateKomentar(data *KomentarDistribusi) error {
	query := `
		INSERT INTO komentar_distribusi (distribusi_id, user_id, komentars)
		VALUES ($1, $2, $3)
		RETURNING id, is_read,  created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.DistribusiID,
		data.UserID,
		data.Komentars,
	).Scan(
		&data.ID,
		&data.IsRead,
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
