BEGIN;

ALTER TABLE healthit_products ADD COLUMN IF NOT EXISTS acb VARCHAR(500);

COMMIT;