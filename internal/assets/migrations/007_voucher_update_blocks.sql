-- +migrate Up

ALTER table blocks
ADD column voucher_update_block bigint;

-- +migrate Down

ALTER table blocks
DROP column voucher_update_block;