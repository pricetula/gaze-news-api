-- This migration reverts changes made on '20250526195751_drop_source_url_unique.up.sql' by adding back the unique constraint on the source_url column in the articles table.
ALTER TABLE sources
    ADD CONSTRAINT sources_url_key UNIQUE (url);