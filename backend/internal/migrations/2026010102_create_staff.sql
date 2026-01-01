--- +up
CREATE TABLE staff (
    id SERIAL PRIMARY KEY,
    member_id INT REFERENCES member(id), -- null if staff is a manager only
    role VARCHAR(20) NOT NULL,           -- 'member' or 'manager'
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE staff;

