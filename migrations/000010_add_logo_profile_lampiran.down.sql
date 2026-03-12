-- Hapus kolom lampiran dari permintaan
ALTER TABLE permintaan DROP COLUMN IF EXISTS lampiran;

-- Hapus kolom profile_picture dari users
ALTER TABLE users DROP COLUMN IF EXISTS profile_picture;

-- Hapus kolom logo dari master_aplikasi
ALTER TABLE master_aplikasi DROP COLUMN IF EXISTS logo;

-- Hapus kolom logo dari master_pemda
ALTER TABLE master_pemda DROP COLUMN IF EXISTS logo;
