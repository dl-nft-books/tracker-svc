-- +migrate Up

create table payments
(
    id                 bigserial primary key,
    contract_address   text,
    token_id       bigint,
    nft_id       bigint,
    book_id       bigint,
    payer_address      text,
    token_address      text,
    token_symbol       text,
    token_name         text,
    token_decimals     int8,
    amount             text,
    price_token        text,
    price_minted       text,
    banner_link       text,
    chain_id integer,
    purchase_timestamp timestamp,
    type     int8
);

-- +migrate Down

drop table payments;

