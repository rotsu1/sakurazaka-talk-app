--- +up
CREATE TABLE fanletter (
    id SERIAL PRIMARY KEY,
    member_id INT NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    talk_user_id INT NOT NULL REFERENCES talk_user(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    template_id INT REFERENCES template(id),
    created_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE fanletter;

