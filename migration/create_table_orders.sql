drop table if exists orders;

create table orders (
    seq bigint UNSIGNED AUTO_INCREMENT UNIQUE,
    id bigint primary key,
    price decimal(10,2) not null,
    tax decimal(10,2) not null,
    final_price decimal (10,2) not null,
    created_at timestamp(6) not null
);
