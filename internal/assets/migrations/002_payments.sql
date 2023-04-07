-- +migrate Up

create table payments
(
    id                 bigserial primary key,
    contract_address   text,
    payer_address      text,
    token_address      text,
    token_symbol       text,
    token_name         text,
    token_decimals     int8,
    amount             text,
    price_token        text,
    price_minted       text,
    book_url           text,
    chain_id integer,
    purchase_timestamp timestamp
);

-- +migrate Down

drop table payments;

