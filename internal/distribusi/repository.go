package distribusi

import "aplikasi-internal/config"

func GetAllDistribusi() ([]Distribusi, error) {
	rows, err := config.DB.Query(
		"SELECT d.id, mp.name, ma.name, u.full_name, d.komentar, d.created_at, d.updated_at FROM distribusi d LEFT JOIN permintaan p ON d.permintaan_id = p.id LEFT JOIN users u ON d.admin_id = u.id LEFT JOIN master_pemda mp ON p.pemda_id = mp.id LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribusi []Distribusi

	for rows.Next() {
		var data Distribusi
		err := rows.Scan(&data.ID, &data.Pemda, &data.Aplikasi, &data.Admin, &data.Komentar, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
		distribusi = append(distribusi, data)
	}

	return distribusi, err

}