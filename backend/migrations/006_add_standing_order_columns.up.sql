alter table bookings
    add standing_order_id varchar(36);

alter table bookings
    add standing_order_period varchar;

alter table bookings
    add standing_order_last_day timestamptz;

