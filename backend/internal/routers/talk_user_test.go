package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"net/http"
	"testing"
)

func TestGetTalkUsers(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tu1 := createNewTalkUser()
	tu2 := createNewTalkUser()
	saveTalkUser(t, db, tu1)
	saveTalkUser(t, db, tu2)

	rr := helperServeHTTP(
		"GET", "/talk_user/", nil,
		withDB(db, routers.GetTalkUsers),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var users []models.TalkUser
	decodeResponse(t, rr.Body, &users)

	checkResponseNumber(t, "talk users", len(users), 2)
}

func TestGetTalkUserByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	rr := helperServeHTTP(
		"GET", "/talk_user/"+strconvInt(tu.ID), nil,
		withDB(db, routers.GetTalkUserByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateTalkUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// TalkUser creation doesn't need body currently as it just inserts default values
	rr := helperServeHTTP(
		"POST", "/talk_user/", nil,
		withDB(db, routers.CreateTalkUser),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestDeleteTalkUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	rr := helperServeHTTP(
		"DELETE", "/talk_user/"+strconvInt(tu.ID), nil,
		withDB(db, routers.DeleteTalkUser),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindTalkUserByID(db, tu.ID)
	verifyDeleted(t, err, "talk user")
}
