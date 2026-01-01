package models_test

import (
	"backend/internal/models"
	"testing"
	"time"
)

// Test saving talk_user
func TestTalkUserSave(t *testing.T) {
	tables := "talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	if tu.ID == 0 {
		t.Fatal("Expected talk_user ID to be set after save")
	}
}

// Test fetching all talk_users
func TestTalkUserGetAll(t *testing.T) {
	tables := "talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	users := []*models.TalkUser{
		createNewTalkUser(),
		createNewTalkUser(),
		createNewTalkUser(),
	}
	for _, u := range users {
		saveTalkUser(t, db, u)
	}

	returnedUsers, err := models.GetAllTalkUsers(db)
	assertNoError(t, err, "Failed to fetch all talk_users")
	assertCount(t, len(returnedUsers), len(users), "talk_users count mismatch")
}

// Test fetching empty talk_user table
func TestTalkUserGetAllEmpty(t *testing.T) {
	tables := "talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	users, err := models.GetAllTalkUsers(db)
	assertNoError(t, err, "Failed to fetch all talk_users")
	assertCount(t, len(users), 0, "Expected talk_users to be empty")
}

// Test updating talk_user
func TestTalkUserUpdate(t *testing.T) {
	tables := "talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	originalUpdatedAt := tu.UpdatedAt

	// Sleep to ensure the updated_at timestamp changes
	time.Sleep(1 * time.Second)

	err := tu.Update(db)
	assertNoError(t, err, "Failed to update talk_user")

	// Verify update
	updated, _ := models.FindTalkUserByID(db, tu.ID)
	if updated.UpdatedAt.Equal(originalUpdatedAt) {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding talk_user by ID
func TestFindTalkUserByID(t *testing.T) {
	tables := "talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	found, err := models.FindTalkUserByID(db, tu.ID)
	assertNoError(t, err, "Failed to find talk_user by ID")
	assertNotNil(t, found, "Expected talk_user to be found")
	if found.ID != tu.ID {
		t.Errorf("Expected talk_user ID to be %d, got %d", tu.ID, found.ID)
	}
}

// Test finding non-existent talk_user
func TestFindTalkUserByIDNotFound(t *testing.T) {
	tables := "talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	found, err := models.FindTalkUserByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent talk_user ID")
	assertNil(t, found, "Expected talk_user to be nil when not found")
}
