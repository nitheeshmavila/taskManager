package main

import (
	"fmt"
	"log"
	"net/http"

	"taskManager/internal/api"
	"taskManager/internal/database"
)

func main() {
	fmt.Println("Task Management API")
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	http.HandleFunc("/tasks", api.TasksHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
