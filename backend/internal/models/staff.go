package models

import (
	"database/sql"
	"time"
)

type Staff struct {
	ID           int
	MemberID     *int // Optional ID of the member this staff is associated with (e.g., manager)
	Role         string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *Staff) Save(db *sql.DB) error {
	query := `
	INSERT INTO staff (
			member_id, 
			role, 
			username, 
			password_hash
	) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, created_at, updated_at`
	return db.QueryRow(
		query,
		s.MemberID,
		s.Role,
		s.Username,
		s.PasswordHash,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

func (s *Staff) Update(db *sql.DB) error {
	query := `
	UPDATE staff SET 
			member_id = $1, 
			role = $2, 
			username = $3, 
			password_hash = $4, 
			updated_at = NOW() 
	WHERE id = $5`
	_, err := db.Exec(
		query,
		s.MemberID,
		s.Role,
		s.Username,
		s.PasswordHash,
		s.ID,
	)
	return err
}

func FindStaffByID(db *sql.DB, id int) (*Staff, error) {
	var s Staff
	query := `
	SELECT 
			id, 
			member_id, 
			role, 
			username, 
			password_hash, 
			created_at, 
			updated_at 
	FROM staff 
	WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&s.ID,
		&s.MemberID,
		&s.Role,
		&s.Username,
		&s.PasswordHash,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func GetAllStaff(db *sql.DB) ([]*Staff, error) {
	query := `
	SELECT 
			id, 
			member_id, 
			role, 
			username, 
			password_hash, 
			created_at, 
			updated_at 
	FROM staff`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffs []*Staff
	for rows.Next() {
		var s Staff
		if err := rows.Scan(
			&s.ID,
			&s.MemberID,
			&s.Role,
			&s.Username,
			&s.PasswordHash,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		staffs = append(staffs, &s)
	}
	return staffs, nil
}
