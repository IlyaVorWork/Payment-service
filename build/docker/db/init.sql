create table public.wallet
(
    address char(64)                     not null
        constraint wallet_pk
            primary key,
    balance  numeric default 0.0 not null
);

alter table public.wallet
    owner to root;

create table public.transaction
(
    from_address char(64)         not null
        constraint transaction_wallet_address_fk
            references public.wallet,
    to_address   char(64)         not null
        constraint transaction_wallet_address_fk_2
            references public.wallet,
    created_at timestamp not null,
    amount       numeric not null,
    constraint transaction_pk
        primary key (from_address, to_address, created_at)
);

alter table public.transaction
    owner to root;

create extension if not exists pgcrypto;

insert into wallet (address, balance)
select
    encode(digest(gen_random_uuid()::text || random()::text, 'sha256'), 'hex') as address,
    100.0 as balance
from
    generate_series(1, 10);