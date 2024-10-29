package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"server/db"
	"server/handlers"
	kaf "server/kafka_client"
)

func main() {

	db := db.ConnectToDB()
	defer db.Close()

	kafkaClient, err := kaf.NewClient()
	if err != nil {
		log.Fatalf("Failed to connect to kafka: %v", err)
	}
	defer kafkaClient.Close()

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	router.Route("/api", func(r chi.Router) {
		r.Get("/users", handlers.ListUsersHandler(db))
		r.Get("/users/{userId}", handlers.RetrieveUserHandler(db))
		r.Post("/users", handlers.CreateUserHandler(db, kafkaClient))
		r.Patch("/users/{userId}", handlers.UpdateUserHandler(db))
		r.Delete("/users/{userId}", handlers.DeleteUserHandler(db))
	})

	serverPort := os.Getenv("SERVER_PORT")

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: router,
	}

	fmt.Printf("Starting server on localhost:%s\n", serverPort)

	err = server.ListenAndServe()
	if err == http.ErrServerClosed {
		fmt.Println("Server shutting down")
	}
	if err != nil {
		fmt.Println(err)
	}
}
