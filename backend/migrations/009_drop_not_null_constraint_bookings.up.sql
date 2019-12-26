alter table bookings
    alter column title drop not null;

alter table bookings
    alter column category_id drop not null;

alter table bookings
    alter column account_id drop not null;
