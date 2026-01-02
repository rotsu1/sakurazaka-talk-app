package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetFanletters(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Fanletter Receiver")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	f1 := createNewFanletter(m.ID, tu.ID)
	saveFanletter(t, db, f1)
	f2 := createNewFanletter(m.ID, tu.ID)
	saveFanletter(t, db, f2)

	rr := helperServeHTTP(
		"GET", "/fanletter/", nil,
		withDB(db, routers.GetFanletters),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var fanletters []models.Fanletter
	decodeResponse(t, rr.Body, &fanletters)

	checkResponseNumber(t, "fanletters", len(fanletters), 2)
}

func TestGetFanletterByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Fanletter Receiver")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	f := createNewFanletter(m.ID, tu.ID)
	saveFanletter(t, db, f)

	rr := helperServeHTTP(
		"GET", "/fanletter/"+strconvInt(f.ID), nil,
		withDB(db, routers.GetFanletterByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateFanletter(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Fanletter Receiver")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	newFanletter := createNewFanletter(m.ID, tu.ID)
	body, _ := json.Marshal(newFanletter)

	rr := helperServeHTTP(
		"POST", "/fanletter/", bytes.NewBuffer(body),
		withDB(db, routers.CreateFanletter),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestDeleteFanletter(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Fanletter Receiver")
	saveMember(t, db, m)
	tu := createNewTalkUser()
	saveTalkUser(t, db, tu)

	f := createNewFanletter(m.ID, tu.ID)
	saveFanletter(t, db, f)

	rr := helperServeHTTP(
		"DELETE", "/fanletter/"+strconvInt(f.ID), nil,
		withDB(db, routers.DeleteFanletter),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindFanletterByID(db, f.ID)
	verifyDeleted(t, err, "fanletter")
}
