package routers_test

import (
	"backend/internal/models"
	"backend/internal/routers"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetBlogs(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup dependency
	m := createNewMember("Blog Writer")
	saveMember(t, db, m)

	b1 := createNewBlog(m.ID, "Addiction")
	saveBlog(t, db, b1)
	b2 := createNewBlog(m.ID, "Unhappy Birthday構文")
	saveBlog(t, db, b2)

	rr := helperServeHTTP(
		"GET", "/blog/", nil,
		withDB(db, routers.GetBlogs),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var blogs []models.Blog
	decodeResponse(t, rr.Body, &blogs)

	checkResponseNumber(t, "blogs", len(blogs), 2)
}

func TestGetBlogByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Blog Writer")
	saveMember(t, db, m)

	b := createNewBlog(m.ID, "Addiction")
	saveBlog(t, db, b)

	rr := helperServeHTTP(
		"GET", "/blog/"+strconvInt(b.ID), nil,
		withDB(db, routers.GetBlogByID),
	)

	checkStatusCode(t, rr.Code, http.StatusOK)

	var blog models.Blog
	decodeResponse(t, rr.Body, &blog)
}

func TestCreateBlog(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Blog Writer")
	saveMember(t, db, m)

	newBlog := createNewBlog(m.ID, "New Blog Title")
	body, _ := json.Marshal(newBlog)

	rr := helperServeHTTP(
		"POST", "/blog/", bytes.NewBuffer(body),
		withDB(db, routers.CreateBlog),
	)

	checkStatusCode(t, rr.Code, http.StatusCreated)
}

func TestUpdateBlog(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Blog Writer")
	saveMember(t, db, m)

	b := createNewBlog(m.ID, "Addiction")
	saveBlog(t, db, b)

	updatedBlog := createNewBlog(m.ID, "Updated Title")
	updatedBlog.ID = b.ID
	body, _ := json.Marshal(updatedBlog)

	rr := helperServeHTTP(
		"PUT", "/blog/"+strconvInt(b.ID), bytes.NewBuffer(body),
		withDB(db, routers.UpdateBlog),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	updated, _ := models.FindBlogByID(db, b.ID)
	checkUpdated(t, "title", "Updated Title", updated.Title)
}

func TestDeleteBlog(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	m := createNewMember("Blog Writer")
	saveMember(t, db, m)

	b := createNewBlog(m.ID, "Addiction")
	saveBlog(t, db, b)

	rr := helperServeHTTP(
		"DELETE", "/blog/"+strconvInt(b.ID), nil,
		withDB(db, routers.DeleteBlog),
	)

	checkStatusCode(t, rr.Code, http.StatusNoContent)

	_, err := models.FindBlogByID(db, b.ID)
	verifyDeleted(t, err, "blog")
}
