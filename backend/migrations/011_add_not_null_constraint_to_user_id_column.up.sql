alter table accounts
    alter column user_id set not null;

drop index accounts_name_idx;

create unique index accounts_name_idx
    on accounts (user_id, lower(name::text));


alter table categories
    alter column user_id set not null;

drop index categories_name_uindex;

create unique index categories_name_uindex
    on categories (user_id, lower(name::text));

alter table bookings
    alter column user_id set not null;

drop index bookings_account_id_index;

create index bookings_user_id_account_id_index
    on bookings (user_id, account_id);

drop index bookings_category_id_index;

create index bookings_user_id_category_id_index
    on bookings (user_id, category_id);

drop index bookings_date_index;

create index bookings_user_id_date_index
    on bookings (user_id, date);

drop index bookings_category_id_date_index;

create index bookings_user_id_category_id_date_index
    on bookings (user_id, category_id, date);

drop index bookings_account_id_date_index;

create index bookings_user_id_account_id_date_index
    on bookings (user_id, account_id, date);

