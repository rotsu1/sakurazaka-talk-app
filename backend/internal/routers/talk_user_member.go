package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterTalkUserMemberRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/talk_user_member/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/talk_user_member/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetTalkUserMembers(db, w, r)
			} else {
				GetTalkUserMemberByID(db, w, r)
			}
		case http.MethodPost:
			CreateTalkUserMember(db, w, r)
		case http.MethodPut:
			UpdateTalkUserMember(db, w, r)
		case http.MethodDelete:
			DeleteTalkUserMember(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetTalkUserMembers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	talkUserMembers, err := models.GetAllTalkUserMembers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(talkUserMembers); err != nil {
		http.Error(w, "Failed to encode talk user members", http.StatusInternalServerError)
	}
}

func GetTalkUserMemberByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/talk_user_member/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid talk user member ID", http.StatusBadRequest)
		return
	}

	tum, err := models.FindTalkUserMemberByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "TalkUserMember not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tum); err != nil {
		http.Error(w, "Failed to encode talk user member", http.StatusInternalServerError)
	}
}

func CreateTalkUserMember(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var tum models.TalkUserMember
	if err := json.NewDecoder(r.Body).Decode(&tum); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.EnsureTalkUserExists(db, tum.UserID); err != nil {
		http.Error(w, "Failed to ensure user exists: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tum.Save(db); err != nil {
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
	if err := json.NewEncoder(w).Encode(tum); err != nil {
		http.Error(w, "Failed to encode talk user member", http.StatusInternalServerError)
	}
}

func UpdateTalkUserMember(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/talk_user_member/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid talk user member ID", http.StatusBadRequest)
		return
	}

	existing, err := models.FindTalkUserMemberByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "TalkUserMember not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var tum models.TalkUserMember
	if err := json.NewDecoder(r.Body).Decode(&tum); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	tum.ID = id
	tum.UserID = existing.UserID     // Preserve UserID
	tum.MemberID = existing.MemberID // Preserve MemberID

	if err := tum.Update(db); err != nil {
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

func DeleteTalkUserMember(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/talk_user_member/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid talk user member ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM talk_user_member WHERE id = $1`, id)
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
		http.Error(w, "TalkUserMember not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
