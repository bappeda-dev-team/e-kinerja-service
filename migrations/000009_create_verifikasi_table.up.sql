CREATE TABLE verifikasi (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    laporan_id       UUID NOT NULL REFERENCES laporan_kinerja(id) ON DELETE CASCADE,
    verifikator_id   UUID NOT NULL REFERENCES users(id),
    komentar         TEXT,
    status_verified  VARCHAR(20) NOT NULL DEFAULT 'pending'
                     CHECK (status_verified IN ('pending', 'approved', 'revision')),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);