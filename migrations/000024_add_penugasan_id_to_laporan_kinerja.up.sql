ALTER TABLE laporan_kinerja
  ADD COLUMN penugasan_id UUID REFERENCES penugasan(id) ON DELETE SET NULL;
