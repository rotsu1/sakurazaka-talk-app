package models

import (
	"database/sql"
	"time"
)

type Fanletter struct {
	ID         int       `json:"id"`
	MemberID   int       `json:"member_id"`
	TalkUserID int       `json:"talk_user_id"`
	Content    string    `json:"content"`
	TemplateID *int      `json:"template_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (f *Fanletter) Save(db *sql.DB) error {
	query := `
	INSERT INTO fanletter (
			member_id, 
			talk_user_id, 
			content, 
			template_id
	) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, created_at`
	return db.QueryRow(
		query,
		f.MemberID,
		f.TalkUserID,
		f.Content,
		f.TemplateID,
	).Scan(&f.ID, &f.CreatedAt)
}

func (f *Fanletter) Update(db *sql.DB) error {
	query := `
	UPDATE fanletter SET 
			member_id = $1, 
			talk_user_id = $2, 
			content = $3, 
			template_id = $4 
	WHERE id = $5`
	_, err := db.Exec(
		query,
		f.MemberID,
		f.TalkUserID,
		f.Content,
		f.TemplateID,
		f.ID,
	)
	return err
}

func FindFanletterByID(db *sql.DB, id int) (*Fanletter, error) {
	var f Fanletter
	query := `
	SELECT 
			id, 
			member_id, 
			talk_user_id, 
			content, 
			template_id, 
			created_at 
	FROM fanletter 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&f.ID,
		&f.MemberID,
		&f.TalkUserID,
		&f.Content,
		&f.TemplateID,
		&f.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func GetAllFanletters(db *sql.DB) ([]*Fanletter, error) {
	query := `
	SELECT 
			id, 
			member_id, 
			talk_user_id, 
			content, 
			template_id, 
			created_at 
	FROM fanletter`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fanletters []*Fanletter
	for rows.Next() {
		var f Fanletter
		if err := rows.Scan(
			&f.ID,
			&f.MemberID,
			&f.TalkUserID,
			&f.Content,
			&f.TemplateID,
			&f.CreatedAt,
		); err != nil {
			return nil, err
		}
		fanletters = append(fanletters, &f)
	}
	return fanletters, nil
}
