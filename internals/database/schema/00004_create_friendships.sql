-- +goose up

CREATE TABLE friendships (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    friend_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, friend_id)
);

-- +goose down
    DROP TABLE friendships;