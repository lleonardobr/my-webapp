package main

import (
	"fmt"
	"log"
	"my-webapp/handlers" // Importação corrigida
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/reset", handlers.ResetHandler)
	http.HandleFunc("/auth/google", handlers.GoogleAuthHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
