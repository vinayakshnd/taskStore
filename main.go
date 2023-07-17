package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vinayakshnd/taskStore/task"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", task.ListTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", task.GetTask).Methods("GET")
	router.HandleFunc("/tasks", task.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", task.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", task.DeleteTask).Methods("DELETE")

	// Cobra CLI to start port
	endpoint := "0.0.0.0:8083"
	log.Printf("Starting REST server on %s...\n", endpoint)
	log.Fatal(http.ListenAndServe(endpoint, router))
}
