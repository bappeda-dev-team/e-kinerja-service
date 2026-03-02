package master_aplikasi

import "aplikasi-internal/config"

func GetAllMasterAplikasi() ([]MasterAplikasi, error) {
	rows, err := config.DB.Query("SELECT id, name, created_at, updated_at FROM master_aplikasi")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_aplikasi []MasterAplikasi

	for rows.Next() {
		var aplikasi MasterAplikasi
		err := rows.Scan(&aplikasi.ID, &aplikasi.Name, &aplikasi.CreatedAt, &aplikasi.UpdatedAt)
		if err != nil {
			return nil, err
		}
		master_aplikasi = append(master_aplikasi, aplikasi)
	}

	return master_aplikasi, err

}

func GetMasterAplikasiId(id string) (MasterAplikasi, error){
	var aplikasi MasterAplikasi
	err := config.DB.QueryRow("SELECT id, name, created_at, updated_at FROM master_aplikasi WHERE id=$1", id).
		Scan(&aplikasi.ID, &aplikasi.Name, &aplikasi.CreatedAt, &aplikasi.UpdatedAt)

	if err != nil {
		return MasterAplikasi{}, err
	}

	return aplikasi, err
}