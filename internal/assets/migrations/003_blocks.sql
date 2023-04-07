-- +migrate Up

create table blocks
(
    id              bigserial primary key,
    contract_address text,
    chain_id  bigint,
    mint_block bigint,
    mint_by_nft_block bigint
);

-- +migrate Down

drop table blocks;

