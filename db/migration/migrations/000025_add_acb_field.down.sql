BEGIN;

ALTER TABLE healthit_products DROP COLUMN IF EXISTS acb CASCADE; 

COMMIT;