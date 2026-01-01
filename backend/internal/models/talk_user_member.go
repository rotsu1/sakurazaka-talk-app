package models

import (
	"database/sql"
	"time"
)

type TalkUserMember struct {
	ID        int
	UserID    int
	MemberID  int
	Status    string // Status of subscription (e.g., "active", "cancelled")
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *TalkUserMember) Save(db *sql.DB) error {
	query := `
	INSERT INTO talk_user_member (
			user_id, 
			member_id, 
			status
	) 
	VALUES ($1, $2, $3) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		t.UserID,
		t.MemberID,
		t.Status,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (t *TalkUserMember) Update(db *sql.DB) error {
	query := `
	UPDATE talk_user_member SET 
			user_id = $1, 
			member_id = $2, 
			status = $3, 
			updated_at = NOW() 
	WHERE id = $4`
	_, err := db.Exec(
		query,
		t.UserID,
		t.MemberID,
		t.Status,
		t.ID,
	)
	return err
}

func FindTalkUserMemberByID(db *sql.DB, id int) (*TalkUserMember, error) {
	var t TalkUserMember
	query := `
	SELECT 
			id, 
			user_id, 
			member_id, 
			status, 
			created_at, 
			updated_at 
	FROM talk_user_member 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&t.ID,
		&t.UserID,
		&t.MemberID,
		&t.Status,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetAllTalkUserMembers(db *sql.DB) ([]*TalkUserMember, error) {
	query := `
	SELECT 
			id, 
			user_id, 
			member_id, 
			status, 
			created_at, 
			updated_at 
	FROM talk_user_member`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*TalkUserMember
	for rows.Next() {
		var t TalkUserMember
		if err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.MemberID,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		members = append(members, &t)
	}
	return members, nil
}
