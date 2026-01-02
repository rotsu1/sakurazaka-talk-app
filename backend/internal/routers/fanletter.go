package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterFanletterRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/fanletter/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/fanletter/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetFanletters(db, w, r)
			} else {
				GetFanletterByID(db, w, r)
			}
		case http.MethodPost:
			CreateFanletter(db, w, r)
		case http.MethodDelete:
			DeleteFanletter(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetFanletters(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fanletters, err := models.GetAllFanletters(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fanletters); err != nil {
		http.Error(w, "Failed to encode fanletters", http.StatusInternalServerError)
	}
}

func GetFanletterByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/fanletter/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid fanletter ID", http.StatusBadRequest)
		return
	}

	f, err := models.FindFanletterByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Fanletter not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(f); err != nil {
		http.Error(w, "Failed to encode fanletter", http.StatusInternalServerError)
	}
}

func CreateFanletter(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var f models.Fanletter
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := f.Save(db); err != nil {
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
	if err := json.NewEncoder(w).Encode(f); err != nil {
		http.Error(w, "Failed to encode fanletter", http.StatusInternalServerError)
	}
}

func DeleteFanletter(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/fanletter/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid fanletter ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM fanletter WHERE id = $1`, id)
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
		http.Error(w, "Fanletter not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
