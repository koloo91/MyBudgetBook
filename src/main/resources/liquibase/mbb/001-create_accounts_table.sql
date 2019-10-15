create table accounts
(
    id varchar(36) not null,
    name varchar(55) not null,
    created timestamp default now() not null,
    updated timestamp default now() not null
);

create unique index accounts_id_uindex
    on accounts (id);

create unique index accounts_name_uindex
    on accounts (name);

alter table accounts
    add constraint accounts_pk
        primary key (id);

