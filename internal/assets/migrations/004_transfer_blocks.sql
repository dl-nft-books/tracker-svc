-- +migrate Up

create table blocks
(
    id              bigserial primary key,
    contract_id     bigint    unique references contracts(id),
    transfer_block  bigint
);

-- +migrate Down

drop table blocks;

