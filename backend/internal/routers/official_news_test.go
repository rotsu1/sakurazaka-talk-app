package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetOfficialNews(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n1 := createNewOfficialNews()
	n2 := createNewOfficialNews()
	saveOfficialNews(t, db, n1)
	saveOfficialNews(t, db, n2)

	rr := helperServeHTTP(
		"GET", "/official_news/", nil,
		withDB(db, routers.GetOfficialNews),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var news []models.OfficialNews
	decodeResponse(t, rr.Body, &news)

	checkResponseNumber(t, "official news", len(news), 2)
}

func TestGetOfficialNewsByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n := createNewOfficialNews()
	saveOfficialNews(t, db, n)

	rr := helperServeHTTP(
		"GET", "/official_news/"+strconvInt(n.ID), nil,
		withDB(db, routers.GetOfficialNewsByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateOfficialNews(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newNews := createNewOfficialNews()
	body, _ := json.Marshal(newNews)

	rr := helperServeHTTP(
		"POST", "/official_news/", bytes.NewBuffer(body),
		withDB(db, routers.CreateOfficialNews),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateOfficialNews(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n := createNewOfficialNews()
	saveOfficialNews(t, db, n)

	tag := "Update"
	updatedNews := createNewOfficialNews()
	updatedNews.Title = "Updated Official News"
	updatedNews.Tag = &tag
	body, _ := json.Marshal(updatedNews)

	rr := helperServeHTTP(
		"PUT", "/official_news/"+strconvInt(n.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateOfficialNews),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindOfficialNewsByID(db, n.ID)
	checkUpdated(t, "title", "Updated Official News", updated.Title)
}

func TestDeleteOfficialNews(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	n := createNewOfficialNews()
	saveOfficialNews(t, db, n)

	rr := helperServeHTTP(
		"DELETE", "/official_news/"+strconvInt(n.ID), nil,
		withDB(db, routers.DeleteOfficialNews),
	)
	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindOfficialNewsByID(db, n.ID)
	verifyDeleted(t, err, "official news")
}
