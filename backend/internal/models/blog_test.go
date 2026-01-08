package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving blog
func TestBlogSave(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	b := createNewBlog(m.ID, "Test blog", "Hello world", "draft")
	saveBlog(t, db, b)

	if b.ID == 0 {
		t.Fatal("Expected blog ID to be set after save")
	}
}

// Test saving blog with non existing foreign key
func TestBlogSaveFailsOnInvalidForeignKey(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	b := createNewBlog(1, "Test blog", "Hello world", "draft")

	err := b.Save(db)
	assertError(t, err, "Expected Blog.Save to fail due to invalid foreign key")
}

// Test fetching all blogs
func TestBlogGetAll(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	blogs := []*models.Blog{
		createNewBlog(m.ID, "CDTV", "content", "pending"),
		createNewBlog(m.ID, "Tokyo Dome", "content", "pending"),
		createNewBlog(m.ID, "5th Aniversary", "content", "pending"),
	}
	for _, blog := range blogs {
		saveBlog(t, db, blog)
	}

	returnedBlogs, err := models.GetAllBlogs(db, "pending")
	assertNoError(t, err, "Failed to fetch all blogs")
	assertCount(t, len(returnedBlogs), len(blogs), "blogs count mismatch")
}

// Test fetching empty blog table
func TestBlogGetAllEmpty(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	blogs, err := models.GetAllBlogs(db, "pending")
	assertNoError(t, err, "Failed to fetch all blogs")
	assertCount(t, len(blogs), 0, "Expected blogs to be empty")
}

// Test updating blog
func TestBlogUpdate(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Fujiyoshi")
	saveMember(t, db, m)

	b := createNewBlog(m.ID, "Original Title", "Original Content", "draft")
	saveBlog(t, db, b)

	// Update fields
	b.Title = "Updated Title"
	b.Status = "verified"

	err := b.Update(db)
	assertNoError(t, err, "Failed to update blog")

	// Verify update
	updated, err := models.FindBlogByID(db, b.ID, "verified")
	assertNoError(t, err, "Failed to find blog by ID")
	if updated.Title != "Updated Title" || updated.Status != "verified" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding blog by ID
func TestFindBlogByID(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	b := createNewBlog(m.ID, "Test blog", "Hello world", "pending")
	saveBlog(t, db, b)

	blog, err := models.FindBlogByID(db, b.ID, "pending")
	assertNoError(t, err, "Failed to find blog by ID")
	assertNotNil(t, blog, "Expected blog to be found")
	if blog.ID != b.ID {
		t.Errorf("Expected blog ID to be 1, got %d", blog.ID)
	}
}

// Test finding non-existent blog
func TestFindBlogByIDNotFound(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	blog, err := models.FindBlogByID(db, 9999, "pending")
	assertError(t, err, "Expected error when searching for non-existent blog ID")
	assertNil(t, blog, "Expected blog to be nil when not found")
}

// Test large content body
func TestBlogLargeContent(t *testing.T) {
	tables := "blog, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Tamura")
	saveMember(t, db, m)

	largeContent := make([]byte, 10000) // 10KB of data
	for i := range largeContent {
		largeContent[i] = 'a'
	}

	b := createNewBlog(m.ID, "Long Post", string(largeContent), "pending")
	saveBlog(t, db, b)
}
