create table bookings
(
    id          varchar(36)                  not null,
    title       varchar                      not null,
    comment     varchar                      not null,
    date        date                         not null,
    amount      numeric(10, 2) default 0.0   not null,
    category_id varchar(36)                  not null
        constraint bookings_categories_id_fk
            references categories
            on delete cascade,
    account_id  varchar(36)                  not null
        constraint bookings_accounts_id_fk
            references accounts
            on delete cascade,
    created     timestamp      default now() not null,
    updated     timestamp      default now() not null
);

create index bookings_account_id_date_index
    on bookings (account_id, date);

create index bookings_account_id_index
    on bookings (account_id);

create index bookings_category_id_date_index
    on bookings (category_id, date);

create index bookings_category_id_index
    on bookings (category_id);

create index bookings_date_index
    on bookings (date);

create unique index bookings_id_uindex
    on bookings (id);

alter table bookings
    add constraint bookings_pk
        primary key (id);

