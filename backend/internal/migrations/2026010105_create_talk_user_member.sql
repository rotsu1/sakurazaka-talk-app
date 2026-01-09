--- +up
CREATE TABLE talk_user_member (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES talk_user(id) ON DELETE CASCADE,
    member_id INT NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'subscribed', -- 'subscribed'/'unsubscribed'
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

--- +down
DROP TABLE talk_user_member;

