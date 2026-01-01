package models

import (
	"database/sql"
	"time"
)

type Template struct {
	ID          int
	TemplateURL string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Template) Save(db *sql.DB) error {
	query := `
	INSERT INTO template (
			template_url
	) 
	VALUES ($1) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		t.TemplateURL,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (t *Template) Update(db *sql.DB) error {
	query := `
	UPDATE template SET 
			template_url = $1, 
			updated_at = NOW() 
	WHERE id = $2`
	_, err := db.Exec(
		query,
		t.TemplateURL,
		t.ID,
	)
	return err
}

func FindTemplateByID(db *sql.DB, id int) (*Template, error) {
	var t Template
	query := `
	SELECT 
			id, 
			template_url, 
			created_at, 
			updated_at 
	FROM template 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&t.ID,
		&t.TemplateURL,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetAllTemplates(db *sql.DB) ([]*Template, error) {
	query := `
	SELECT 
			id, 
			template_url, 
			created_at, 
			updated_at 
	FROM template`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*Template
	for rows.Next() {
		var t Template
		if err := rows.Scan(
			&t.ID,
			&t.TemplateURL,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		templates = append(templates, &t)
	}
	return templates, nil
}
