package models_test

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
)

// setupTestDB
func setupTestDB(t *testing.T, tables string) *sql.DB {
	db, err := db.InitDB()
	if err != nil {
		log.Println("DB connection failed:", err)
	}
	_, err = db.Exec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE", tables))
	if err != nil {
		t.Fatal("Failed to truncate tables:", err)
	}

	return db
}

func assertNoError(t *testing.T, err error, msg string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %v", msg, err)
	}
}

func assertError(t *testing.T, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Fatal(msg)
	}
}

func assertCount(t *testing.T, got, want int, msg string) {
	t.Helper()
	if got != want {
		t.Fatalf("%s: expected %d, got %d", msg, want, got)
	}
}

func assertNotNil[T any](t *testing.T, v *T, msg string) {
	t.Helper()
	if v == nil {
		t.Fatal(msg)
	}
}

func assertNil[T any](t *testing.T, v *T, msg string) {
	t.Helper()
	if v != nil {
		t.Fatalf("%s: expected nil, got %v", msg, v)
	}
}

// CreateNewMember creates new member instance
func createNewMember(name string) *models.Member {
	gen := 1
	return &models.Member{
		Name:       name,
		Generation: &gen,
	}
}

func saveMember(t *testing.T, db *sql.DB, m *models.Member) {
	if err := m.Save(db); err != nil {
		t.Fatalf("Failed to save member: %v", err)
	}
}

// NewBlog creates new blog instance
func createNewBlog(
	memberID int,
	title, content, status string,
	opts ...func(*models.Blog),
) *models.Blog {
	b := &models.Blog{
		MemberID: memberID,
		Title:    title,
		Content:  content,
		Status:   status,
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func saveBlog(t *testing.T, db *sql.DB, b *models.Blog) {
	if err := b.Save(db); err != nil {
		t.Fatalf("Failed to save blog: %v", err)
	}
}

// Optional helper functions
func withVerifiedBy(name int) func(*models.Blog) {
	return func(b *models.Blog) {
		b.VerifiedBy = &name
	}
}
func withVerifiedAt(t time.Time) func(*models.Blog) {
	return func(b *models.Blog) {
		b.VerifiedAt = &t
	}
}

// createNewTalkUser creates new talk_user instance
func createNewTalkUser() *models.TalkUser {
	return &models.TalkUser{}
}

func saveTalkUser(t *testing.T, db *sql.DB, u *models.TalkUser) {
	if err := u.Save(db); err != nil {
		t.Fatalf("Failed to save talk_user: %v", err)
	}
}

// createNewFanletter creates new fanletter instance
func createNewFanletter(memberID, talkUserID int, content string) *models.Fanletter {
	return &models.Fanletter{
		MemberID:   memberID,
		TalkUserID: talkUserID,
		Content:    content,
	}
}

func saveFanletter(t *testing.T, db *sql.DB, f *models.Fanletter) {
	if err := f.Save(db); err != nil {
		t.Fatalf("Failed to save fanletter: %v", err)
	}
}

// createNewMessage creates new message instance
func createNewMessage(memberID int, content, msgType, status string) *models.Message {
	return &models.Message{
		MemberID: memberID,
		Content:  content,
		Type:     msgType,
		Status:   status,
	}
}

func saveMessage(t *testing.T, db *sql.DB, m *models.Message) {
	if err := m.Save(db); err != nil {
		t.Fatalf("Failed to save message: %v", err)
	}
}

// createNewNotification creates new notification instance
func createNewNotification(title, content string) *models.Notification {
	return &models.Notification{
		Title:   title,
		Content: content,
	}
}

func saveNotification(t *testing.T, db *sql.DB, n *models.Notification) {
	if err := n.Save(db); err != nil {
		t.Fatalf("Failed to save notification: %v", err)
	}
}

// createNewOfficialNews creates new official_news instance
func createNewOfficialNews(title string, tag *string) *models.OfficialNews {
	return &models.OfficialNews{
		Title: title,
		Tag:   tag,
	}
}

func saveOfficialNews(t *testing.T, db *sql.DB, n *models.OfficialNews) {
	if err := n.Save(db); err != nil {
		t.Fatalf("Failed to save official_news: %v", err)
	}
}

// createNewStaff creates new staff instance
func createNewStaff(memberID *int, role, username, password string) *models.Staff {
	return &models.Staff{
		MemberID:     memberID,
		Role:         role,
		Username:     username,
		PasswordHash: password,
	}
}

func saveStaff(t *testing.T, db *sql.DB, s *models.Staff) {
	if err := s.Save(db); err != nil {
		t.Fatalf("Failed to save staff: %v", err)
	}
}

// createNewTalkUserMember creates new talk_user_member instance
func createNewTalkUserMember(userID, memberID int, status string) *models.TalkUserMember {
	return &models.TalkUserMember{
		UserID:   userID,
		MemberID: memberID,
		Status:   status,
	}
}

func saveTalkUserMember(t *testing.T, db *sql.DB, tm *models.TalkUserMember) {
	if err := tm.Save(db); err != nil {
		t.Fatalf("Failed to save talk_user_member: %v", err)
	}
}

// createNewTemplate creates new template instance
func createNewTemplate(url string) *models.Template {
	return &models.Template{
		TemplateURL: url,
	}
}

func saveTemplate(t *testing.T, db *sql.DB, tp *models.Template) {
	if err := tp.Save(db); err != nil {
		t.Fatalf("Failed to save template: %v", err)
	}
}

// createNewMessageRead creates new message_read instance
func createNewMessageRead(messageID, userID int) *models.MessageRead {
	return &models.MessageRead{
		MessageID: messageID,
		UserID:    userID,
		ReadAt:    time.Now().Truncate(time.Second), // Truncate to avoid microsecond differences in DB tests
	}
}

func saveMessageRead(t *testing.T, db *sql.DB, mr *models.MessageRead) {
	if err := mr.Save(db); err != nil {
		t.Fatalf("Failed to save message_read: %v", err)
	}
}
