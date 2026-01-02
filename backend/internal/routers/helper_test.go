package routers_test

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// Helper to initialize database
func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := db.InitDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Truncate all tables to ensure clean state
	tables := []string{
		"member", "blog", "fanletter", "message", "notification",
		"official_news", "staff", "talk_user", "talk_user_member", "template",
	}

	for _, table := range tables {
		_, err = db.Exec("TRUNCATE " + table + " RESTART IDENTITY CASCADE")
		if err != nil {
			log.Fatalf("Failed to truncate table: %v", err)
		}
	}

	return db
}

func strconvInt(i int) string {
	return strconv.Itoa(i)
}

func createNewMember(name string) *models.Member {
	gen := 1
	return &models.Member{
		Name:       name,
		Generation: &gen,
	}
}

func saveMember(t *testing.T, db *sql.DB, m *models.Member) {
	t.Helper()
	if err := m.Save(db); err != nil {
		t.Fatalf("Failed to save member: %v", err)
	}
}

func createNewStaff(mID *int, role string, username string) *models.Staff {
	return &models.Staff{
		MemberID: mID,
		Role:     role,
		Username: username,
	}
}

func saveStaff(t *testing.T, db *sql.DB, s *models.Staff) {
	t.Helper()
	if err := s.Save(db); err != nil {
		t.Fatalf("Failed to save staff: %v", err)
	}
}

func createNewBlog(memberID int, title string) *models.Blog {
	return &models.Blog{
		MemberID: memberID,
		Title:    title,
		Content:  "Test Content",
		Status:   "pending",
	}
}

func saveBlog(t *testing.T, db *sql.DB, b *models.Blog) {
	t.Helper()
	if err := b.Save(db); err != nil {
		t.Fatalf("Failed to save blog: %v", err)
	}
}

func createNewFanletter(memberID, talkUserID int) *models.Fanletter {
	return &models.Fanletter{
		MemberID:   memberID,
		TalkUserID: talkUserID,
		Content:    "Test Fanletter",
	}
}

func saveFanletter(t *testing.T, db *sql.DB, f *models.Fanletter) {
	t.Helper()
	if err := f.Save(db); err != nil {
		t.Fatalf("Failed to save fanletter: %v", err)
	}
}

func createNewMessage(memberID int) *models.Message {
	return &models.Message{
		MemberID: memberID,
		Type:     "text",
		Content:  "Test Message",
		Status:   "approved",
	}
}

func saveMessage(t *testing.T, db *sql.DB, m *models.Message) {
	t.Helper()
	if err := m.Save(db); err != nil {
		t.Fatalf("Failed to save message: %v", err)
	}
}

func createNewNotification() *models.Notification {
	return &models.Notification{
		Title:   "Test Notification",
		Content: "Test Content",
	}
}

func saveNotification(t *testing.T, db *sql.DB, n *models.Notification) {
	t.Helper()
	if err := n.Save(db); err != nil {
		t.Fatalf("Failed to save notification: %v", err)
	}
}

func createNewOfficialNews() *models.OfficialNews {
	tag := "Info"
	return &models.OfficialNews{
		Title: "Test News",
		Tag:   &tag,
	}
}

func saveOfficialNews(t *testing.T, db *sql.DB, o *models.OfficialNews) {
	t.Helper()
	if err := o.Save(db); err != nil {
		t.Fatalf("Failed to save official news: %v", err)
	}
}

func createNewTalkUser() *models.TalkUser {
	return &models.TalkUser{}
}

func saveTalkUser(t *testing.T, db *sql.DB, tu *models.TalkUser) {
	t.Helper()
	if err := tu.Save(db); err != nil {
		t.Fatalf("Failed to save talk user: %v", err)
	}
}

func createNewTalkUserMember(userID, memberID int) *models.TalkUserMember {
	return &models.TalkUserMember{
		UserID:   userID,
		MemberID: memberID,
		Status:   "active",
	}
}

func saveTalkUserMember(t *testing.T, db *sql.DB, tum *models.TalkUserMember) {
	t.Helper()
	if err := tum.Save(db); err != nil {
		t.Fatalf("Failed to save talk user member: %v", err)
	}
}

func createNewTemplate() *models.Template {
	return &models.Template{
		TemplateURL: "http://example.com/template.jpg",
	}
}

func saveTemplate(t *testing.T, db *sql.DB, tmpl *models.Template) {
	t.Helper()
	if err := tmpl.Save(db); err != nil {
		t.Fatalf("Failed to save template: %v", err)
	}
}

// Helper to serve HTTP request and returns response recorder
func helperServeHTTP(
	method string,
	path string,
	body io.Reader,
	handlerFunc func(w http.ResponseWriter, r *http.Request),
) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	})

	handler.ServeHTTP(rr, req)
	return rr
}

// Helper to wrap handler function with database connection
func withDB(
	db *sql.DB,
	handler func(db *sql.DB, w http.ResponseWriter, r *http.Request),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(db, w, r)
	}
}

func checkStatusCode(t *testing.T, status int, expected int) {
	t.Helper()
	if status != expected {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expected)
	}
}

func decodeResponse[T any](t *testing.T, body io.Reader, v *T) {
	t.Helper()
	if err := json.NewDecoder(body).Decode(v); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
}

func checkUpdated(t *testing.T, field string, expected string, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected updated %s '%s', got '%s'", field, expected, actual)
	}
}

func verifyDeleted(t *testing.T, err error, model string) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error finding deleted %s, got nil", model)
	}
}

func checkResponseNumber(t *testing.T, field string, length int, expected int) {
	t.Helper()
	if length != expected {
		t.Errorf("Expected %d %s, got %d", expected, field, length)
	}
}
