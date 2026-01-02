package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterStaffRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/staff/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/staff/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetStaffs(db, w, r)
			} else {
				GetStaffByID(db, w, r)
			}
		case http.MethodPost:
			CreateStaff(db, w, r)
		case http.MethodPut:
			UpdateStaff(db, w, r)
		case http.MethodDelete:
			DeleteStaff(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetStaffs(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	staffs, err := models.GetAllStaff(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(staffs); err != nil {
		http.Error(w, "Failed to encode staffs", http.StatusInternalServerError)
	}
}

func GetStaffByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/staff/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid staff ID", http.StatusBadRequest)
		return
	}

	s, err := models.FindStaffByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Staff not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, "Failed to encode staff", http.StatusInternalServerError)
	}
}

func CreateStaff(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var s models.Staff
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.Save(db); err != nil {
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
	if err := json.NewEncoder(w).Encode(s); err != nil {
		http.Error(w, "Failed to encode staff", http.StatusInternalServerError)
	}
}

func UpdateStaff(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/staff/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid staff ID", http.StatusBadRequest)
		return
	}

	if _, err := models.FindStaffByID(db, id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Staff not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var s models.Staff
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	s.ID = id

	if err := s.Update(db); err != nil {
		switch models.ClassifyDBError(err) {
		case models.ErrInvalidReference, models.ErrInvalidData:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteStaff(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/staff/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid staff ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM staff WHERE id = $1`, id)
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
		http.Error(w, "Staff not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
