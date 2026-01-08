package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving official_news
func TestOfficialNewsSave(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	tag := "test"
	n := createNewOfficialNews("Test OfficialNews", &tag, "test content")
	saveOfficialNews(t, db, n)

	if n.ID == 0 {
		t.Fatal("Expected official_news ID to be set after save")
	}
}

// Test saving official_news with complex content
func TestOfficialNewsContent(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	content := `Line 1
Line 2
Special chars: !@#$%^&*()`
	n := createNewOfficialNews("Content Test", nil, content)
	saveOfficialNews(t, db, n)

	saved, err := models.FindOfficialNewsByID(db, n.ID)
	assertNoError(t, err, "Failed to find official_news")
	if saved.Content != content {
		t.Errorf("Content mismatch. Expected:\n%s\nGot:\n%s", content, saved.Content)
	}
}

// Test fetching all official_news
func TestOfficialNewsGetAll(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	tag1 := "tag1"
	tag2 := "tag2"
	news := []*models.OfficialNews{
		createNewOfficialNews("CDTV", &tag1, "content"),
		createNewOfficialNews("Tokyo Dome", &tag2, "content"),
		createNewOfficialNews("5th Aniversary", nil, "content"),
	}
	for _, n := range news {
		saveOfficialNews(t, db, n)
	}

	returnedNews, err := models.GetAllOfficialNews(db)
	assertNoError(t, err, "Failed to fetch all official_news")
	assertCount(t, len(returnedNews), len(news), "official_news count mismatch")
}

// Test fetching empty official_news table
func TestOfficialNewsGetAllEmpty(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	news, err := models.GetAllOfficialNews(db)
	assertNoError(t, err, "Failed to fetch all official_news")
	assertCount(t, len(news), 0, "Expected official_news to be empty")
}

// Test updating official_news
func TestOfficialNewsUpdate(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	n := createNewOfficialNews("Original Title", nil, "content")
	saveOfficialNews(t, db, n)

	// Update fields
	n.Title = "Updated Title"
	tag := "updated"
	n.Tag = &tag
	n.Content = "Updated Content"

	err := n.Update(db)
	assertNoError(t, err, "Failed to update official_news")

	// Verify update
	updated, _ := models.FindOfficialNewsByID(db, n.ID)
	if updated.Title != "Updated Title" || *updated.Tag != "updated" || updated.Content != "Updated Content" {
		t.Errorf("Update did not persist changes correctly: got Title=%s, Tag=%s, Content=%s", updated.Title, *updated.Tag, updated.Content)
	}
}

// Test finding official_news by ID
func TestFindOfficialNewsByID(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	n := createNewOfficialNews("Test OfficialNews", nil, "content")
	saveOfficialNews(t, db, n)

	news, err := models.FindOfficialNewsByID(db, n.ID)
	assertNoError(t, err, "Failed to find official_news by ID")
	assertNotNil(t, news, "Expected official_news to be found")
	if news.ID != n.ID {
		t.Errorf("Expected official_news ID to be %d, got %d", n.ID, news.ID)
	}
}

// Test finding non-existent official_news
func TestFindOfficialNewsByIDNotFound(t *testing.T) {
	tables := "official_news"
	db := setupTestDB(t, tables)
	defer db.Close()

	news, err := models.FindOfficialNewsByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent official_news ID")
	assertNil(t, news, "Expected official_news to be nil when not found")
}
