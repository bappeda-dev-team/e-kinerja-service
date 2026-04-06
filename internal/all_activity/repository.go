package all_activity

import (
	"aplikasi-internal/config"
	"database/sql"
)

func GetAllActivities() ([]ActivityItem, error) {
	query := `
		SELECT
			activity_type, activity_id, ref_id,
			actor_id, actor_username, actor_full_name, actor_profile_picture,
			pemda_name, aplikasi_name, menu,
			status, komentar,
			created_at, updated_at
		FROM (

			-- permintaan
			SELECT
				'permintaan'      AS activity_type,
				p.id              AS activity_id,
				p.id              AS ref_id,
				u.id              AS actor_id,
				u.username        AS actor_username,
				u.full_name       AS actor_full_name,
				u.profile_picture AS actor_profile_picture,
				mp.name           AS pemda_name,
				ma.name           AS aplikasi_name,
				p.menu            AS menu,
				p.status          AS status,
				''                AS komentar,
				p.created_at, p.updated_at
			FROM permintaan p
			LEFT JOIN users u  ON p.created_by    = u.id
			LEFT JOIN master_pemda mp  ON p.pemda_id     = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id = ma.id

			UNION ALL

			-- distribusi
			SELECT
				'distribusi'      AS activity_type,
				d.id              AS activity_id,
				p.id              AS ref_id,
				u.id, u.username, u.full_name, u.profile_picture,
				mp.name, ma.name, p.menu,
				''                AS status,
				d.komentar        AS komentar,
				d.created_at, d.updated_at
			FROM distribusi d
			LEFT JOIN permintaan p      ON d.permintaan_id = p.id
			LEFT JOIN users u           ON d.admin_id      = u.id
			LEFT JOIN master_pemda mp   ON p.pemda_id      = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id  = ma.id

			UNION ALL

			-- pelaksana (distribusi_pelaksana)
			SELECT
				'pelaksana'       AS activity_type,
				dp.id             AS activity_id,
				p.id              AS ref_id,
				u.id, u.username, u.full_name, u.profile_picture,
				mp.name, ma.name, p.menu,
				''                AS status,
				''                AS komentar,
				dp.created_at, dp.updated_at
			FROM distribusi_pelaksana dp
			LEFT JOIN distribusi d       ON dp.distribusi_id  = d.id
			LEFT JOIN permintaan p       ON d.permintaan_id   = p.id
			LEFT JOIN users u            ON dp.programmer_id  = u.id
			LEFT JOIN master_pemda mp    ON p.pemda_id        = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id     = ma.id

			UNION ALL

			-- laporan (laporan_kinerja)
			SELECT
				'laporan'         AS activity_type,
				l.id              AS activity_id,
				p.id              AS ref_id,
				u.id, u.username, u.full_name, u.profile_picture,
				mp.name, ma.name, p.menu,
				l.status          AS status,
				l.laporan_progress AS komentar,
				l.created_at, l.updated_at
			FROM laporan_kinerja l
			LEFT JOIN permintaan p       ON l.permintaan_id  = p.id
			LEFT JOIN users u            ON l.programmer_id  = u.id
			LEFT JOIN master_pemda mp    ON p.pemda_id       = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id    = ma.id

			UNION ALL

			-- verifikasi
			SELECT
				'verifikasi'      AS activity_type,
				v.id              AS activity_id,
				p.id              AS ref_id,
				u.id, u.username, u.full_name, u.profile_picture,
				mp.name, ma.name, p.menu,
				v.status_verified AS status,
				COALESCE(v.komentar, '') AS komentar,
				v.created_at, v.updated_at
			FROM verifikasi v
			LEFT JOIN laporan_kinerja l  ON v.laporan_id     = l.id
			LEFT JOIN permintaan p       ON l.permintaan_id  = p.id
			LEFT JOIN users u            ON v.verifikator_id = u.id
			LEFT JOIN master_pemda mp    ON p.pemda_id       = mp.id
			LEFT JOIN master_aplikasi ma ON p.aplikasi_id    = ma.id

		) AS activities
		ORDER BY created_at DESC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []ActivityItem
	for rows.Next() {
		var item ActivityItem
		var refID, actorID, actorUsername, actorFullName, actorPP sql.NullString
		var pemdaName, aplikasiName, menu sql.NullString

		err := rows.Scan(
			&item.ActivityType, &item.ActivityID, &refID,
			&actorID, &actorUsername, &actorFullName, &actorPP,
			&pemdaName, &aplikasiName, &menu,
			&item.Status, &item.Komentar,
			&item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		item.RefID = refID.String
		item.Actor = ActorInfo{
			ID:             actorID.String,
			Username:       actorUsername.String,
			FullName:       actorFullName.String,
			ProfilePicture: actorPP.String,
		}
		item.PemdaName = pemdaName.String
		item.AplikasiName = aplikasiName.String
		item.Menu = menu.String

		activities = append(activities, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return activities, nil
}
