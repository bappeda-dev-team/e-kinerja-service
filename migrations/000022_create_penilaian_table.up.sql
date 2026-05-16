CREATE TABLE penilaian (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  distribusi_id UUID NOT NULL REFERENCES distribusi(id) ON DELETE CASCADE,
  penilai_id UUID NOT NULL REFERENCES users(id),
  tingkat_keberhasilan INTEGER NOT NULL CHECK (tingkat_keberhasilan BETWEEN 0 AND 100),
  ketepatan_waktu VARCHAR(20) NOT NULL CHECK (ketepatan_waktu IN ('tepat_waktu', 'terlambat', 'lebih_awal')),
  komentar TEXT,
  tanggal_selesai TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  UNIQUE(distribusi_id)
);
