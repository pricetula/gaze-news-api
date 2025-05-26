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

INSERT INTO countries (id, name) VALUES
    ('ae', 'United Arab Emirates'),
    ('ar', 'Argentina'),
    ('at', 'Austria'),
    ('au', 'Australia'),
    ('be', 'Belgium'),
    ('bg', 'Bulgaria'),
    ('br', 'Brazil'),
    ('ca', 'Canada'),
    ('ch', 'Switzerland'),
    ('cn', 'China'),
    ('co', 'Colombia'),
    ('cu', 'Cuba'),
    ('cz', 'Czechia'),
    ('de', 'Germany'),
    ('eg', 'Egypt'),
    ('es', 'Spain'),
    ('fr', 'France'),
    ('gb', 'United Kingdom'),
    ('gr', 'Greece'),
    ('hk', 'Hong Kong'),
    ('hu', 'Hungary'),
    ('id', 'Indonesia'),
    ('ie', 'Ireland'),
    ('il', 'Israel'),
    ('in', 'India'),
    ('it', 'Italy'),
    ('jp', 'Japan'),
    ('kr', 'South Korea'),
    ('lt', 'Lithuania'),
    ('lv', 'Latvia'),
    ('ma', 'Morocco'),
    ('mx', 'Mexico'),
    ('my', 'Malaysia'),
    ('ng', 'Nigeria'),
    ('nl', 'Netherlands'),
    ('no', 'Norway'),
    ('nz', 'New Zealand'),
    ('ph', 'Philippines'),
    ('pl', 'Poland'),
    ('pt', 'Portugal'),
    ('ro', 'Romania'),
    ('rs', 'Serbia'),
    ('ru', 'Russia'),
    ('sa', 'Saudi Arabia'),
    ('se', 'Sweden'),
    ('sg', 'Singapore'),
    ('si', 'Slovenia'),
    ('sk', 'Slovakia'),
    ('th', 'Thailand'),
    ('tr', 'Turkey'),
    ('tw', 'Taiwan'),
    ('ua', 'Ukraine'),
    ('us', 'United States'),
    ('ve', 'Venezuela'),
    ('za', 'South Africa');

-- 2. Create the 'languages' table
CREATE TABLE languages (
    id VARCHAR(255) PRIMARY KEY, -- Language ID should be a string
    name VARCHAR(255) NOT NULL UNIQUE -- Language name should be unique and not null
);

INSERT INTO languages (id, name) VALUES
    ('ar', 'Arabic'),
    ('de', 'German'),
    ('en', 'English'),
    ('es', 'Spanish'),
    ('fr', 'French'),
    ('it', 'Italian'),
    ('nl', 'Dutch'),
    ('no', 'Norwegian'),
    ('pt', 'Portuguese'),
    ('ru', 'Russian'),
    ('se', 'Swedish');

-- 3. Create the 'categories' table
CREATE TABLE categories (
    id VARCHAR(255) PRIMARY KEY, -- Category ID should be a string
    name VARCHAR(255) NOT NULL UNIQUE -- Category name should be unique and not null
);

INSERT INTO categories (id, name) VALUES
    ('business', 'Business'),
    ('entertainment', 'Entertainment'),
    ('general', 'General'),
    ('health', 'Health'),
    ('science', 'Science'),
    ('sports', 'Sports'),
    ('technology', 'Technology');

-- 4. Alter the 'articles' table to change the source_id field to string
ALTER TABLE articles
    DROP CONSTRAINT articles_source_id_fkey,
    DROP CONSTRAINT articles_author_id_fkey,
    ALTER COLUMN source_id TYPE VARCHAR(255) USING source_id::VARCHAR(255),
    ALTER COLUMN author_id TYPE VARCHAR(255) USING author_id::VARCHAR(255); -- Change UUID to string

-- 5. Alter the 'sources' table to change the id field to string
ALTER TABLE sources
    ALTER COLUMN id TYPE VARCHAR(255) USING id::VARCHAR(255), -- Change UUID to string
    ADD COLUMN description TEXT, -- Description of the source
    ADD COLUMN url TEXT UNIQUE NOT NULL, -- URL of the source
    ADD COLUMN category_id VARCHAR(255) REFERENCES categories(id) ON DELETE SET NULL, -- Foreign key to categories table
    ADD COLUMN language_id VARCHAR(255) REFERENCES languages(id) ON DELETE SET NULL, -- Foreign key to languages table
    ADD COLUMN country_id VARCHAR(255) REFERENCES countries(id) ON DELETE SET NULL; -- Foreign key to countries table

-- 6. Alter the 'articles' table by adding foreign key constraint to 'sources' table
ALTER TABLE articles
    ADD CONSTRAINT articles_source_id_fkey
        FOREIGN KEY (source_id)
        REFERENCES sources (id);

-- 7. Alter the 'authors' table to change the id field to string
ALTER TABLE authors
    ALTER COLUMN id TYPE VARCHAR(255) USING id::VARCHAR(255); -- Change UUID to string

-- 8. Alter the 'articles' table by adding foreign key constraint to 'authors' table
ALTER TABLE articles
    ADD CONSTRAINT articles_author_id_fkey
        FOREIGN KEY (author_id)
        REFERENCES authors (id);