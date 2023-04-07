-- +migrate Up

create table nft_payments
(
    id                 bigserial primary key,
    contract_address   text,
    payer_address      text,
    nft_address        text,
    nft_id             bigint,
    floor_price        text,
    price_minted       text,
    book_url           text,
    chain_id integer,
    purchase_timestamp timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS nft_payments;