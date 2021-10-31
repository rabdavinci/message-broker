CREATE TABLE IF NOT EXISTS users (
    id uuid primary key,
    name varchar not null,
    created_at timestamp default current_timestamp,
	updated_at timestamp,
	deleted_at bigint default 0
);