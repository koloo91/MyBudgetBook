alter table accounts
    alter column created type timestamptz using created::timestamptz;

alter table accounts
    alter column updated type timestamptz using updated::timestamptz;


alter table categories
    alter column created type timestamptz using created::timestamptz;

alter table categories
    alter column updated type timestamptz using updated::timestamptz;



alter table bookings
    alter column date type timestamptz using date::timestamptz;

alter table bookings
    alter column created type timestamptz using created::timestamptz;

alter table bookings
    alter column updated type timestamptz using updated::timestamptz;

