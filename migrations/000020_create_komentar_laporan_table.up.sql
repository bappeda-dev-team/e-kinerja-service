CREATE TABLE komentar_laporan (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    laporan_id     UUID NOT NULL REFERENCES laporan_kinerja(id) ON DELETE CASCADE,
    user_id        UUID NOT NULL REFERENCES users(id),
    komentar       TEXT,
    is_read        BOOLEAN NOT NULL DEFAULT FALSE,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);