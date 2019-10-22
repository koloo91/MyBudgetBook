create table accounts
(
    id varchar(36) not null,
    name varchar(55) not null,
    created timestamp default now() not null,
    updated timestamp default now() not null
);

create unique index accounts_id_idx
    on accounts (id);

create unique index accounts_name_idx
    on accounts (lower(name));

alter table accounts
    add constraint accounts_pk
        primary key (id);

