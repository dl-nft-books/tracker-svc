-- +migrate Up

ALTER TABLE contracts ADD COLUMN IF NOT EXISTS chain_id bigint NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE contracts DROP COLUMN chain_id;

