package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving template
func TestTemplateSave(t *testing.T) {
	tables := "template"
	db := setupTestDB(t, tables)
	defer db.Close()

	tpl := createNewTemplate("http://example.com/template.jpg")
	saveTemplate(t, db, tpl)

	if tpl.ID == 0 {
		t.Fatal("Expected template ID to be set after save")
	}
}

// Test fetching all templates
func TestTemplateGetAll(t *testing.T) {
	tables := "template"
	db := setupTestDB(t, tables)
	defer db.Close()

	templates := []*models.Template{
		createNewTemplate("http://example.com/1.jpg"),
		createNewTemplate("http://example.com/2.jpg"),
		createNewTemplate("http://example.com/3.jpg"),
	}
	for _, tpl := range templates {
		saveTemplate(t, db, tpl)
	}

	returnedTemplates, err := models.GetAllTemplates(db)
	assertNoError(t, err, "Failed to fetch all templates")
	assertCount(t, len(returnedTemplates), len(templates), "templates count mismatch")
}

// Test fetching empty template table
func TestTemplateGetAllEmpty(t *testing.T) {
	tables := "template"
	db := setupTestDB(t, tables)
	defer db.Close()

	templates, err := models.GetAllTemplates(db)
	assertNoError(t, err, "Failed to fetch all templates")
	assertCount(t, len(templates), 0, "Expected templates to be empty")
}

// Test updating template
func TestTemplateUpdate(t *testing.T) {
	tables := "template"
	db := setupTestDB(t, tables)
	defer db.Close()

	tpl := createNewTemplate("http://example.com/original.jpg")
	saveTemplate(t, db, tpl)

	// Update fields
	tpl.TemplateURL = "http://example.com/updated.jpg"

	err := tpl.Update(db)
	assertNoError(t, err, "Failed to update template")

	// Verify update
	updated, _ := models.FindTemplateByID(db, tpl.ID)
	if updated.TemplateURL != "http://example.com/updated.jpg" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding template by ID
func TestFindTemplateByID(t *testing.T) {
	tables := "template"
	db := setupTestDB(t, tables)
	defer db.Close()

	tpl := createNewTemplate("http://example.com/template.jpg")
	saveTemplate(t, db, tpl)

	found, err := models.FindTemplateByID(db, tpl.ID)
	assertNoError(t, err, "Failed to find template by ID")
	assertNotNil(t, found, "Expected template to be found")
	if found.ID != tpl.ID {
		t.Errorf("Expected template ID to be %d, got %d", tpl.ID, found.ID)
	}
}

// Test finding non-existent template
func TestFindTemplateByIDNotFound(t *testing.T) {
	tables := "template"
	db := setupTestDB(t, tables)
	defer db.Close()

	found, err := models.FindTemplateByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent template ID")
	assertNil(t, found, "Expected template to be nil when not found")
}
