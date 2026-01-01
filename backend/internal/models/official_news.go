package models

import (
	"database/sql"
	"time"
)

type OfficialNews struct {
	ID        int
	Title     string
	Tag       *string // Optional tag for the news (e.g., "Event", "Release")
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *OfficialNews) Save(db *sql.DB) error {
	query := `
	INSERT INTO official_news (
			title, 
			tag
	) 
	VALUES ($1, $2) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		o.Title,
		o.Tag,
	).Scan(&o.ID, &o.CreatedAt, &o.UpdatedAt)
}

func (o *OfficialNews) Update(db *sql.DB) error {
	query := `
	UPDATE official_news SET 
			title = $1, 
			tag = $2, 
			updated_at = NOW() 
	WHERE id = $3`
	_, err := db.Exec(
		query,
		o.Title,
		o.Tag,
		o.ID,
	)
	return err
}

func FindOfficialNewsByID(db *sql.DB, id int) (*OfficialNews, error) {
	var o OfficialNews
	query := `
	SELECT 
			id, 
			title, 
			tag, 
			created_at, 
			updated_at 
	FROM official_news 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&o.ID,
		&o.Title,
		&o.Tag,
		&o.CreatedAt,
		&o.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func GetAllOfficialNews(db *sql.DB) ([]*OfficialNews, error) {
	query := `
	SELECT 
			id, 
			title, 
			tag, 
			created_at, 
			updated_at 
	FROM official_news`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []*OfficialNews
	for rows.Next() {
		var o OfficialNews
		if err := rows.Scan(
			&o.ID,
			&o.Title,
			&o.Tag,
			&o.CreatedAt,
			&o.UpdatedAt,
		); err != nil {
			return nil, err
		}
		news = append(news, &o)
	}
	return news, nil
}
