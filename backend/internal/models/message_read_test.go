package models_test

import (
	"backend/internal/models"
	"testing"
)

// Test saving message_read
func TestMessageReadSave(t *testing.T) {
	tables := "message_read, message, talk_user, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// 1. Create dependencies
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Hello world", "text", "draft")
	saveMessage(t, db, msg)

	user := createNewTalkUser()
	saveTalkUser(t, db, user)

	// 2. Create message_read
	mr := createNewMessageRead(msg.ID, user.ID)

	// 3. Save
	saveMessageRead(t, db, mr)
}

// Test saving message_read with invalid FK
func TestMessageReadSaveFailsOnInvalidForeignKey(t *testing.T) {
	tables := "message_read"
	db := setupTestDB(t, tables)
	defer db.Close()

	mr := createNewMessageRead(999, 999)

	err := mr.Save(db)
	assertError(t, err, "Expected Save to fail due to invalid foreign key")
}

// Test fetching all message_reads
func TestMessageReadGetAll(t *testing.T) {
	tables := "message_read, message, talk_user, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	// Setup dependencies
	m := createNewMember("Sugai")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Hello world", "text", "draft")
	saveMessage(t, db, msg)

	user1 := createNewTalkUser()
	saveTalkUser(t, db, user1)
	user2 := createNewTalkUser()
	saveTalkUser(t, db, user2)

	// Create reads
	reads := []*models.MessageRead{
		createNewMessageRead(msg.ID, user1.ID),
		createNewMessageRead(msg.ID, user2.ID),
	}

	for _, mr := range reads {
		saveMessageRead(t, db, mr)
	}

	// Fetch all
	returnedReads, err := models.GetAllMessageReads(db)
	assertNoError(t, err, "Failed to fetch all message_reads")
	assertCount(t, len(returnedReads), len(reads), "message_reads count mismatch")
}

// Test finding message_read by IDs
func TestFindMessageReadByIDs(t *testing.T) {
	tables := "message_read, message, talk_user, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Target Message", "text", "draft")
	saveMessage(t, db, msg)

	user := createNewTalkUser()
	saveTalkUser(t, db, user)

	mr := createNewMessageRead(msg.ID, user.ID)
	saveMessageRead(t, db, mr)

	found, err := models.FindMessageReadByIDs(db, msg.ID, user.ID)
	assertNoError(t, err, "Failed to find message_read")
	assertNotNil(t, found, "Expected message_read to be found")
	if found.MessageID != msg.ID || found.UserID != user.ID {
		t.Errorf("Expected IDs %d, %d, got %d, %d", msg.ID, user.ID, found.MessageID, found.UserID)
	}
}

// Test deleting message_read
func TestMessageReadDelete(t *testing.T) {
	tables := "message_read, message, talk_user, member"
	db := setupTestDB(t, tables)
	defer db.Close()

	m := createNewMember("Sugai")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID, "Delete Me", "text", "draft")
	saveMessage(t, db, msg)

	user := createNewTalkUser()
	saveTalkUser(t, db, user)

	mr := createNewMessageRead(msg.ID, user.ID)
	saveMessageRead(t, db, mr)

	// Delete
	err := mr.Delete(db)
	assertNoError(t, err, "Failed to delete message_read")

	// Verify deletion
	_, err = models.FindMessageReadByIDs(db, msg.ID, user.ID)
	assertError(t, err, "Expected error when searching for deleted message_read")
}
