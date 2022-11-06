-- +migrate Up

create table payments
(
    id            bigserial primary key,
    contract_id   bigint    references contracts(id),
    payer_address text,
    token_address text,
    token_symbol  text,
    token_name    text,
    amount        text,
    price         text
);

-- +migrate Down

drop table payments;

