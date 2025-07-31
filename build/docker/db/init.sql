create table public.wallet
(
    address char(64)                     not null
        constraint wallet_pk
            primary key,
    amount  double precision default 0.0 not null
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
    amount       double precision not null,
    constraint transaction_pk
        primary key (from_address, to_address)
);

alter table public.transaction
    owner to root;

create extension if not exists pgcrypto;

insert into wallet (address, amount)
select
    encode(digest(gen_random_uuid()::text || random()::text, 'sha256'), 'hex') as address,
    100.0 as amount
from
    generate_series(1, 10);