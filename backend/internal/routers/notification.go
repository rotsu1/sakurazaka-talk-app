package routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/models"
)

func RegisterNotificationRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/notification/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/notification/"):]

		switch r.Method {
		case http.MethodGet:
			if path == "" {
				GetNotifications(db, w, r)
			} else {
				GetNotificationByID(db, w, r)
			}
		case http.MethodPost:
			CreateNotification(db, w, r)
		case http.MethodPut:
			UpdateNotification(db, w, r)
		case http.MethodDelete:
			DeleteNotification(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func GetNotifications(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	notifications, err := models.GetAllNotifications(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notifications); err != nil {
		http.Error(w, "Failed to encode notifications", http.StatusInternalServerError)
	}
}

func GetNotificationByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/notification/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	n, err := models.FindNotificationByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Notification not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(n); err != nil {
		http.Error(w, "Failed to encode notification", http.StatusInternalServerError)
	}
}

func CreateNotification(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var n models.Notification
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := n.Save(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(n); err != nil {
		http.Error(w, "Failed to encode notification", http.StatusInternalServerError)
	}
}

func UpdateNotification(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/notification/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	if _, err := models.FindNotificationByID(db, id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Notification not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var n models.Notification
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	n.ID = id

	if err := n.Update(db); err != nil {
		http.Error(w, "Update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteNotification(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/notification/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(`DELETE FROM notification WHERE id = $1`, id)
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
		http.Error(w, "Notification not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
