ALTER TABLE laporan_kinerja DROP CONSTRAINT laporan_kinerja_status_check;

UPDATE laporan_kinerja SET status = CASE status
  WHEN '0'   THEN 'putih'
  WHEN '25'  THEN 'merah'
  WHEN '50'  THEN 'orange'
  WHEN '75'  THEN 'kuning'
  WHEN '100' THEN 'hijau'
  ELSE 'putih'
END;

ALTER TABLE laporan_kinerja
  ALTER COLUMN status SET DEFAULT 'putih',
  ADD CONSTRAINT laporan_kinerja_status_check
    CHECK (status IN ('putih', 'merah', 'orange', 'kuning', 'hijau'));
