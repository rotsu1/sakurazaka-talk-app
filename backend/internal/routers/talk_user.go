package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterTalkUserRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/talk_user/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/talk_user/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetTalkUsers(db, w, r)
			} else {
				GetTalkUserByID(db, w, r)
			}
		case http.MethodPost:
			CreateTalkUser(db, w, r)
		case http.MethodDelete:
			DeleteTalkUser(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetTalkUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	talkUsers, err := models.GetAllTalkUsers(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(talkUsers); err != nil {
		http.Error(w, "Failed to encode talk users", http.StatusInternalServerError)
	}
}

func GetTalkUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/talk_user/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid talk user ID", http.StatusBadRequest)
		return
	}

	tu, err := models.FindTalkUserByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "TalkUser not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tu); err != nil {
		http.Error(w, "Failed to encode talk user", http.StatusInternalServerError)
	}
}

func CreateTalkUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var tu models.TalkUser
	if err := tu.Save(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tu); err != nil {
		http.Error(w, "Failed to encode talk user", http.StatusInternalServerError)
	}
}

func DeleteTalkUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/talk_user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid talk user ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM talk_user WHERE id = $1`, id)
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
		http.Error(w, "TalkUser not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
