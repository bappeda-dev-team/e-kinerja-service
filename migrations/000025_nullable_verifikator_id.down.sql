UPDATE verifikasi SET verifikator_id = (SELECT id FROM users WHERE id IS NOT NULL LIMIT 1) WHERE verifikator_id IS NULL;
ALTER TABLE verifikasi ALTER COLUMN verifikator_id SET NOT NULL;
