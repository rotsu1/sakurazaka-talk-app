package models

import (
	"database/sql"
	"time"
)

type Message struct {
	ID         int
	MemberID   int
	Type       string // Type of the message (e.g., "text", "image", "video")
	Content    string
	Status     string // Status of the message (e.g., "pending", "approved", "rejected")
	VerifiedBy *int   // ID of the staff who verified the message
	VerifiedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (m *Message) Save(db *sql.DB) error {
	query := `
	INSERT INTO message (
			member_id, 
			type, 
			content, 
			status, 
			verified_by, 
			verified_at
	) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		m.MemberID,
		m.Type,
		m.Content,
		m.Status,
		m.VerifiedBy,
		m.VerifiedAt,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
}

func (m *Message) Update(db *sql.DB) error {
	query := `
	UPDATE message SET 
			member_id = $1, 
			type = $2, 
			content = $3, 
			status = $4, 
			verified_by = $5, 
			verified_at = $6, 
			updated_at = NOW() 
	WHERE id = $7`
	_, err := db.Exec(
		query,
		m.MemberID,
		m.Type,
		m.Content,
		m.Status,
		m.VerifiedBy,
		m.VerifiedAt,
		m.ID,
	)
	return err
}

func FindMessageByID(db *sql.DB, id int) (*Message, error) {
	var m Message
	query := `
	SELECT 
			id, 
			member_id, 
			type, 
			content, 
			status, 
			verified_by, 
			verified_at, 
			created_at, 
			updated_at 
	FROM message 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&m.ID,
		&m.MemberID,
		&m.Type,
		&m.Content,
		&m.Status,
		&m.VerifiedBy,
		&m.VerifiedAt,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func GetAllMessages(db *sql.DB) ([]*Message, error) {
	query := `
	SELECT 
			id, 
			member_id, 
			type, 
			content, 
			status, 
			verified_by, 
			verified_at, 
			created_at, 
			updated_at 
	FROM message`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(
			&m.ID,
			&m.MemberID,
			&m.Type,
			&m.Content,
			&m.Status,
			&m.VerifiedBy,
			&m.VerifiedAt,
			&m.CreatedAt,
			&m.UpdatedAt,
		); err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	return messages, nil
}
