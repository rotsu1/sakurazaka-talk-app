--- +up
CREATE TABLE member (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    generation INT,
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE member