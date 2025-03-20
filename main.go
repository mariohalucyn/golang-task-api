package main

import (
	"github.com/gorilla/mux"
	"github.com/mariohalucyn/todo-app/handlers"
	"github.com/mariohalucyn/todo-app/initializers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
	initializers.LoadAndParsePrivateKey()
	initializers.LoadAndParsePublicKey()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/api/verify", handlers.Verify).Methods("GET")
	router.HandleFunc("/api/update-user", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/authorization", handlers.Authorization).Methods("GET")
	router.HandleFunc("/api/logout", handlers.Logout).Methods("GET")
	router.HandleFunc("/api/create-todo", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/get-todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/api/update-todo/{id:[0-9]+}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/delete-todo/{id:[0-9]+}", handlers.DeleteTodo).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_ADDRESS")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
