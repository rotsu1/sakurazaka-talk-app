package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetMessages(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Message Sender")
	saveMember(t, db, m)

	msg1 := createNewMessage(m.ID)
	msg2 := createNewMessage(m.ID)
	saveMessage(t, db, msg1)
	saveMessage(t, db, msg2)

	rr := helperServeHTTP(
		"GET", "/message/", nil,
		withDB(db, routers.GetMessages),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var messages []models.Message
	decodeResponse(t, rr.Body, &messages)

	checkResponseNumber(t, "messages", len(messages), 2)
}

func TestGetMessageByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Message Sender")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)

	rr := helperServeHTTP(
		"GET", "/message/"+strconvInt(msg.ID), nil,
		withDB(db, routers.GetMessageByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateMessage(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Message Sender")
	saveMember(t, db, m)

	newMessage := createNewMessage(m.ID)
	body, _ := json.Marshal(newMessage)

	rr := helperServeHTTP(
		"POST", "/message/", bytes.NewBuffer(body),
		withDB(db, routers.CreateMessage),
	)
	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateMessage(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Message Sender")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)

	updatedMessage := createNewMessage(m.ID)
	updatedMessage.Type = "image"
	updatedMessage.Content = "http://example.com/image.jpg"
	updatedMessage.Status = "approved"
	body, _ := json.Marshal(updatedMessage)

	rr := helperServeHTTP(
		"PUT", "/message/"+strconvInt(msg.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateMessage),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindMessageByID(db, msg.ID)
	checkUpdated(t, "type", "image", updated.Type)
}

func TestDeleteMessage(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Message Sender")
	saveMember(t, db, m)

	msg := createNewMessage(m.ID)
	saveMessage(t, db, msg)

	rr := helperServeHTTP(
		"DELETE", "/message/"+strconvInt(msg.ID), nil,
		withDB(db, routers.DeleteMessage),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindMessageByID(db, msg.ID)
	verifyDeleted(t, err, "message")
}
