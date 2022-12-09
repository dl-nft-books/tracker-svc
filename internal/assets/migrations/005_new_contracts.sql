-- +migrate Up

-- DO NOT EXECUTE IT IN ONE `ALTER` STATEMENT!
ALTER TABLE contracts RENAME COLUMN contract TO address;
ALTER TABLE contracts RENAME COLUMN last_block TO previous_mint_block;

-- +migrate Down

