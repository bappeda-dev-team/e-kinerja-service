CREATE TABLE progress_history (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
laporan_id UUID NOT NULL REFERENCES laporan_kinerja(id) ON DELETE CASCADE,
programmer_id UUID NOT NULL REFERENCES users(id),
old_status VARCHAR(20),
new_status VARCHAR(20),
old_progress TEXT,
new_progress TEXT,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);