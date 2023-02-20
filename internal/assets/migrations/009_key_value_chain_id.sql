-- +migrate Up

UPDATE key_value set key = CONCAT(chain_id, '-', key);
ALTER table key_value
DROP column IF EXISTS chain_id,
ADD UNIQUE (key);

