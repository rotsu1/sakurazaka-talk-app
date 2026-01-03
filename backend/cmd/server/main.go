package main

import (
	"backend/internal/db"
	"backend/internal/routers"
	"log"
	"net/http"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Println("DB connection failed:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	routers.RegisterRoutes(mux, db)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
