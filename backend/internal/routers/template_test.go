package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetTemplates(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tmpl1 := createNewTemplate()
	tmpl2 := createNewTemplate()
	saveTemplate(t, db, tmpl1)
	saveTemplate(t, db, tmpl2)

	rr := helperServeHTTP(
		"GET", "/template/", nil,
		withDB(db, routers.GetTemplates),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var templates []models.Template
	decodeResponse(t, rr.Body, &templates)

	checkResponseNumber(t, "templates", len(templates), 2)
}

func TestGetTemplateByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tmpl := createNewTemplate()
	saveTemplate(t, db, tmpl)

	rr := helperServeHTTP(
		"GET", "/template/"+strconvInt(tmpl.ID), nil,
		withDB(db, routers.GetTemplateByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)
}

func TestCreateTemplate(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newTemplate := createNewTemplate()
	body, _ := json.Marshal(newTemplate)

	rr := helperServeHTTP(
		"POST", "/template/", bytes.NewBuffer(body),
		withDB(db, routers.CreateTemplate),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateTemplate(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tmpl := createNewTemplate()
	saveTemplate(t, db, tmpl)

	updatedTemplate := createNewTemplate()
	updatedTemplate.TemplateURL = "http://example.com/updated_template.jpg"
	body, _ := json.Marshal(updatedTemplate)

	rr := helperServeHTTP(
		"PUT", "/template/"+strconvInt(tmpl.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateTemplate),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindTemplateByID(db, tmpl.ID)
	checkUpdated(
		t,
		"templateURL",
		"http://example.com/updated_template.jpg",
		updated.TemplateURL,
	)
}

func TestDeleteTemplate(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	tmpl := createNewTemplate()
	saveTemplate(t, db, tmpl)

	rr := helperServeHTTP(
		"DELETE", "/template/"+strconvInt(tmpl.ID), nil,
		withDB(db, routers.DeleteTemplate),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindTemplateByID(db, tmpl.ID)
	verifyDeleted(t, err, "template")
}
