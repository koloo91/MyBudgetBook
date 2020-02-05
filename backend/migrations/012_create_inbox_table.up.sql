create table inbox
(
    id           varchar(36)                not null,
    user_id      varchar(36)                not null,
    booking_date timestamptz                not null,
    value_date   timestamptz,
    intended_use varchar                    not null,
    amount       numeric(10, 2) default 0.0 not null,
    created      timestamptz,
    updated      timestamptz
);

create unique index inbox_id_uindex
    on inbox (id);

alter table inbox
    add constraint inbox_pk
        primary key (id);

create index inbox_id_user_id_index
    on inbox (id, user_id);

create index inbox_user_id_index
    on inbox (user_id);

