package permintaan

import "aplikasi-internal/config"

func GetAllPermintaan() ([]Permintaan, error) {
	rows, err := config.DB.Query(
		"SELECT p.id, mp.name, ma.name, p.menu, p.kondisi_awal, p.kondisi_diharapkan, p.tanggal_pesanan, p.tanggal_deadline, u.full_name FROM permintaan p LEFT JOIN master_pemda mp ON p.pemda_id = mp.id LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id LEFT JOIN users u ON p.created_by = u.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permintaan []Permintaan

	for rows.Next() {
		var data Permintaan
		err := rows.Scan(&data.ID, &data.PemdaID, &data.AplikasiID, &data.Menu, &data.KondisiAwal, &data.KondisiDiharapkan, &data.TanggalPesanan, &data.TanggalDeadline, &data.CreatedBy)
		if err != nil {
			return nil, err
		}
		permintaan = append(permintaan, data)
	}

	return permintaan, err

}