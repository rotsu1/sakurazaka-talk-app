--- +up
CREATE TABLE template (
    id SERIAL PRIMARY KEY,
    template_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE template;

