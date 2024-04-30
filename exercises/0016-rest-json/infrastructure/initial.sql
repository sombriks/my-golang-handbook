-- initial database state
create table if not exists todos(
    id serial primary key,
    description text not null,
    done boolean not null default false,
    created timestamp not null default now(),
    updated timestamp not null default now()
);
