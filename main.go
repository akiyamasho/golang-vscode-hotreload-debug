package main

import (
	"log"
	"net/http"

	"github.com/akiyamasho/golang-vscode-hotreload-debug/handlers"
)

const PORT = "8000"

func main() {
	http.HandleFunc("/health", handlers.GetHealth)

	log.Println("Server running on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
