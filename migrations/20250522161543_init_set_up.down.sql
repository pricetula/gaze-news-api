-- This migration script reverts the changes made by ('20250522161543_init_set_up.up.sql').
-- It safely drops tables, indexes, and extensions in the correct order to avoid dependency errors.

-- 1. Drop the 'articles' table first, as it has foreign keys referencing 'sources' and 'authors'.
-- CASCADE is implicit for the indexes when the table is dropped, but explicitly dropping indexes is good practice if they were created separately.
DROP TABLE IF EXISTS articles CASCADE;

-- 2. Drop the 'authors' table.
DROP TABLE IF EXISTS authors CASCADE;

-- 3. Drop the 'sources' table.
DROP TABLE IF EXISTS sources CASCADE;

-- 4. Drop the pgcrypto extension if it was created specifically for this application
-- This should only be done if you are certain no other parts of your database rely on it.
-- Using DROP EXTENSION IF EXISTS ensures it won't error if it's already gone or never existed.
DROP EXTENSION IF EXISTS pgcrypto;