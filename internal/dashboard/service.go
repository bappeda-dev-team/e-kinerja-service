package dashboard

func GetSuperAdminDashboardService() (SuperAdminDashboardResponse, error) {
	items, err := fetchAllPermintaan()
	if err != nil {
		return SuperAdminDashboardResponse{}, err
	}
	if len(items) == 0 {
		return SuperAdminDashboardResponse{Permintaan: []PermintaanItem{}}, nil
	}

	permIDs := make([]string, len(items))
	for i, item := range items {
		permIDs[i] = item.ID
	}

	distribusiMap, distribusiIDs, err := fetchDistribusiByPermintaanIDs(permIDs)
	if err != nil {
		return SuperAdminDashboardResponse{}, err
	}

	pelaksanaMap, err := fetchPelaksanaByDistribusiIDs(distribusiIDs)
	if err != nil {
		return SuperAdminDashboardResponse{}, err
	}

	totalDistribusi := 0
	for permID, dList := range distribusiMap {
		for i, d := range dList {
			if pList, ok := pelaksanaMap[d.ID]; ok {
				distribusiMap[permID][i].Pelaksana = pList
			}
		}
		totalDistribusi += len(dList)
	}

	laporanMap, totalLaporan, err := fetchLaporanByPermintaanIDs(permIDs)
	if err != nil {
		return SuperAdminDashboardResponse{}, err
	}

	for i, item := range items {
		if dList, ok := distribusiMap[item.ID]; ok {
			items[i].Distribusi = dList
		}
		if lList, ok := laporanMap[item.ID]; ok {
			items[i].Laporan = lList
		}
	}

	return SuperAdminDashboardResponse{
		TotalPermintaan: len(items),
		TotalDistribusi: totalDistribusi,
		TotalLaporan:    totalLaporan,
		Permintaan:      items,
	}, nil
}

func GetAdminDashboardService(adminID string) (AdminDashboardResponse, error) {
	distribusiMap, distribusiIDs, err := fetchDistribusiByAdminID(adminID)
	if err != nil {
		return AdminDashboardResponse{}, err
	}

	pelaksanaMap, err := fetchPelaksanaByDistribusiIDs(distribusiIDs)
	if err != nil {
		return AdminDashboardResponse{}, err
	}

	totalDistribusi := 0
	totalPelaksana := 0
	permIDSet := make(map[string]struct{})

	for permID, dList := range distribusiMap {
		permIDSet[permID] = struct{}{}
		for i, d := range dList {
			if pList, ok := pelaksanaMap[d.ID]; ok {
				distribusiMap[permID][i].Pelaksana = pList
				totalPelaksana += len(pList)
			}
		}
		totalDistribusi += len(dList)
	}

	permIDs := make([]string, 0, len(permIDSet))
	for id := range permIDSet {
		permIDs = append(permIDs, id)
	}

	items, err := fetchPermintaanByIDs(permIDs)
	if err != nil {
		return AdminDashboardResponse{}, err
	}

	laporanMap, _, err := fetchLaporanByPermintaanIDs(permIDs)
	if err != nil {
		return AdminDashboardResponse{}, err
	}

	for i, item := range items {
		if dList, ok := distribusiMap[item.ID]; ok {
			items[i].Distribusi = dList
		}
		if lList, ok := laporanMap[item.ID]; ok {
			items[i].Laporan = lList
		}
	}

	return AdminDashboardResponse{
		TotalPermintaan: len(items),
		TotalDistribusi: totalDistribusi,
		TotalPelaksana:  totalPelaksana,
		Permintaan:      items,
	}, nil
}

func GetProgrammerDashboardService(programmerID string) (ProgrammerDashboardResponse, error) {
	penugasan, err := fetchPenugasanByProgrammerID(programmerID)
	if err != nil {
		return ProgrammerDashboardResponse{}, err
	}
	if penugasan == nil {
		penugasan = []PenugasanItem{}
	}

	laporan, err := fetchLaporanByProgrammerID(programmerID)
	if err != nil {
		return ProgrammerDashboardResponse{}, err
	}
	if laporan == nil {
		laporan = []LaporanItem{}
	}

	return ProgrammerDashboardResponse{
		TotalPenugasan: len(penugasan),
		TotalLaporan:   len(laporan),
		Penugasan:      penugasan,
		Laporan:        laporan,
	}, nil
}

func GetVerifikatorDashboardService() (VerifikatorDashboardResponse, error) {
	laporan, err := fetchLaporanPendingVerifikasi()
	if err != nil {
		return VerifikatorDashboardResponse{}, err
	}
	if laporan == nil {
		laporan = []VerifikatorLaporanItem{}
	}

	return VerifikatorDashboardResponse{
		TotalMenunggu: len(laporan),
		Laporan:       laporan,
	}, nil
}

func GetSuperAdminActivityService() ([]ActivityItem, error) {
	items, err := fetchActivitiesSuperAdmin()
	if err != nil {
		return nil, err
	}
	if items == nil {
		return []ActivityItem{}, nil
	}
	return items, nil
}

func GetAdminActivityService(adminID string) ([]ActivityItem, error) {
	items, err := fetchActivitiesAdmin(adminID)
	if err != nil {
		return nil, err
	}
	if items == nil {
		return []ActivityItem{}, nil
	}
	return items, nil
}

func GetProgrammerActivityService(programmerID string) ([]ActivityItem, error) {
	items, err := fetchActivitiesProgrammer(programmerID)
	if err != nil {
		return nil, err
	}
	if items == nil {
		return []ActivityItem{}, nil
	}
	return items, nil
}

func GetVerifikatorActivityService(verifikatorID string) ([]ActivityItem, error) {
	items, err := fetchActivitiesVerifikator(verifikatorID)
	if err != nil {
		return nil, err
	}
	if items == nil {
		return []ActivityItem{}, nil
	}
	return items, nil
}
