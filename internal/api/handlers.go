package api

import (
	"encoding/json"
	"log"
	"net/http"

	"taskManager/internal/database"
	"taskManager/internal/models"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, err := database.DB.Query("SELECT id, title, description, status, priority, user_id, created_at, updated_at FROM tasks")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		tasks := []models.Task{}
		for rows.Next() {
			var t models.Task
			if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.UserID, &t.CreatedAt, &t.UpdatedAt); err != nil {
				log.Println("Error scanning row:", err)
				continue
			}
			tasks = append(tasks, t)
		}
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
