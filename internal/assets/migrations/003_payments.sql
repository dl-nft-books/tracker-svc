-- +migrate Up

create table payments
(
    id                 bigserial primary key,
    contract_id        bigint    references contracts(id),
    contract_address   text      references contracts(contract),
    payer_address      text,
    token_address      text,
    token_symbol       text,
    token_name         text,
    token_decimals     int8,
    amount             text,
    price_token        text,
    price_minted       text,
    book_url           text,
    purchase_timestamp timestamp
);

-- +migrate Down

drop table payments;

