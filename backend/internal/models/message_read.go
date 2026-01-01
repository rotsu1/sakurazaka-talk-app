package models

import (
	"database/sql"
	"time"
)

type MessageRead struct {
	MessageID int
	UserID    int
	ReadAt    time.Time
}

func (mr *MessageRead) Save(db *sql.DB) error {
	query := `
	INSERT INTO message_read (
			message_id, 
			user_id, 
			read_at
	) 
	VALUES ($1, $2, $3)`
	_, err := db.Exec(
		query,
		mr.MessageID,
		mr.UserID,
		mr.ReadAt,
	)
	return err
}

func (mr *MessageRead) Delete(db *sql.DB) error {
	query := `
	DELETE FROM message_read 
	WHERE message_id = $1 AND user_id = $2`
	_, err := db.Exec(query, mr.MessageID, mr.UserID)
	return err
}

func FindMessageReadByIDs(db *sql.DB, messageID, userID int) (*MessageRead, error) {
	var mr MessageRead
	query := `
	SELECT 
			message_id, 
			user_id, 
			read_at 
	FROM message_read 
	WHERE message_id = $1 AND user_id = $2`
	err := db.QueryRow(query, messageID, userID).Scan(
		&mr.MessageID,
		&mr.UserID,
		&mr.ReadAt,
	)
	if err != nil {
		return nil, err
	}
	return &mr, nil
}

func GetAllMessageReads(db *sql.DB) ([]*MessageRead, error) {
	query := `
	SELECT 
			message_id, 
			user_id, 
			read_at 
	FROM message_read`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messageReads []*MessageRead
	for rows.Next() {
		var mr MessageRead
		if err := rows.Scan(
			&mr.MessageID,
			&mr.UserID,
			&mr.ReadAt,
		); err != nil {
			return nil, err
		}
		messageReads = append(messageReads, &mr)
	}
	return messageReads, nil
}
