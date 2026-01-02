package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterMemberRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/member/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/member/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetMembers(db, w, r)
			} else {
				GetMemberByID(db, w, r)
			}
		case http.MethodPost:
			CreateMember(db, w, r)
		case http.MethodPut:
			UpdateMember(db, w, r)
		case http.MethodDelete:
			DeleteMember(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetMembers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	members, err := models.GetAllMembers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(members); err != nil {
		http.Error(w, "Failed to encode members", http.StatusInternalServerError)
	}
}

func GetMemberByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/member/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	m, err := models.FindMemberByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Member not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		http.Error(w, "Failed to encode member", http.StatusInternalServerError)
	}
}

func CreateMember(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var m models.Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := m.Save(db); err != nil {
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
	if err := json.NewEncoder(w).Encode(m); err != nil {
		http.Error(w, "Failed to encode member", http.StatusInternalServerError)
	}
}

func UpdateMember(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/member/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	if _, err := models.FindMemberByID(db, id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Member not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var m models.Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	m.ID = id

	if err := m.Update(db); err != nil {
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

func DeleteMember(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/member/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM member WHERE id = $1`, id)
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
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
