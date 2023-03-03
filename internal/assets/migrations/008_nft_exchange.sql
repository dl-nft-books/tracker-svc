-- +migrate Up

ALTER table contracts
ADD column IF NOT EXISTS previous_mint_by_nft_block bigint;

UPDATE contracts SET previous_mint_by_nft_block = previous_mint_block;

create table nft_payments
(
    id                 bigserial primary key,
    contract_id        bigint    references contracts(id),
    contract_address   text      references contracts(address),
    payer_address      text,
    nft_address        text,
    nft_id             bigint,
    floor_price        text,
    price_minted       text,
    book_url           text,
    purchase_timestamp timestamp
);

-- +migrate Down

ALTER table contracts
DROP column IF EXISTS previous_mint_by_nft_block;

DROP TABLE IF EXISTS nft_payments;