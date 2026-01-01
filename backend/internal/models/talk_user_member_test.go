package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving talk_user_member
func TestTalkUserMemberSave(t *testing.T) {
	tables := "talk_user_member, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum := createNewTalkUserMember(tu.ID, m.ID, "active")
	saveTalkUserMember(t, db, tum)

	if tum.ID == 0 {
		t.Fatal("Expected talk_user_member ID to be set after save")
	}
}

// Test saving talk_user_member with non existing foreign key
func TestTalkUserMemberSaveFailsOnInvalidForeignKey(t *testing.T) {
	tables := "talk_user_member, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	tum := createNewTalkUserMember(999, 999, "active")

	err := tum.Save(db)
	assertError(t, err, "Expected TalkUserMember.Save to fail due to invalid foreign key")
}

// Test fetching all talk_user_members
func TestTalkUserMemberGetAll(t *testing.T) {
	tables := "talk_user_member, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tums := []*models.TalkUserMember{
		createNewTalkUserMember(tu.ID, m.ID, "active"),
		createNewTalkUserMember(tu.ID, m.ID, "inactive"),
		createNewTalkUserMember(tu.ID, m.ID, "pending"),
	}
	for _, tum := range tums {
		// Create new members for each subscription
		member := createNewMember("member")
		saveMember(t, db, member)
		tum.MemberID = member.ID
		saveTalkUserMember(t, db, tum)
	}

	returnedTums, err := models.GetAllTalkUserMembers(db)
	assertNoError(t, err, "Failed to fetch all talk_user_members")
	assertCount(t, len(returnedTums), len(tums), "talk_user_members count mismatch")
}

// Test fetching empty talk_user_member table
func TestTalkUserMemberGetAllEmpty(t *testing.T) {
	tables := "talk_user_member"
	db := setupTestDB(t, tables)
	defer db.Close()

	tums, err := models.GetAllTalkUserMembers(db)
	assertNoError(t, err, "Failed to fetch all talk_user_members")
	assertCount(t, len(tums), 0, "Expected talk_user_members to be empty")
}

// Test updating talk_user_member
func TestTalkUserMemberUpdate(t *testing.T) {
	tables := "talk_user_member, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Fujiyoshi")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum := createNewTalkUserMember(tu.ID, m.ID, "active")
	saveTalkUserMember(t, db, tum)

	// Update fields
	tum.Status = "inactive"

	err := tum.Update(db)
	assertNoError(t, err, "Failed to update talk_user_member")

	// Verify update
	updated, _ := models.FindTalkUserMemberByID(db, tum.ID)
	if updated.Status != "inactive" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding talk_user_member by ID
func TestFindTalkUserMemberByID(t *testing.T) {
	tables := "talk_user_member, member, talk_user"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum := createNewTalkUserMember(tu.ID, m.ID, "active")
	saveTalkUserMember(t, db, tum)

	found, err := models.FindTalkUserMemberByID(db, tum.ID)
	assertNoError(t, err, "Failed to find talk_user_member by ID")
	assertNotNil(t, found, "Expected talk_user_member to be found")
	if found.ID != tum.ID {
		t.Errorf("Expected talk_user_member ID to be %d, got %d", tum.ID, found.ID)
	}
}

// Test finding non-existent talk_user_member
func TestFindTalkUserMemberByIDNotFound(t *testing.T) {
	tables := "talk_user_member"
	db := setupTestDB(t, tables)
	defer db.Close()

	found, err := models.FindTalkUserMemberByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent talk_user_member ID")
	assertNil(t, found, "Expected talk_user_member to be nil when not found")
}
