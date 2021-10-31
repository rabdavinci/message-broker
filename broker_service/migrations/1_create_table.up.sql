CREATE TABLE IF NOT EXISTS broker (
    user_id uuid,
    topic_id uuid,
    created_at timestamp default current_timestamp,
	updated_at timestamp,
	deleted_at bigint default 0,
    primary key (user_id, topic_id)
);