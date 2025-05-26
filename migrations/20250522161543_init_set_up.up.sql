-- This SQL script sets up the initial database schema for the news aggregator application.
-- 1. Enable the pgcrypto extension for UUID generation
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 2. Create the 'sources' table
CREATE TABLE sources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE -- Source name should be unique and not null
);

-- 3. Create the 'authors' table
CREATE TABLE authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE -- Author name should be unique and not null
);

-- 4. Create the 'articles' table
CREATE TABLE articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_id UUID NOT NULL REFERENCES sources(id) ON DELETE CASCADE, -- Delete articles if source is deleted
    author_id UUID NOT NULL REFERENCES authors(id) ON DELETE CASCADE, -- Delete articles if author is deleted
    title VARCHAR(512) NOT NULL,
    description TEXT,
    url TEXT UNIQUE NOT NULL,
    url_to_image TEXT,
    published_at TIMESTAMPTZ NOT NULL,
    content TEXT
);

-- 5. Optional: Add an index on published_at for sorting
CREATE INDEX idx_articles_published_at ON articles (published_at);

-- 6. Optional: Add a full-text index on title for search
CREATE INDEX idx_articles_title ON articles USING GIN (to_tsvector('english', title));
