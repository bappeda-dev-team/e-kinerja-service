package master_pemda

import "aplikasi-internal/config"

func GetAllMasterPemda() ([]MasterPemda, error) {
	rows, err := config.DB.Query("SELECT id, name FROM master_pemda")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_pemda []MasterPemda

	for rows.Next() {
		var pemda MasterPemda
		err := rows.Scan(&pemda.ID, &pemda.Name)
		if err != nil {
			return nil, err
		}
		master_pemda = append(master_pemda, pemda)
	}

	return master_pemda, err

}