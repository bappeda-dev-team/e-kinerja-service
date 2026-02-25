CREATE TABLE permintaan (
    id                   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pemda_id             UUID NOT NULL REFERENCES master_pemda(id),
    aplikasi_id          UUID NOT NULL REFERENCES master_aplikasi(id),
    menu                 VARCHAR(255) NOT NULL,
    kondisi_awal         TEXT NOT NULL,           -- Kondisi Sekarang
    kondisi_diharapkan   TEXT NOT NULL,
    tanggal_pesanan      DATE NOT NULL,
    tanggal_deadline     DATE NOT NULL,
    created_by           UUID NOT NULL REFERENCES users(id),  -- Super Admin
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW()
);