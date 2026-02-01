package main

import (
	"log"
	"net/http"
	"time"

	"YeahMusic/internal/handlers"
	"YeahMusic/internal/services"
)

func main() {
	store := services.NewStore()
	services.Seed(store)

	cleanupStop := make(chan struct{})
	go services.SessionCleanupWorker(store, 30*time.Second, cleanupStop)

	app := handlers.NewApp(store)
	mux := handlers.NewRouter(app)

	addr := ":8080"
	log.Println("API listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		close(cleanupStop)
		log.Fatal(err)
	}
}
