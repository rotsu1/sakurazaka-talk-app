package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving staff
func TestStaffSave(t *testing.T) {
	tables := "staff, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	s := createNewStaff(&m.ID, "admin", "testuser", "password")
	saveStaff(t, db, s)

	if s.ID == 0 {
		t.Fatal("Expected staff ID to be set after save")
	}
}

// Test saving staff with non existing foreign key
func TestStaffSaveFailsOnInvalidForeignKey(t *testing.T) {
	tables := "staff"
	db := setupTestDB(t, tables)
	defer db.Close()

	invalidMemberID := 999
	s := createNewStaff(&invalidMemberID, "admin", "testuser", "password")

	err := s.Save(db)
	assertError(t, err, "Expected Staff.Save to fail due to invalid foreign key")
}

// Test fetching all staff
func TestStaffGetAll(t *testing.T) {
	tables := "staff, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	staffs := []*models.Staff{
		createNewStaff(&m.ID, "admin", "user1", "pass1"),
		createNewStaff(&m.ID, "editor", "user2", "pass2"),
		createNewStaff(nil, "viewer", "user3", "pass3"),
	}
	for _, s := range staffs {
		saveStaff(t, db, s)
	}

	returnedStaffs, err := models.GetAllStaff(db)
	assertNoError(t, err, "Failed to fetch all staff")
	assertCount(t, len(returnedStaffs), len(staffs), "staffs count mismatch")
}

// Test fetching empty staff table
func TestStaffGetAllEmpty(t *testing.T) {
	tables := "staff"
	db := setupTestDB(t, tables)
	defer db.Close()

	staffs, err := models.GetAllStaff(db)
	assertNoError(t, err, "Failed to fetch all staff")
	assertCount(t, len(staffs), 0, "Expected staffs to be empty")
}

// Test updating staff
func TestStaffUpdate(t *testing.T) {
	tables := "staff, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Fujiyoshi")
	saveMember(t, db, m)

	s := createNewStaff(&m.ID, "admin", "user1", "pass1")
	saveStaff(t, db, s)

	// Update fields
	s.Username = "updatedUser"
	s.Role = "super_admin"

	err := s.Update(db)
	assertNoError(t, err, "Failed to update staff")

	// Verify update
	updated, _ := models.FindStaffByID(db, s.ID)
	if updated.Username != "updatedUser" || updated.Role != "super_admin" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding staff by ID
func TestFindStaffByID(t *testing.T) {
	tables := "staff, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	s := createNewStaff(&m.ID, "admin", "user1", "pass1")
	saveStaff(t, db, s)

	staff, err := models.FindStaffByID(db, s.ID)
	assertNoError(t, err, "Failed to find staff by ID")
	assertNotNil(t, staff, "Expected staff to be found")
	if staff.ID != s.ID {
		t.Errorf("Expected staff ID to be %d, got %d", s.ID, staff.ID)
	}
}

// Test finding non-existent staff
func TestFindStaffByIDNotFound(t *testing.T) {
	tables := "staff"
	db := setupTestDB(t, tables)
	defer db.Close()

	staff, err := models.FindStaffByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent staff ID")
	assertNil(t, staff, "Expected staff to be nil when not found")
}
