CREATE TABLE penugasan (
  id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  distribusi_pelaksana_id UUID NOT NULL REFERENCES distribusi_pelaksana(id) ON DELETE CASCADE,
  judul                   VARCHAR(255) NOT NULL,
  deskripsi               TEXT,
  deadline                TIMESTAMPTZ,
  prioritas               VARCHAR(10) NOT NULL DEFAULT 'medium'
                          CHECK (prioritas IN ('low', 'medium', 'high')),
  estimasi_hari           INTEGER,
  urutan                  INTEGER NOT NULL DEFAULT 1,
  status                  VARCHAR(20) NOT NULL DEFAULT 'belum_mulai'
                          CHECK (status IN ('belum_mulai', 'sedang_berjalan', 'selesai', 'revisi')),
  created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
