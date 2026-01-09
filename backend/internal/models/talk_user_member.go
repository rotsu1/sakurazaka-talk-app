package models

import (
	"database/sql"
	"time"
)

type TalkUserMember struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	MemberID  int        `json:"member_id"`
	Status    string     `json:"status"` // Status of subscription (e.g., "active", "cancelled")
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (t *TalkUserMember) Save(db *sql.DB) error {
	query := `
	INSERT INTO talk_user_member (
			user_id, 
			member_id, 
			status,
			expires_at
	) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, created_at, updated_at`

	return db.QueryRow(
		query,
		t.UserID,
		t.MemberID,
		t.Status,
		t.ExpiresAt, // driver usually handles *time.Time -> NULL if nil
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (t *TalkUserMember) Update(db *sql.DB) error {
	query := `
	UPDATE talk_user_member SET 
			user_id = $1, 
			member_id = $2, 
			status = $3, 
			expires_at = $4,
			updated_at = NOW() 
	WHERE id = $5`
	_, err := db.Exec(
		query,
		t.UserID,
		t.MemberID,
		t.Status,
		t.ExpiresAt,
		t.ID,
	)
	return err
}

func FindTalkUserMemberByID(db *sql.DB, id int) (*TalkUserMember, error) {
	var t TalkUserMember
	var expiresAt sql.NullTime
	query := `
	SELECT 
			id, 
			user_id, 
			member_id, 
			status, 
			expires_at,
			created_at, 
			updated_at 
	FROM talk_user_member 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&t.ID,
		&t.UserID,
		&t.MemberID,
		&t.Status,
		&expiresAt,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if expiresAt.Valid {
		t.ExpiresAt = &expiresAt.Time
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
			expires_at,
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
		var expiresAt sql.NullTime
		if err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.MemberID,
			&t.Status,
			&expiresAt,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		if expiresAt.Valid {
			t.ExpiresAt = &expiresAt.Time
		}
		members = append(members, &t)
	}
	return members, nil
}
