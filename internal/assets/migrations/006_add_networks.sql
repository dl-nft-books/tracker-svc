-- +migrate Up

ALTER TABLE contracts ADD COLUMN IF NOT EXISTS chain_id bigint NOT NULL DEFAULT 0;
ALTER TABLE key_value
DROP CONSTRAINT IF EXISTS key_value_key_key,
ADD COLUMN IF NOT EXISTS chain_id int NOT NULL DEFAULT 0,
ADD UNIQUE ( key, chain_id);
-- +migrate Down
ALTER TABLE contracts DROP COLUMN IF EXISTS chain_id;
ALTER TABLE key_value
DROP COLUMN IF EXISTS chain_id,
DROP CONSTRAINT IF EXISTS key_value_key_chain_id_key;