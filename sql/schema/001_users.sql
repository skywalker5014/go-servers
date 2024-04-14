-- +goose up

CREATE TABLE users (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null
);

-- +goose down

drop table users;