package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetTalkUserMembers(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Talk Member")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum1 := createNewTalkUserMember(tu.ID, m.ID)
	tum2 := createNewTalkUserMember(tu.ID, m.ID)
	saveTalkUserMember(t, db, tum1)
	saveTalkUserMember(t, db, tum2)

	rr := helperServeHTTP(
		"GET", "/talk_user_member/", nil,
		withDB(db, routers.GetTalkUserMembers),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var tums []models.TalkUserMember
	decodeResponse(t, rr.Body, &tums)

	checkResponseNumber(t, "talk user members", len(tums), 2)
}

func TestGetTalkUserMemberByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Talk Member")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum := createNewTalkUserMember(tu.ID, m.ID)
	saveTalkUserMember(t, db, tum)

	rr := helperServeHTTP(
		"GET", "/talk_user_member/"+strconvInt(tum.ID), nil,
		withDB(db, routers.GetTalkUserMemberByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateTalkUserMember(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Talk Member")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	newTum := createNewTalkUserMember(tu.ID, m.ID)
	body, _ := json.Marshal(newTum)

	rr := helperServeHTTP(
		"POST", "/talk_user_member/", bytes.NewBuffer(body),
		withDB(db, routers.CreateTalkUserMember),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateTalkUserMember(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Talk Member")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum := createNewTalkUserMember(tu.ID, m.ID)
	saveTalkUserMember(t, db, tum)

	updatedTum := createNewTalkUserMember(tu.ID, m.ID)
	updatedTum.Status = "cancelled"
	body, _ := json.Marshal(updatedTum)

	rr := helperServeHTTP(
		"PUT", "/talk_user_member/"+strconvInt(tum.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateTalkUserMember),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindTalkUserMemberByID(db, tum.ID)
	checkUpdated(t, "status", "cancelled", updated.Status)
}

func TestDeleteTalkUserMember(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Talk Member")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	tum := createNewTalkUserMember(tu.ID, m.ID)
	saveTalkUserMember(t, db, tum)

	rr := helperServeHTTP(
		"DELETE", "/talk_user_member/"+strconvInt(tum.ID), nil,
		withDB(db, routers.DeleteTalkUserMember),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindTalkUserMemberByID(db, tum.ID)
	verifyDeleted(t, err, "talk user member")
}
