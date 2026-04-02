package superadmin_dashboard

import (
	"aplikasi-internal/config"
	"aplikasi-internal/internal/permintaan"
	"fmt"
	"strings"
	"time"
)

func GetDashboard() (DashboardResponse, error) {
	// 1. Ambil semua permintaan beserta info pemda, aplikasi, pembuat
	rows, err := config.DB.Query(`
		SELECT
			p.id,
			mp.id, mp.name, mp.logo,
			ma.id, ma.name, ma.logo,
			p.menu, p.kondisi_awal, p.kondisi_diharapkan,
			p.tanggal_pesanan, p.tanggal_deadline, p.lampiran, p.status,
			u.id, u.username, u.full_name, u.profile_picture,
			p.created_at, p.updated_at
		FROM permintaan p
		LEFT JOIN master_pemda mp ON p.pemda_id = mp.id
		LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id
		LEFT JOIN users u ON p.created_by = u.id
		ORDER BY p.created_at DESC
	`)
	if err != nil {
		return DashboardResponse{}, err
	}
	defer rows.Close()

	var items []DashboardPermintaanItem

	for rows.Next() {
		var item DashboardPermintaanItem
		var lamp permintaan.StringArray
		err := rows.Scan(
			&item.ID,
			&item.Pemda.ID, &item.Pemda.Name, &item.Pemda.Logo,
			&item.Aplikasi.ID, &item.Aplikasi.Name, &item.Aplikasi.Logo,
			&item.Menu, &item.KondisiAwal, &item.KondisiDiharapkan,
			&item.TanggalPesanan, &item.TanggalDeadline, &lamp, &item.Status,
			&item.Pembuat.ID, &item.Pembuat.Username, &item.Pembuat.FullName, &item.Pembuat.ProfilePicture,
			&item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return DashboardResponse{}, err
		}
		item.Lampiran = lamp
		item.Distribusi = []DistribusiInfo{}
		item.Laporan = []LaporanInfo{}
		items = append(items, item)
	}

	if len(items) == 0 {
		return DashboardResponse{
			TotalPermintaan: 0,
			TotalDistribusi: 0,
			TotalLaporan:    0,
			Permintaan:      []DashboardPermintaanItem{},
		}, nil
	}

	permintaanIDs := make([]string, len(items))
	for i, item := range items {
		permintaanIDs[i] = item.ID
	}

	// 2. Batch fetch distribusi
	distribusiPerPermintaan, distribusiIDs, err := fetchDistribusi(permintaanIDs)
	if err != nil {
		return DashboardResponse{}, err
	}

	// 3. Batch fetch pelaksana per distribusi
	pelaksanaMap, err := fetchPelaksana(distribusiIDs)
	if err != nil {
		return DashboardResponse{}, err
	}

	// Pasang pelaksana ke distribusi
	totalDistribusi := 0
	for permID, dList := range distribusiPerPermintaan {
		for i, d := range dList {
			if pList, ok := pelaksanaMap[d.ID]; ok {
				distribusiPerPermintaan[permID][i].Pelaksana = pList
			}
		}
		totalDistribusi += len(dList)
	}

	// 4. Batch fetch laporan
	laporanPerPermintaan, totalLaporan, err := fetchLaporan(permintaanIDs)
	if err != nil {
		return DashboardResponse{}, err
	}

	// 5. Gabungkan ke items
	for i, item := range items {
		if dList, ok := distribusiPerPermintaan[item.ID]; ok {
			items[i].Distribusi = dList
		}
		if lList, ok := laporanPerPermintaan[item.ID]; ok {
			items[i].Laporan = lList
		}
	}

	return DashboardResponse{
		TotalPermintaan: len(items),
		TotalDistribusi: totalDistribusi,
		TotalLaporan:    totalLaporan,
		Permintaan:      items,
	}, nil
}

func fetchDistribusi(permintaanIDs []string) (map[string][]DistribusiInfo, []string, error) {
	placeholders := makePlaceholders(permintaanIDs)
	args := toArgs(permintaanIDs)

	query := fmt.Sprintf(`
		SELECT
			d.id, d.permintaan_id,
			u.id, u.username, u.full_name, u.profile_picture,
			d.komentar, d.created_at, d.updated_at
		FROM distribusi d
		LEFT JOIN users u ON d.admin_id = u.id
		WHERE d.permintaan_id IN (%s)
		ORDER BY d.created_at ASC
	`, placeholders)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := make(map[string][]DistribusiInfo)
	var distribusiIDs []string

	for rows.Next() {
		var permID string
		var d DistribusiInfo
		var createdAt, updatedAt time.Time
		err := rows.Scan(
			&d.ID, &permID,
			&d.Admin.ID, &d.Admin.Username, &d.Admin.FullName, &d.Admin.ProfilePicture,
			&d.Komentar, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		d.CreatedAt = createdAt
		d.UpdatedAt = updatedAt
		d.Pelaksana = []PelaksanaInfo{}
		result[permID] = append(result[permID], d)
		distribusiIDs = append(distribusiIDs, d.ID)
	}
	return result, distribusiIDs, nil
}

func fetchPelaksana(distribusiIDs []string) (map[string][]PelaksanaInfo, error) {
	result := make(map[string][]PelaksanaInfo)
	if len(distribusiIDs) == 0 {
		return result, nil
	}

	placeholders := makePlaceholders(distribusiIDs)
	args := toArgs(distribusiIDs)

	query := fmt.Sprintf(`
		SELECT dp.distribusi_id, u.id, u.username, u.full_name
		FROM distribusi_pelaksana dp
		LEFT JOIN users u ON dp.programmer_id = u.id
		WHERE dp.distribusi_id IN (%s)
	`, placeholders)

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

func fetchLaporan(permintaanIDs []string) (map[string][]LaporanInfo, int, error) {
	result := make(map[string][]LaporanInfo)
	if len(permintaanIDs) == 0 {
		return result, 0, nil
	}

	placeholders := makePlaceholders(permintaanIDs)
	args := toArgs(permintaanIDs)

	query := fmt.Sprintf(`
		SELECT
			l.id, l.permintaan_id,
			u.id, u.username, u.full_name, u.profile_picture,
			l.laporan_progress, l.status,
			l.created_at, l.updated_at
		FROM laporan_kinerja l
		LEFT JOIN users u ON l.programmer_id = u.id
		WHERE l.permintaan_id IN (%s)
		ORDER BY l.created_at ASC
	`, placeholders)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	total := 0
	for rows.Next() {
		var permID string
		var l LaporanInfo
		err := rows.Scan(
			&l.ID, &permID,
			&l.Programmer.ID, &l.Programmer.Username, &l.Programmer.FullName, &l.Programmer.ProfilePicture,
			&l.LaporanProgress, &l.Status,
			&l.CreatedAt, &l.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		result[permID] = append(result[permID], l)
		total++
	}
	return result, total, nil
}

func makePlaceholders(ids []string) string {
	ph := make([]string, len(ids))
	for i := range ids {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(ph, ",")
}

func toArgs(ids []string) []interface{} {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	return args
}
