package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetMessageRead(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("member")
	saveMember(t, db, m)
	msg1 := createNewMessage(m.ID)
	saveMessage(t, db, msg1)
	msg2 := createNewMessage(m.ID)
	saveMessage(t, db, msg2)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)
	mr1 := createNewMessageRead(msg1.ID, tu.ID)
	saveMessageRead(t, db, mr1)
	mr2 := createNewMessageRead(msg2.ID, tu.ID)
	saveMessageRead(t, db, mr2)

	rr := helperServeHTTP(
		"GET", "/message_read/", nil,
		withDB(db, routers.GetMessageRead),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var MessageReads []models.MessageRead
	decodeResponse(t, rr.Body, &MessageReads)
	checkResponseNumber(t, "Messages read", len(MessageReads), 2)
}

func TestGetMessageReadByIDs(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("member")
	saveMember(t, db, m)
	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	mr := createNewMessageRead(msg.ID, tu.ID)
	saveMessageRead(t, db, mr)

	url := "/message_read/" + strconvInt(msg.ID) + "/" + strconvInt(tu.ID)
	rr := helperServeHTTP(
		"GET", url, nil,
		withDB(db, routers.GetMessageReadByIDs),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var responseMR models.MessageRead
	decodeResponse(t, rr.Body, &responseMR)

	if responseMR.MessageID != msg.ID || responseMR.UserID != tu.ID {
		t.Errorf("Expected MessageID %d, UserID %d; got MessageID %d, UserID %d",
			msg.ID, tu.ID, responseMR.MessageID, responseMR.UserID)
	}
}

func TestCreateMessageRead(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("member")
	saveMember(t, db, m)
	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	newMR := createNewMessageRead(msg.ID, tu.ID)
	body, _ := json.Marshal(newMR)

	rr := helperServeHTTP(
		"POST", "/message_read/", bytes.NewReader(body),
		withDB(db, routers.CreateMessageReadByID),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)

	savedMR, err := models.FindMessageReadByIDs(db, msg.ID, tu.ID)
	if err != nil {
		t.Fatalf("Failed to find created message read: %v", err)
	}
	if savedMR.MessageID != msg.ID || savedMR.UserID != tu.ID {
		t.Errorf("Saved message read does not match expected IDs")
	}
}

func TestUpdateMessageRead(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("member")
	saveMember(t, db, m)
	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	mr := createNewMessageRead(msg.ID, tu.ID)
	saveMessageRead(t, db, mr)

	// Update the message read (sets ReadAt to NOW)
	body, _ := json.Marshal(mr)
	rr := helperServeHTTP(
		"PUT", "/message_read/", bytes.NewReader(body),
		withDB(db, routers.UpdateMessageRead),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindMessageReadByIDs(db, msg.ID, tu.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve updated message read: %v", err)
	}
}

func TestDeleteMessageRead(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("member")
	saveMember(t, db, m)
	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	mr := createNewMessageRead(msg.ID, tu.ID)
	saveMessageRead(t, db, mr)

	url := "/message_read/" + strconvInt(msg.ID) + "/" + strconvInt(tu.ID)
	rr := helperServeHTTP(
		"DELETE", url, nil,
		withDB(db, routers.DeleteMessageRead),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindMessageReadByIDs(db, msg.ID, tu.ID)
	if err == nil {
		t.Error("Expected error finding deleted message read, got nil")
	}
}
