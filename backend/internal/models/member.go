package models

import (
	"database/sql"
	"time"
)

type Member struct {
	ID         int
	Name       string
	Generation *int
	AvatarURL  *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (m *Member) Save(db *sql.DB) error {
	query := `
	INSERT INTO member (
			name, 
			generation, 
			avatar_url
	) 
	VALUES ($1, $2, $3) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		m.Name,
		m.Generation,
		m.AvatarURL,
	).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
}

func (m *Member) Update(db *sql.DB) error {
	query := `
	UPDATE member SET 
			name = $1, 
			generation = $2, 
			avatar_url = $3, 
			updated_at = NOW() 
	WHERE id = $4`
	_, err := db.Exec(
		query,
		m.Name,
		m.Generation,
		m.AvatarURL,
		m.ID,
	)
	return err
}

func FindMemberByID(db *sql.DB, id int) (*Member, error) {
	var m Member
	query := `
	SELECT 
			id, 
			name, 
			generation, 
			avatar_url, 
			created_at, 
			updated_at 
	FROM member 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&m.ID,
		&m.Name,
		&m.Generation,
		&m.AvatarURL,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func GetAllMembers(db *sql.DB) ([]*Member, error) {
	query := `
	SELECT 
			id, 
			name, 
			generation, 
			avatar_url, 
			created_at, 
			updated_at 
	FROM member`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*Member
	for rows.Next() {
		var m Member
		if err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Generation,
			&m.AvatarURL,
			&m.CreatedAt,
			&m.UpdatedAt,
		); err != nil {
			return nil, err
		}
		members = append(members, &m)
	}
	return members, nil
}
