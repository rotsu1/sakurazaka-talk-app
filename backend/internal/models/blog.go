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

type BlogWithAuthor struct {
	ID        int       `json:"id"`
	MemberID  int       `json:"member_id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
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

func FindBlogByID(db *sql.DB, id int, status string) (*BlogWithAuthor, error) {
	var b BlogWithAuthor
	query := `
	SELECT 
			b.id, 
			b.member_id, 
			m.name,
			b.title, 
			b.content, 
			b.status, 
			b.created_at
	FROM blog b
	JOIN member m ON m.id = b.member_id
	WHERE b.id = $1`
	if status == "verified" {
		query += ` AND b.status = 'verified'`
	} else if status == "pending" {
		query += ` AND b.status = 'pending'`
	}
	err := db.QueryRow(query, id).Scan(
		&b.ID,
		&b.MemberID,
		&b.Author,
		&b.Title,
		&b.Content,
		&b.Status,
		&b.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetAllBlogs(db *sql.DB, status string) ([]*BlogWithAuthor, error) {
	query := `
	SELECT 
			b.id, 
			b.member_id, 
			m.name,
			b.title, 
			b.content, 
			b.status, 
			b.created_at
	FROM blog b
	JOIN member m ON m.id = b.member_id`

	if status == "verified" {
		query += ` WHERE status = 'verified'`
	} else if status == "pending" {
		query += ` WHERE status = 'pending'`
	}

	query += ` ORDER BY created_at DESC`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogsWithAuthor []*BlogWithAuthor
	for rows.Next() {
		var b BlogWithAuthor
		if err := rows.Scan(
			&b.ID,
			&b.MemberID,
			&b.Author,
			&b.Title,
			&b.Content,
			&b.Status,
			&b.CreatedAt,
		); err != nil {
			return nil, err
		}
		blogsWithAuthor = append(blogsWithAuthor, &b)
	}
	return blogsWithAuthor, nil
}
