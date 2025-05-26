-- This migration drops the unique constraint on the source_url column in the articles table.
ALTER TABLE sources
    DROP CONSTRAINT sources_url_key;