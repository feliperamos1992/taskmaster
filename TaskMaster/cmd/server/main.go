package main

import (
	"TaskMaster/internal/db"
	"log"
	"net/http"

	"TaskMaster/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/user/create", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/find/{id}", handlers.GetUserId).Methods("GET")

	// Task Routes
	router.HandleFunc("/task/create", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/task/find", handlers.FindTask).Methods("GET")

	//User Tasks Routers
	router.HandleFunc("/user-task/find/{userId}", handlers.GetUserAndTasks).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
