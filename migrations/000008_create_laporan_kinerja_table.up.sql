CREATE TABLE laporan_kinerja (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    permintaan_id   UUID NOT NULL REFERENCES permintaan(id) ON DELETE CASCADE,
    programmer_id   UUID NOT NULL REFERENCES users(id),
    laporan_progress TEXT NOT NULL,
    status          VARCHAR(20) NOT NULL DEFAULT 'putih'
                    CHECK (status IN ('putih', 'merah', 'orange', 'kuning', 'hijau')),
    -- putih=0%, merah=25%, orange=50%, kuning=75%, hijau=100%
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);