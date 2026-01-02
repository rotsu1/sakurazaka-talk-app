package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetMembers(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Seed data
	member := createNewMember("Sugai")
	saveMember(t, db, member)

	rr := helperServeHTTP(
		"GET", "/member/", nil,
		withDB(db, routers.GetMembers),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var members []models.Member
	decodeResponse(t, rr.Body, &members)

	checkResponseNumber(t, "members", len(members), 1)
}

func TestGetMemberByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	member := createNewMember("Target Member")
	saveMember(t, db, member)

	tests := []struct {
		name       string
		id         string
		wantStatus int
	}{
		{"Valid ID", strconvInt(member.ID), http.StatusOK},
		{"Non-existent ID", "9999", http.StatusNotFound},
		{"Invalid ID", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := helperServeHTTP(
				"GET", "/member/"+tt.id, nil,
				withDB(db, routers.GetMemberByID),
			)

			checkStatusCode(t, rr.Code, tt.wantStatus)
		})
	}
}

func TestCreateMember(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newMember := createNewMember("New Member")
	body, _ := json.Marshal(newMember)

	rr := helperServeHTTP(
		"POST", "/member/", bytes.NewBuffer(body),
		withDB(db, routers.CreateMember),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateMember(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	member := createNewMember("Original Name")
	saveMember(t, db, member)

	updatedMember := createNewMember("Updated Name")
	body, _ := json.Marshal(updatedMember)

	rr := helperServeHTTP(
		"PUT", "/member/"+strconvInt(member.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateMember),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	// Verify update
	m, err := models.FindMemberByID(db, member.ID)
	if err != nil {
		t.Fatalf("Failed to find member: %v", err)
	}
	checkUpdated(t, "name", "Updated Name", m.Name)
}

func TestDeleteMember(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	member := createNewMember("To Delete")
	saveMember(t, db, member)

	rr := helperServeHTTP(
		"DELETE", "/member/"+strconvInt(member.ID), nil,
		withDB(db, routers.DeleteMember),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	// Verify deletion
	_, err := models.FindMemberByID(db, member.ID)
	verifyDeleted(t, err, "member")
}
