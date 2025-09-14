-- +goose up

CREATE TABLE tags (
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      name VARCHAR(30) NOT NULL UNIQUE,
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_tags (
                           user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                           tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
                           PRIMARY KEY (user_id, tag_id)
);

-- Pre-populate with common tags
INSERT INTO tags (name) VALUES
                            ('actor'), ('director'), ('athlete'), ('musician'), ('entrepreneur'),
                            ('comedian'), ('author'), ('scientist'), ('model'), ('chef'),
                            ('youtuber'), ('tiktoker'), ('podcast_host'), ('activist');

-- +goose down
DROP TABLE user_tags;
DROP TABLE tags;