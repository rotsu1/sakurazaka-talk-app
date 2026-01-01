package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving message
func TestMessageSave(t *testing.T) {
	tables := "message, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Hello world", "text", "draft")
	saveMessage(t, db, msg)

	if msg.ID == 0 {
		t.Fatal("Expected message ID to be set after save")
	}
}

// Test saving message with non existing foreign key
func TestMessageSaveFailsOnInvalidForeignKey(t *testing.T) {
	tables := "message"
	db := setupTestDB(t, tables)
	defer db.Close()

	msg := createNewMessage(999, "Hello world", "text", "draft")

	err := msg.Save(db)
	assertError(t, err, "Expected Message.Save to fail due to invalid foreign key")
}

// Test fetching all messages
func TestMessageGetAll(t *testing.T) {
	tables := "message, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	messages := []*models.Message{
		createNewMessage(m.ID, "CDTV", "text", "pending"),
		createNewMessage(m.ID, "Tokyo Dome", "text", "pending"),
		createNewMessage(m.ID, "5th Aniversary", "text", "pending"),
	}
	for _, msg := range messages {
		saveMessage(t, db, msg)
	}

	returnedMessages, err := models.GetAllMessages(db)
	assertNoError(t, err, "Failed to fetch all messages")
	assertCount(t, len(returnedMessages), len(messages), "messages count mismatch")
}

// Test fetching empty message table
func TestMessageGetAllEmpty(t *testing.T) {
	tables := "message"
	db := setupTestDB(t, tables)
	defer db.Close()

	messages, err := models.GetAllMessages(db)
	assertNoError(t, err, "Failed to fetch all messages")
	assertCount(t, len(messages), 0, "Expected messages to be empty")
}

// Test updating message
func TestMessageUpdate(t *testing.T) {
	tables := "message, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Fujiyoshi")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Original Content", "text", "draft")
	saveMessage(t, db, msg)

	// Update fields
	msg.Content = "Updated Content"
	msg.Status = "published"

	err := msg.Update(db)
	assertNoError(t, err, "Failed to update message")

	// Verify update
	updated, _ := models.FindMessageByID(db, msg.ID)
	if updated.Content != "Updated Content" || updated.Status != "published" {
		t.Errorf("Update did not persist changes correctly")
	}
}

// Test finding message by ID
func TestFindMessageByID(t *testing.T) {
	tables := "message, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Hello world", "text", "draft")
	saveMessage(t, db, msg)

	message, err := models.FindMessageByID(db, msg.ID)
	assertNoError(t, err, "Failed to find message by ID")
	assertNotNil(t, message, "Expected message to be found")
	if message.ID != msg.ID {
		t.Errorf("Expected message ID to be %d, got %d", msg.ID, message.ID)
	}
}

// Test finding non-existent message
func TestFindMessageByIDNotFound(t *testing.T) {
	tables := "message"
	db := setupTestDB(t, tables)
	defer db.Close()

	message, err := models.FindMessageByID(db, 9999)
	assertError(t, err, "Expected error when searching for non-existent message ID")
	assertNil(t, message, "Expected message to be nil when not found")
}
