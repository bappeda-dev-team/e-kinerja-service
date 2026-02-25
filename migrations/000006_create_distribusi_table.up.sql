CREATE TABLE distribusi (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    permintaan_id UUID NOT NULL REFERENCES permintaan(id) ON DELETE CASCADE,
    admin_id      UUID NOT NULL REFERENCES users(id),
    komentar      TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);