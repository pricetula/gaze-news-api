-- This migration script reverts the changes made by ('20250523090312_add_countries.up.sql').
-- It safely drops tables, indexes, and extensions in the correct order to avoid dependency errors.

-- 1. Alter the 'articles' table to change the id field back to UUID
ALTER TABLE articles
    DROP CONSTRAINT articles_source_id_fkey;
ALTER TABLE articles
    ALTER COLUMN source_id TYPE UUID USING source_id::UUID; -- Change string back to UUID
ALTER TABLE articles
    ADD CONSTRAINT articles_source_id_fkey
        FOREIGN KEY (source_id)
        REFERENCES sources (id);

-- 2. Alter the 'authors' table to change the id field back to UUID
ALTER TABLE authors
    ALTER COLUMN id TYPE UUID USING id::UUID; -- Change string back to UUID

-- 3. Alter the 'sources' table to change the id field back to UUID
ALTER TABLE sources
    ALTER COLUMN id TYPE UUID USING id::UUID; -- Change string back to UUID

-- 4. Drop the 'sources' table added fields
ALTER TABLE sources
    DROP COLUMN IF EXISTS description, -- Drop description of the source
    DROP COLUMN IF EXISTS url, -- Drop URL of the source
    DROP COLUMN IF EXISTS category_id, -- Drop foreign key to categories table
    DROP COLUMN IF EXISTS language_id, -- Drop foreign key to languages table
    DROP COLUMN IF EXISTS country_id; -- Drop foreign key to countries table

-- 5. Drop the 'categories' table
DROP TABLE IF EXISTS categories;

-- 6. Drop the 'languages' table
DROP TABLE IF EXISTS languages;

-- 7. Drop the 'countries' table
DROP TABLE IF EXISTS countries;