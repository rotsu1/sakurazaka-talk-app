package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterBlogRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/blog/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetBlogs(db, w, r)
			} else {
				GetBlogByID(db, w, r)
			}
		case http.MethodPost:
			CreateBlog(db, w, r)
		case http.MethodPut:
			UpdateBlog(db, w, r)
		case http.MethodDelete:
			DeleteBlog(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetBlogs(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	blogs, err := models.GetAllBlogs(db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		http.Error(w, "Failed to encode blogs", http.StatusInternalServerError)
	}
}

func GetBlogByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/blog/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	b, err := models.FindBlogByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Blog not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, "Failed to encode blog", http.StatusInternalServerError)
	}
}

func CreateBlog(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var b models.Blog
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := b.Save(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, "Failed to encode blog", http.StatusInternalServerError)
	}
}

func UpdateBlog(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/blog/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	var b models.Blog
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		`UPDATE blog
		 SET title = $1, content = $2, status = $3, verified_by = $4, verified_at = $5, updated_at = NOW()
		 WHERE id = $6`,
		b.Title,
		b.Content,
		b.Status,
		b.VerifiedBy,
		b.VerifiedAt,
		id,
	)
	if err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Could not get affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteBlog(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/blog/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM blog WHERE id = $1`, id)
	if err != nil {
		http.Error(w, "Delete failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Could not get affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
