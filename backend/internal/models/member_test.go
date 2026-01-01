package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving member
func TestMemberSave(t *testing.T) {
	tables := "member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	if m.ID == 0 {
		t.Fatal("Expected member ID to be set after save")
	}
}

// Test fetching all members
func TestMemberGetAll(t *testing.T) {
	tables := "member"
	db := setupTestDB(t, tables)
	defer db.Close()

	members := []*models.Member{
		createNewMember("Sugai"),
		createNewMember("Fujiyoshi"),
		createNewMember("Tamura"),
	}
	for _, m := range members {
		saveMember(t, db, m)
	}

	returnedMembers, err := models.GetAllMembers(db)
	assertNoError(t, err, "Failed to fetch all members")
	assertCount(t, len(returnedMembers), len(members), "member count mismatch")
}

// Test fetching empty member table
func TestMemberGetAllEmpty(t *testing.T) {
	tables := "member"
	db := setupTestDB(t, tables)
	defer db.Close()

	members, err := models.GetAllMembers(db)
	assertNoError(t, err, "Failed to fetch all members")
	assertCount(t, len(members), 0, "Expected members to be empty")
}

// Test updating member
func TestMemberUpdate(t *testing.T) {
	tables := "member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Fujiyoshi")
	saveMember(t, db, m)

	// Update fields
	m.Name = "Karin"
	gen := 2
	m.Generation = &gen

	err := m.Update(db)
	assertNoError(t, err, "Failed to update member")

	// Verify update
	updated, _ := models.FindMemberByID(db, m.ID)
	if updated.Name != "Karin" || *updated.Generation != 2 {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding member by ID
func TestFindMemberByID(t *testing.T) {
	tables := "member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	found, err := models.FindMemberByID(db, m.ID)
	assertNoError(t, err, "Failed to find member by ID")
	assertNotNil(t, found, "Expected member to be found")
	if found.ID != m.ID {
		t.Errorf("Expected member ID to be %d, got %d", m.ID, found.ID)
	}
}

// Test finding non-existent member
func TestFindMemberByIDNotFound(t *testing.T) {
	tables := "member"
	db := setupTestDB(t, tables)
	defer db.Close()

	found, err := models.FindMemberByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent member ID")
	assertNil(t, found, "Expected member to be nil when not found")
}
