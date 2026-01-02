package routers

import (
	"backend/internal/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func registerMessageReadRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc(
		"/message_read/",
		func(w http.ResponseWriter, r *http.Request,
		) {
			path := r.URL.Path[len("/message_read/"):]

			switch r.Method {
			case http.MethodGet:
				if path == "" {
					GetMessageRead(db, w, r)
				} else {
					GetMessageReadByIDs(db, w, r)
				}
			case http.MethodPost:
				CreateMessageReadByID(db, w, r)
			case http.MethodPut:
				UpdateMessageRead(db, w, r)
			case http.MethodDelete:
				DeleteMessageRead(db, w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
}

func GetMessageRead(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	messageRead, err := models.GetAllMessageReads(db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messageRead); err != nil {
		http.Error(w, "Failed to encode message_read", http.StatusInternalServerError)
	}
}

func GetMessageReadByIDs(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		http.Error(
			w,
			"Invalid URL, must be /message_read/{messageID}/{userID}",
			http.StatusBadRequest,
		)
	}

	messageID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	messageRead, err := models.FindMessageReadByIDs(db, messageID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messageRead); err != nil {
		http.Error(w, "Failed to encode message_read", http.StatusInternalServerError)
	}
}

func CreateMessageReadByID(
	db *sql.DB,
	w http.ResponseWriter,
	r *http.Request,
) {
	var mr models.MessageRead
	if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := mr.Save(db); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(mr); err != nil {
		http.Error(w, "Failed to encode message_read", http.StatusInternalServerError)
	}
}

func UpdateMessageRead(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var mr models.MessageRead
	if err := json.NewDecoder(r.Body).Decode(&mr); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := mr.Update(db); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Message Read not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteMessageRead(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	messageID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(
			w,
			"Invalid URL, must be /message_read/{messageID}/{userID}",
			http.StatusBadRequest,
		)
	}
	userID, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(
			w,
			"Invalid URL, must be /message_read/{messageID}/{userID}",
			http.StatusBadRequest,
		)
	}

	mr := &models.MessageRead{
		MessageID: messageID,
		UserID:    userID,
	}

	if err := mr.Delete(db); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Message Read not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
