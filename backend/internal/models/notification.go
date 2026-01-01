package models

import (
	"database/sql"
	"time"
)

type Notification struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n *Notification) Save(db *sql.DB) error {
	query := `
	INSERT INTO notification (
			title, 
			content
	) 
	VALUES ($1, $2) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		n.Title,
		n.Content,
	).Scan(&n.ID, &n.CreatedAt, &n.UpdatedAt)
}

func (n *Notification) Update(db *sql.DB) error {
	query := `
	UPDATE notification SET 
			title = $1, 
			content = $2, 
			updated_at = NOW() 
	WHERE id = $3`
	_, err := db.Exec(
		query,
		n.Title,
		n.Content,
		n.ID,
	)
	return err
}

func FindNotificationByID(db *sql.DB, id int) (*Notification, error) {
	var n Notification
	query := `
	SELECT 
			id, 
			title, 
			content, 
			created_at, 
			updated_at 
	FROM notification 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&n.ID,
		&n.Title,
		&n.Content,
		&n.CreatedAt,
		&n.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func GetAllNotifications(db *sql.DB) ([]*Notification, error) {
	query := `
	SELECT 
			id, 
			title, 
			content, 
			created_at, 
			updated_at 
	FROM notification`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		var n Notification
		if err := rows.Scan(
			&n.ID,
			&n.Title,
			&n.Content,
			&n.CreatedAt,
			&n.UpdatedAt,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, &n)
	}
	return notifications, nil
}
