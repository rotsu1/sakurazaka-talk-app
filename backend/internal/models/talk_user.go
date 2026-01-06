package models

import (
	"database/sql"
	"time"
)

type TalkUser struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *TalkUser) Save(db *sql.DB) error {
	query := `
	INSERT INTO talk_user 
	DEFAULT VALUES 
	RETURNING id, created_at, updated_at`
	// Assuming no other fields for now as per struct definition
	return db.QueryRow(query).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (t *TalkUser) Update(db *sql.DB) error {
	query := `
	UPDATE talk_user SET 
			updated_at = NOW() 
	WHERE id = $1`
	_, err := db.Exec(query, t.ID)
	return err
}

func FindTalkUserByID(db *sql.DB, id int) (*TalkUser, error) {
	var t TalkUser
	query := `
	SELECT 
			id, 
			created_at, 
			updated_at 
	FROM talk_user 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&t.ID,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetAllTalkUsers(db *sql.DB) ([]*TalkUser, error) {
	query := `
	SELECT 
			id, 
			created_at, 
			updated_at 
	FROM talk_user`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*TalkUser
	for rows.Next() {
		var t TalkUser
		if err := rows.Scan(
			&t.ID,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &t)
	}
	return users, nil
}
