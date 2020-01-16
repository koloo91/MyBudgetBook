alter table accounts
    add user_id varchar(36);

alter table bookings
    add user_id varchar(36);

alter table categories
    add user_id varchar(36);

