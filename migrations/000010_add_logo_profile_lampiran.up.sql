-- Tambah kolom logo ke master_pemda
ALTER TABLE master_pemda ADD COLUMN IF NOT EXISTS logo TEXT NOT NULL DEFAULT '';

-- Tambah kolom logo ke master_aplikasi
ALTER TABLE master_aplikasi ADD COLUMN IF NOT EXISTS logo TEXT NOT NULL DEFAULT '';

-- Tambah kolom profile_picture ke users
ALTER TABLE users ADD COLUMN IF NOT EXISTS profile_picture TEXT NOT NULL DEFAULT '';

-- Tambah kolom lampiran ke permintaan (max 3 URL, disimpan sebagai TEXT[])
ALTER TABLE permintaan ADD COLUMN IF NOT EXISTS lampiran TEXT[] NOT NULL DEFAULT '{}';
