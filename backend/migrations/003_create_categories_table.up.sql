create table categories
(
    id        varchar(36)             not null,
    parent_id varchar(36),
    name      varchar                 not null,
    created   timestamp default now() not null,
    updated   timestamp default now() not null
);

create unique index categories_id_uindex
    on categories (id);

create unique index categories_name_uindex
    on categories (lower(name));

alter table categories
    add constraint categories_pk
        primary key (id);
