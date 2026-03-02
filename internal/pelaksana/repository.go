package pelaksana

import "aplikasi-internal/config"

func GetAllPelaksana() ([]Pelaksana, error) {
	rows, err := config.DB.Query(
		"SELECT dp.id, mp.name, ma.name, u.full_name, dp.created_at FROM distribusi_pelaksana dp LEFT JOIN distribusi d ON dp.distribusi_id = d.id LEFT JOIN permintaan p ON d.permintaan_id = p.id LEFT JOIN users u ON dp.programmer_id = u.id LEFT JOIN master_pemda mp ON p.pemda_id = mp.id LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pelaksana []Pelaksana

	for rows.Next() {
		var data Pelaksana
		err := rows.Scan(&data.ID, &data.Pemda, &data.Aplikasi, &data.Programmer, &data.CreatedAt)
		if err != nil {
			return nil, err
		}
		pelaksana = append(pelaksana, data)
	}

	return pelaksana, err

}