CREATE TABLE komentar_distribusi (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    distribusi_id  UUID NOT NULL REFERENCES distribusi(id) ON DELETE CASCADE,
    user_id        UUID NOT NULL REFERENCES users(id),
    komentars      TEXT,
    is_read        BOOLEAN NOT NULL DEFAULT FALSE,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);