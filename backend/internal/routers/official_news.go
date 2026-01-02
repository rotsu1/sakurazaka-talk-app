package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterOfficialNewsRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/official_news/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/official_news/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetOfficialNews(db, w, r)
			} else {
				GetOfficialNewsByID(db, w, r)
			}
		case http.MethodPost:
			CreateOfficialNews(db, w, r)
		case http.MethodPut:
			UpdateOfficialNews(db, w, r)
		case http.MethodDelete:
			DeleteOfficialNews(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetOfficialNews(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	officialNews, err := models.GetAllOfficialNews(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(officialNews); err != nil {
		http.Error(w, "Failed to encode official news", http.StatusInternalServerError)
	}
}

func GetOfficialNewsByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/official_news/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid official news ID", http.StatusBadRequest)
		return
	}

	on, err := models.FindOfficialNewsByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "OfficialNews not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(on); err != nil {
		http.Error(w, "Failed to encode official news", http.StatusInternalServerError)
	}
}

func CreateOfficialNews(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var on models.OfficialNews
	if err := json.NewDecoder(r.Body).Decode(&on); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := on.Save(db); err != nil {
		switch models.ClassifyDBError(err) {
		case models.ErrInvalidReference, models.ErrInvalidData:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(on); err != nil {
		http.Error(w, "Failed to encode official news", http.StatusInternalServerError)
	}
}

func UpdateOfficialNews(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/official_news/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid official news ID", http.StatusBadRequest)
		return
	}

	if _, err := models.FindOfficialNewsByID(db, id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "OfficialNews not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var on models.OfficialNews
	if err := json.NewDecoder(r.Body).Decode(&on); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	on.ID = id

	if err := on.Update(db); err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteOfficialNews(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/official_news/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid official news ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM official_news WHERE id = $1`, id)
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
		http.Error(w, "OfficialNews not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
