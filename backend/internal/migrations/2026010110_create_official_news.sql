--- +up
CREATE TABLE official_news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    tag VARCHAR(50),
    content TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE official_news;