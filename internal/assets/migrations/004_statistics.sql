-- +migrate Up

create table book_statistics
(
    id              bigserial primary key,
    book_id        bigint UNIQUE default 0,
    amount        bigint,
    usd float
);

create table token_statistics
(
    id              bigserial primary key,
    book_id         bigint UNIQUE default 0,
    token_symbol    text,
    usd_price       float,
    token_price     float
);

create table date_statistics
(
    id              bigserial primary key,
    date        text,
    amount        bigint,
    book_id bigint UNIQUE default 0
);

create table chain_statistics
(
    id              bigserial primary key,
    amount        bigint,
    book_id bigint UNIQUE default 0,
    chain_id bigint
);

-- +migrate Down

drop table book_statistics;
drop table token_statistics;
drop table date_statistics;
drop table chain_statistics;

