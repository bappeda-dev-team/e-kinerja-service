CREATE TYPE status_permintaan AS ENUM ('proses', 'selesai', 'revisi');

ALTER TABLE permintaan
    ADD COLUMN status status_permintaan NOT NULL DEFAULT 'proses';
