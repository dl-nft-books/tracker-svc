-- +migrate Up

create table tokens
(
    id         bigserial primary key,
    contract   text,
    name       text,
    symbol     text,
    last_block bigint,
    unique(contract)
);

-- +migrate Down

drop table tokens;

