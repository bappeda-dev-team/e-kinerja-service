CREATE TABLE users (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_id      UUID NOT NULL REFERENCES roles(id),
    username     VARCHAR(100) NOT NULL UNIQUE,
    full_name    VARCHAR(255) NOT NULL,
    password     TEXT NOT NULL,                -- bcrypt hashed
    is_active    BOOLEAN NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);