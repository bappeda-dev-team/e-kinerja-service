ALTER TABLE laporan_kinerja DROP CONSTRAINT laporan_kinerja_status_check;

UPDATE laporan_kinerja SET status = CASE status
  WHEN 'putih'  THEN '0'
  WHEN 'merah'  THEN '25'
  WHEN 'orange' THEN '50'
  WHEN 'kuning' THEN '75'
  WHEN 'hijau'  THEN '100'
  ELSE '0'
END;

ALTER TABLE laporan_kinerja
  ALTER COLUMN status SET DEFAULT '0',
  ADD CONSTRAINT laporan_kinerja_status_check
    CHECK (status IN ('0', '25', '50', '75', '100'));
