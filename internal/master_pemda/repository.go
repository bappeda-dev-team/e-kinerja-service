package master_pemda

import "aplikasi-internal/config"

func GetAllMasterPemda() ([]MasterPemda, error) {
	rows, err := config.DB.Query("SELECT id, name, created_at, updated_at FROM master_pemda")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var master_pemda []MasterPemda

	for rows.Next() {
		var pemda MasterPemda
		err := rows.Scan(&pemda.ID, &pemda.Name, &pemda.CreatedAt, &pemda.UpdatedAt)
		if err != nil {
			return nil, err
		}
		master_pemda = append(master_pemda, pemda)
	}

	return master_pemda, err

}

func GetMasterPemdaId(id string) (MasterPemda, error){
	var pemda MasterPemda
	err := config.DB.QueryRow("SELECT id, name, created_at, updated_at FROM master_pemda WHERE id=$1", id).
		Scan(&pemda.ID, &pemda.Name, &pemda.CreatedAt, &pemda.UpdatedAt)

	if err != nil {
		return MasterPemda{}, err
	}

	return pemda, err
}