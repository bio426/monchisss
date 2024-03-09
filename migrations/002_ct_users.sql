create table stores(
    id serial primary key,
    name varchar unique,
    wa_token varchar,
  	active bool default true,
    created_at timestamp default now(),
 	admin_id integer references users
);

---- create above / drop below ----

drop table stores;
