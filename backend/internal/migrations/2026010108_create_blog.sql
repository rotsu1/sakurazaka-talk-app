--- +up
CREATE TABLE blog (
    id SERIAL PRIMARY KEY,
    member_id INT NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending', -- verification
    verified_by INT REFERENCES staff(id),
    verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE blog;

