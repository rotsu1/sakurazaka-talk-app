package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterMessageRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/message/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/message/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetMessages(db, w, r)
			} else {
				GetMessageByID(db, w, r)
			}
		case http.MethodPost:
			CreateMessage(db, w, r)
		case http.MethodPut:
			UpdateMessage(db, w, r)
		case http.MethodDelete:
			DeleteMessage(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	messages, err := models.GetAllMessages(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
	}
}

func GetMessageByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/message/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	m, err := models.FindMessageByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Message not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
	}
}

func CreateMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var m models.Message
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
		http.Error(w, "Failed to encode message", http.StatusInternalServerError)
	}
}

func UpdateMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/message/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	existing, err := models.FindMessageByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Message not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var m models.Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	m.ID = id
	m.MemberID = existing.MemberID // Preserve member_id

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

func DeleteMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/message/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM message WHERE id = $1`, id)
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
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
