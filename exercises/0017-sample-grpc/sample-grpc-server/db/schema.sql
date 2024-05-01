create table if not exists todos (
    id integer not null primary key,
    description text not null,
    done boolean not null default false,
    created timestamp not null default CURRENT_TIMESTAMP,
    updated timestamp not null default CURRENT_TIMESTAMP
);
