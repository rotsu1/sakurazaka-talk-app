package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving fanletter
func TestFanletterSave(t *testing.T) {
	tables := "fanletter, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	f := createNewFanletter(m.ID, tu.ID, "Hello world")
	saveFanletter(t, db, f)

	if f.ID == 0 {
		t.Fatal("Expected fanletter ID to be set after save")
	}
}

// Test saving fanletter with non existing foreign key
func TestFanletterSaveFailsOnInvalidForeignKey(t *testing.T) {
	tables := "fanletter, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	f := createNewFanletter(999, 999, "Hello world")

	err := f.Save(db)
	assertError(t, err, "Expected Fanletter Save to fail due to invalid foreign key")
}

// Test fetching all fanletters
func TestFanletterGetAll(t *testing.T) {
	tables := "fanletter, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	fanletters := []*models.Fanletter{
		createNewFanletter(m.ID, tu.ID, "CDTV"),
		createNewFanletter(m.ID, tu.ID, "Tokyo Dome"),
		createNewFanletter(m.ID, tu.ID, "5th Aniversary"),
	}
	for _, f := range fanletters {
		saveFanletter(t, db, f)
	}

	returnedFanletters, err := models.GetAllFanletters(db)
	assertNoError(t, err, "Failed to fetch all fanletters")
	assertCount(t, len(returnedFanletters), len(fanletters), "fanletters count mismatch")
}

// Test fetching empty fanletter table
func TestFanletterGetAllEmpty(t *testing.T) {
	tables := "fanletter"
	db := setupTestDB(t, tables)
	defer db.Close()

	fanletters, err := models.GetAllFanletters(db)
	assertNoError(t, err, "Failed to fetch all fanletters")
	assertCount(t, len(fanletters), 0, "Expected fanletters to be empty")
}

// Test updating fanletter
func TestFanletterUpdate(t *testing.T) {
	tables := "fanletter, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Fujiyoshi")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	f := createNewFanletter(m.ID, tu.ID, "Original Content")
	saveFanletter(t, db, f)

	// Update fields
	f.Content = "Updated Content"

	err := f.Update(db)
	assertNoError(t, err, "Failed to update fanletter")

	// Verify update
	updated, _ := models.FindFanletterByID(db, f.ID)
	if updated.Content != "Updated Content" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding fanletter by ID
func TestFindFanletterByID(t *testing.T) {
	tables := "fanletter, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Insert member first to prevent foreign key violation
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	f := createNewFanletter(m.ID, tu.ID, "Hello world")
	saveFanletter(t, db, f)

	fanletter, err := models.FindFanletterByID(db, f.ID)
	assertNoError(t, err, "Failed to find fanletter by ID")
	assertNotNil(t, fanletter, "Expected fanletter to be found")
	if fanletter.ID != f.ID {
		t.Errorf("Expected fanletter ID to be %d, got %d", f.ID, fanletter.ID)
	}
}

// Test finding non-existent fanletter
func TestFindFanletterByIDNotFound(t *testing.T) {
	tables := "fanletter"
	db := setupTestDB(t, tables)
	defer db.Close()

	fanletter, err := models.FindFanletterByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent fanletter ID")
	assertNil(t, fanletter, "Expected fanletter to be nil when not found")
}
