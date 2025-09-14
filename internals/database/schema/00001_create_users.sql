-- +goose up

-- Create the enum type first
CREATE TYPE celebrity_domain AS ENUM (
    'sports',
    'cinema',
    'politics',
    'music',
    'business',
    'influencer',
    'comedy',
    'literature',
    'science',
    'technology',
    'art',
    'fashion',
    'gaming',
    'travel',
    'food',
    'health',
    'education',
    'environment',
    'history',
    'philosophy',
    'religion',
    'activism',
    'journalism',
    'photography',
    'theater',
    'other'
);

CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       bio TEXT,
                       profile_image_url TEXT,
                       domain celebrity_domain NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose down
DROP TABLE users;
DROP TYPE celebrity_domain;