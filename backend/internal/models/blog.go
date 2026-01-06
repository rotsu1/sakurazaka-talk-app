package models

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID         int        `json:"id"`
	MemberID   int        `json:"member_id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Status     string     `json:"status"`
	VerifiedBy *int       `json:"verified_by"` // Verfied by staff with manager role
	VerifiedAt *time.Time `json:"verified_at"` // Verfied time by staff with manager role
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (b *Blog) Save(db *sql.DB) error {
	query := `
	INSERT INTO blog (
			member_id, 
			title, 
			content, 
			status, 
			verified_by, 
			verified_at
	)
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		b.MemberID,
		b.Title,
		b.Content,
		b.Status,
		b.VerifiedBy,
		b.VerifiedAt,
	).Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt)
}

func (b *Blog) Update(db *sql.DB) error {
	query := `
	UPDATE blog SET 
			member_id = $1, 
			title = $2, 
			content = $3, 
			status = $4, 
			verified_by = $5, 
			verified_at = $6, 
			updated_at = NOW() 
	WHERE id = $7`
	result, err := db.Exec(
		query,
		b.MemberID,
		b.Title,
		b.Content,
		b.Status,
		b.VerifiedBy,
		b.VerifiedAt,
		b.ID,
	)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func FindBlogByID(db *sql.DB, id int) (*Blog, error) {
	var b Blog
	query := `
	SELECT 
			id, 
			member_id, 
			title, 
			content, 
			status, 
			verified_by, 
			verified_at, 
			created_at, 
			updated_at 
	FROM blog
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&b.ID,
		&b.MemberID,
		&b.Title,
		&b.Content,
		&b.Status,
		&b.VerifiedBy,
		&b.VerifiedAt,
		&b.CreatedAt,
		&b.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetAllBlogs(db *sql.DB) ([]*Blog, error) {
	query := `
	SELECT 
			id, 
			member_id, 
			title, 
			content, 
			status, 
			verified_by, 
			verified_at, 
			created_at, 
			updated_at 
	FROM blog`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*Blog
	for rows.Next() {
		var b Blog
		if err := rows.Scan(
			&b.ID,
			&b.MemberID,
			&b.Title,
			&b.Content,
			&b.Status,
			&b.VerifiedBy,
			&b.VerifiedAt,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		blogs = append(blogs, &b)
	}
	return blogs, nil
}
