package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetStaffs(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	s1 := createNewStaff(nil, "manager", "user1")
	s2 := createNewStaff(nil, "manager", "user2")
	s3 := createNewStaff(nil, "manager", "user3")
	saveStaff(t, db, s1)
	saveStaff(t, db, s2)
	saveStaff(t, db, s3)

	rr := helperServeHTTP(
		"GET", "/staff/", nil,
		withDB(db, routers.GetStaffs),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var staffs []models.Staff
	decodeResponse(t, rr.Body, &staffs)

	checkResponseNumber(t, "staffs", len(staffs), 3)
}

func TestGetStaffByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	s := createNewStaff(nil, "admin", "admin_user")
	saveStaff(t, db, s)

	rr := helperServeHTTP(
		"GET", "/staff/"+strconvInt(s.ID), nil,
		withDB(db, routers.GetStaffByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var staff models.Staff
	decodeResponse(t, rr.Body, &staff)
}

func TestCreateStaff(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newStaff := createNewStaff(nil, "manager", "new_manager")
	body, _ := json.Marshal(newStaff)

	rr := helperServeHTTP(
		"POST", "/staff/", bytes.NewBuffer(body),
		withDB(db, routers.CreateStaff),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateStaff(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	s := createNewStaff(nil, "manager", "old_username")
	saveStaff(t, db, s)

	updatedStaff := createNewStaff(nil, "admin", "updated_username")
	body, _ := json.Marshal(updatedStaff)

	rr := helperServeHTTP(
		"PUT", "/staff/"+strconvInt(s.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateStaff),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindStaffByID(db, s.ID)
	checkUpdated(t, "username", "updated_username", updated.Username)
	checkUpdated(t, "role", "admin", updated.Role)
}

func TestDeleteStaff(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	s := createNewStaff(nil, "manager", "to_delete")
	saveStaff(t, db, s)

	rr := helperServeHTTP(
		"DELETE", "/staff/"+strconvInt(s.ID), nil,
		withDB(db, routers.DeleteStaff),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindStaffByID(db, s.ID)
	verifyDeleted(t, err, "staff")
}
