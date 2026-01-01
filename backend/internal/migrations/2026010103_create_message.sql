--- +up
CREATE TABLE message (
    id SERIAL PRIMARY KEY,
    member_id INT NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL, -- 'voice' or 'text'
    content TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending', -- 'pending', 'verified', 'rejected'
    verified_by INT REFERENCES staff(id),
    verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE message;

