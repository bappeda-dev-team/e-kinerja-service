CREATE TABLE distribusi_pelaksana (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    distribusi_id  UUID NOT NULL REFERENCES distribusi(id) ON DELETE CASCADE,
    programmer_id  UUID NOT NULL REFERENCES users(id),
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (distribusi_id, programmer_id)
);