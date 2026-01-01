--- +up
CREATE TABLE message_read (
    message_id INT REFERENCES message(id) ON DELETE CASCADE,
    user_id INT REFERENCES talk_user(id) ON DELETE CASCADE,
    read_at TIMESTAMP NOT NULL,
    PRIMARY KEY (message_id, user_id)
);

--- +down
DROP TABLE message_read;