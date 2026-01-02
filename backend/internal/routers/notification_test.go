package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetNotifications(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n1 := createNewNotification()
	n2 := createNewNotification()
	saveNotification(t, db, n1)
	saveNotification(t, db, n2)

	rr := helperServeHTTP(
		"GET", "/notification/", nil,
		withDB(db, routers.GetNotifications),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var notifications []models.Notification
	decodeResponse(t, rr.Body, &notifications)

	checkResponseNumber(t, "notifications", len(notifications), 2)
}

func TestGetNotificationByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n := createNewNotification()
	saveNotification(t, db, n)

	rr := helperServeHTTP(
		"GET", "/notification/"+strconvInt(n.ID), nil,
		withDB(db, routers.GetNotificationByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateNotification(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newNotification := createNewNotification()
	body, _ := json.Marshal(newNotification)

	rr := helperServeHTTP(
		"POST", "/notification/", bytes.NewBuffer(body),
		withDB(db, routers.CreateNotification),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateNotification(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n := createNewNotification()
	saveNotification(t, db, n)

	updatedNotification := createNewNotification()
	updatedNotification.Title = "Updated Title"
	updatedNotification.Content = "Updated Content"
	body, _ := json.Marshal(updatedNotification)

	rr := helperServeHTTP(
		"PUT", "/notification/"+strconvInt(n.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateNotification),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindNotificationByID(db, n.ID)
	checkUpdated(t, "title", "Updated Title", updated.Title)
}

func TestDeleteNotification(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n := createNewNotification()
	saveNotification(t, db, n)

	rr := helperServeHTTP(
		"DELETE", "/notification/"+strconvInt(n.ID), nil,
		withDB(db, routers.DeleteNotification),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindNotificationByID(db, n.ID)
	verifyDeleted(t, err, "notification")
}
