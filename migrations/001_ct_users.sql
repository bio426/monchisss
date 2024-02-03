create type user_role as enum('super','owner','employee');
create table users(
    id serial primary key,
    username varchar unique,
    password varchar,
    role user_role,
    active bool default true,
    created_at timestamp default now()
);

---- create above / drop below ----

drop table users;
drop type user_role;
