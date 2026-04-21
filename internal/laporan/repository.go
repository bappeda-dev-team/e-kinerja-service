package laporan

import (
	"aplikasi-internal/config"
	"database/sql"
	"fmt"
	"strings"
)

func GetAll() ([]LaporanFullResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			l.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
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

	var laporan []LaporanFullResponse
	for rows.Next() {
		var data LaporanFullResponse
		// var vID, vKomentar, vStatus sql.NullString
		err := rows.Scan(
			&data.ID,
			&data.Permintaan.ID,
			&data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
			&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
			&data.LaporanProgress, &data.Status,
			// &vID, &vKomentar, &vStatus,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		// if vID.Valid || vKomentar.Valid || vStatus.Valid {
		// 	data.Verifikasi = &VerifikasiInfo{}

		// 	if vID.Valid {
		// 		data.Verifikasi.ID = &vID.String
		// 	}
		// 	if vKomentar.Valid {
		// 		data.Verifikasi.Komentar = &vKomentar.String
		// 	}
		// 	if vStatus.Valid {
		// 		data.Verifikasi.StatusVerified = &vStatus.String
		// 	}
		// }
		data.Verifikasi = []VerifikasiInfo{}
		laporan = append(laporan, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(laporan) > 0 {
		ids := make([]string, len(laporan))
		for i, l := range laporan {
			ids[i] = l.ID
		}
		verifikasiMap, err := getVerifikasiByLaporanIDs(ids)
		if err != nil {
			return nil, err
		}
		for i, l := range laporan {
			if v, ok := verifikasiMap[l.ID]; ok {
				laporan[i].Verifikasi = v
			}
		}
	}

	return laporan, nil
}

func GetId(id string) (LaporanFullResponse, error) {
	var data LaporanFullResponse
	// var vID, vKomentar, vStatus sql.NullString
	err := config.DB.QueryRow(`
		SELECT
			l.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
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
		&data.Permintaan.ID,
		&data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
		&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
		&data.LaporanProgress, &data.Status,
		// &vID, &vKomentar, &vStatus,
		&data.CreatedAt, &data.UpdatedAt,
	)
	// if vID.Valid || vKomentar.Valid || vStatus.Valid {
	// 		data.Verifikasi = &VerifikasiInfo{}

	// 		if vID.Valid {
	// 			data.Verifikasi.ID = &vID.String
	// 		}
	// 		if vKomentar.Valid {
	// 			data.Verifikasi.Komentar = &vKomentar.String
	// 		}
	// 		if vStatus.Valid {
	// 			data.Verifikasi.StatusVerified = &vStatus.String
	// 		}
	// 	}
	if err != nil {
		return LaporanFullResponse{}, err
	}

	verifikasiMap, err := getVerifikasiByLaporanIDs([]string{data.ID})
	if err != nil {
		return LaporanFullResponse{}, err
	}
	if p, ok := verifikasiMap[data.ID]; ok {
		data.Verifikasi = p
	} else {
		data.Verifikasi = []VerifikasiInfo{}
	}

	return data, err
}

func GetAllByProgrammer(userID string) ([]LaporanFullResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			l.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN permintaan p ON l.permintaan_id = p.id
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON l.programmer_id = u.id
		WHERE l.programmer_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []LaporanFullResponse
	for rows.Next() {
		var data LaporanFullResponse
		err := rows.Scan(
			&data.ID,
			&data.Permintaan.ID,
			&data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
			&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
			&data.LaporanProgress, &data.Status,
			&data.CreatedAt, &data.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		data.Verifikasi = []VerifikasiInfo{}
		laporan = append(laporan, data)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(laporan) > 0 {
		ids := make([]string, len(laporan))
		for i, l := range laporan {
			ids[i] = l.ID
		}
		verifikasiMap, err := getVerifikasiByLaporanIDs(ids)
		if err != nil {
			return nil, err
		}
		for i, l := range laporan {
			if v, ok := verifikasiMap[l.ID]; ok {
				laporan[i].Verifikasi = v
			}
		}
	}

	return laporan, nil
}

func getVerifikasiByLaporanIDs(ids []string) (map[string][]VerifikasiInfo, error) {
	result := make(map[string][]VerifikasiInfo)
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
		SELECT DISTINCT ON (l.id)
		l.id, v.id, v.komentar, v.status_verified, v.is_submitted_to_verified, v.created_at, v.updated_at
		FROM verifikasi v
		LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
		WHERE l.id IN (%s)
		ORDER BY l.id, v.updated_at DESC
	`, strings.Join(placeholders, ","))

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var laporanID string
		var v VerifikasiInfo
		if err := rows.Scan(&laporanID, &v.ID, &v.Komentar, &v.StatusVerified, &v.IsSubmittedToVerified, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, err
		}
		result[laporanID] = append(result[laporanID], v)
	}
	return result, nil
}

func getAllHistory() ([]HistoryResponse, error) {
	rows, err := config.DB.Query(`SELECT id, laporan_id, programmer_id, old_status, new_status, 
	old_progress, new_progress, created_at FROM progress_history`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []HistoryResponse

	for rows.Next() {
		var data HistoryResponse
		err := rows.Scan(&data.ID, &data.LaporanID, &data.ProgrammerID, &data.OldStatus, &data.NewStatus, &data.OldProgress,
		&data.NewProgress, &data.CreatedAt)
		if err != nil {
			return nil, err
		}
		history = append(history, data)
	}

	return history, err
}

func GetAllDetail() ([]LaporanDetailResponse, error) {
	rows, err := config.DB.Query(`
		SELECT
			l.id,
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
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
			&data.Permintaan.ID,
			&data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
			&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
			&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
			&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
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
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			u.id, u.username, u.full_name, u.profile_picture,
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
		&data.Permintaan.ID,
		&data.Permintaan.Pemda.ID, &data.Permintaan.Pemda.Name, &data.Permintaan.Pemda.Logo,
		&data.Permintaan.Aplikasi.ID, &data.Permintaan.Aplikasi.Name, &data.Permintaan.Aplikasi.Logo,
		&data.Permintaan.Menu, &data.Permintaan.KondisiAwal, &data.Permintaan.KondisiDiharapkan,
		&data.Permintaan.TanggalPesanan, &data.Permintaan.TanggalDeadline, &data.Permintaan.Lampiran,
		&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
		&data.LaporanProgress, &data.Status,
		&data.CreatedAt, &data.UpdatedAt,
	)
	if err != nil {
		return LaporanDetailResponse{}, err
	}
	return data, err
}

func GetVerifId(id string) (VerifikasiResponse, error) {
	var data VerifikasiResponse
	err := config.DB.QueryRow(`
		SELECT 
			v.id, v.laporan_id, 
			p.id, mp.id, mp.name, mp.logo, ma.id, ma.name, ma.logo, p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran,
			l.laporan_progress, l.status,
			v.verifikator_id, u.username, u.full_name, u.profile_picture,
			v.status_verified,
			v.created_at, v.updated_at 
			FROM verifikasi v
			LEFT JOIN laporan_kinerja l ON v.laporan_id = l.id
			LEFT JOIN permintaan p ON l.permintaan_id = p.id
			LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
			LEFT JOIN users u ON v.verifikator_id = u.id
			WHERE v.id=$1`, id).
		Scan(&data.ID,
			&data.Laporan.ID, &data.Laporan.Permintaan.ID,
			&data.Laporan.Permintaan.Pemda.ID, &data.Laporan.Permintaan.Pemda.Name, &data.Laporan.Permintaan.Pemda.Logo,
			&data.Laporan.Permintaan.Aplikasi.ID, &data.Laporan.Permintaan.Aplikasi.Name, &data.Laporan.Permintaan.Aplikasi.Logo,
			&data.Laporan.Permintaan.Menu, &data.Laporan.Permintaan.KondisiAwal, &data.Laporan.Permintaan.KondisiDiharapkan,
			&data.Laporan.Permintaan.TanggalPesanan, &data.Laporan.Permintaan.TanggalDeadline, &data.Laporan.Permintaan.Lampiran,
			&data.Laporan.LaporanProgress, &data.Laporan.Status,
			&data.Programmer.ID, &data.Programmer.Username, &data.Programmer.FullName, &data.Programmer.ProfilePicture,
			&data.StatusVerified, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return VerifikasiResponse{}, err
	}

	return data, err
}

func Create(data *Laporan) error {
	query := `
		INSERT INTO laporan_kinerja (permintaan_id, programmer_id, laporan_progress, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.ProgrammerID,
		data.LaporanProgress,
		data.Status,
	).Scan(
		&data.ID,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	return err
}
func CreateHistory(data *History) error {
	query := `
		INSERT INTO progress_history (laporan_id, programmer_id, old_status, new_status, old_progress, new_progress)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	err := config.DB.QueryRow(
		query,
		data.LaporanID,
		data.ProgrammerID,
		data.OldStatus,
		data.NewStatus,
		data.OldProgress,
		data.NewProgress,
	).Scan(
		&data.ID,
		&data.CreatedAt,
	)

	return err
}

func CreateVerifikasi(data *Verifikasi) error {
	query := `
		INSERT INTO verifikasi (laporan_id, verifikator_id)
		VALUES ($1, $2)
		RETURNING id, status_verified, created_at, updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.LaporanID,
		data.ProgrammerID,
	).Scan(
		&data.ID,
		&data.StatusVerified,
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
			status           = $4,
		    updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	err := config.DB.QueryRow(
		query,
		data.PermintaanID,
		data.ProgrammerID,
		data.LaporanProgress,
		data.Status,
		id,
	).Scan(&data.UpdatedAt)

	return err
}

func UpdateStatusVerified(id string, status string, isSubmitted bool) error {
	query := `
		UPDATE verifikasi
		SET status_verified = $1,
			is_submitted_to_verified = $2,
		    updated_at = NOW()
		WHERE id = $3
	`

	_, err := config.DB.Exec(query, status, isSubmitted, id)
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
