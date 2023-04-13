-- +migrate Up

create table blocks
(
    id              bigserial primary key,
    contract_address text,
    chain_id        bigint,
    token_purchased_block bigint
);

-- +migrate Down

drop table blocks;

