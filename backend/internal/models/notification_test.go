package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving notification
func TestNotificationSave(t *testing.T) {
	tables := "notification"
	db := setupTestDB(t, tables)
	defer db.Close()

	n := createNewNotification("Test Notification", "Hello world")
	saveNotification(t, db, n)

	if n.ID == 0 {
		t.Fatal("Expected notification ID to be set after save")
	}
}

// Test fetching all notifications
func TestNotificationGetAll(t *testing.T) {
	tables := "notification"
	db := setupTestDB(t, tables)
	defer db.Close()

	notifications := []*models.Notification{
		createNewNotification("CDTV", "content"),
		createNewNotification("Tokyo Dome", "content"),
		createNewNotification("5th Aniversary", "content"),
	}
	for _, n := range notifications {
		saveNotification(t, db, n)
	}

	returnedNotifications, err := models.GetAllNotifications(db)
	assertNoError(t, err, "Failed to fetch all notifications")
	assertCount(t, len(returnedNotifications), len(notifications), "notifications count mismatch")
}

// Test fetching empty notification table
func TestNotificationGetAllEmpty(t *testing.T) {
	tables := "notification"
	db := setupTestDB(t, tables)
	defer db.Close()

	notifications, err := models.GetAllNotifications(db)
	assertNoError(t, err, "Failed to fetch all notifications")
	assertCount(t, len(notifications), 0, "Expected notifications to be empty")
}

// Test updating notification
func TestNotificationUpdate(t *testing.T) {
	tables := "notification"
	db := setupTestDB(t, tables)
	defer db.Close()

	n := createNewNotification("Original Title", "Original Content")
	saveNotification(t, db, n)

	// Update fields
	n.Title = "Updated Title"
	n.Content = "Updated Content"

	err := n.Update(db)
	assertNoError(t, err, "Failed to update notification")

	// Verify update
	updated, _ := models.FindNotificationByID(db, n.ID)
	if updated.Title != "Updated Title" || updated.Content != "Updated Content" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding notification by ID
func TestFindNotificationByID(t *testing.T) {
	tables := "notification"
	db := setupTestDB(t, tables)
	defer db.Close()

	n := createNewNotification("Test Notification", "Hello world")
	saveNotification(t, db, n)

	notification, err := models.FindNotificationByID(db, n.ID)
	assertNoError(t, err, "Failed to find notification by ID")
	assertNotNil(t, notification, "Expected notification to be found")
	if notification.ID != n.ID {
		t.Errorf("Expected notification ID to be %d, got %d", n.ID, notification.ID)
	}
}

// Test finding non-existent notification
func TestFindNotificationByIDNotFound(t *testing.T) {
	tables := "notification"
	db := setupTestDB(t, tables)
	defer db.Close()

	notification, err := models.FindNotificationByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent notification ID")
	assertNil(t, notification, "Expected notification to be nil when not found")
}
