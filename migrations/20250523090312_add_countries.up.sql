-- This migration adds a new table 'countries' to the database schema,
-- adds a new table 'languages' to the database schema,
-- adds a new table 'categories' to the database schema,
-- adds missing fields in sources table.
-- changes sources abd authors id field to string and
-- changes articles table source_id to string.

-- 1. Create the 'countries' table
CREATE TABLE countries (
    id VARCHAR(255) PRIMARY KEY, -- Country ID should be a string
    name VARCHAR(255) NOT NULL UNIQUE -- Country name should be unique and not null
);

-- 2. Create the 'languages' table
CREATE TABLE languages (
    id VARCHAR(255) PRIMARY KEY, -- Language ID should be a string
    name VARCHAR(255) NOT NULL UNIQUE -- Language name should be unique and not null
);

-- 3. Create the 'categories' table
CREATE TABLE categories (
    id VARCHAR(255) PRIMARY KEY, -- Category ID should be a string
    name VARCHAR(255) NOT NULL UNIQUE -- Category name should be unique and not null
);

-- 3a. Alter the 'articles' table to change the source_id field to string
ALTER TABLE articles
    DROP CONSTRAINT articles_source_id_fkey,
    DROP CONSTRAINT articles_author_id_fkey,
    ALTER COLUMN source_id TYPE VARCHAR(255) USING source_id::VARCHAR(255),
    ALTER COLUMN author_id TYPE VARCHAR(255) USING author_id::VARCHAR(255); -- Change UUID to string

-- 4. Alter the 'sources' table to change the id field to string
ALTER TABLE sources
    ALTER COLUMN id TYPE VARCHAR(255) USING id::VARCHAR(255), -- Change UUID to string
    ADD COLUMN description TEXT, -- Description of the source
    ADD COLUMN url TEXT UNIQUE NOT NULL, -- URL of the source
    ADD COLUMN category_id VARCHAR(255) REFERENCES categories(id) ON DELETE SET NULL, -- Foreign key to categories table
    ADD COLUMN language_id VARCHAR(255) REFERENCES languages(id) ON DELETE SET NULL, -- Foreign key to languages table
    ADD COLUMN country_id VARCHAR(255) REFERENCES countries(id) ON DELETE SET NULL; -- Foreign key to countries table

ALTER TABLE articles
    ADD CONSTRAINT articles_source_id_fkey
        FOREIGN KEY (source_id)
        REFERENCES sources (id);

-- 6. Alter the 'authors' table to change the id field to string
ALTER TABLE authors
    ALTER COLUMN id TYPE VARCHAR(255) USING id::VARCHAR(255); -- Change UUID to string

ALTER TABLE articles
    ADD CONSTRAINT articles_author_id_fkey
        FOREIGN KEY (author_id)
        REFERENCES authors (id);