package master_aplikasi

import "aplikasi-internal/config"

func GetAllMasterAplikasi() ([]MasterAplikasi, error) {
	rows, err := config.DB.Query("SELECT id, name FROM master_aplikasi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_aplikasi []MasterAplikasi

	for rows.Next() {
		var aplikasi MasterAplikasi
		err := rows.Scan(&aplikasi.ID, &aplikasi.Name)
		if err != nil {
			return nil, err
		}
		master_aplikasi = append(master_aplikasi, aplikasi)
	}

	return master_aplikasi, err

}