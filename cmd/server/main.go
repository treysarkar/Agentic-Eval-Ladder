package main

import (
	"log"
	"net/http"
	"github.com/treysarkar/Agentic-Eval-Ladder/internal/api"
	"github.com/treysarkar/Agentic-Eval-Ladder/internal/store"
)

func main() {
	handler := api.NewRouter(store.New())
	// will give me a new store struct defined 
	// which I could use whenever required here
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",handler))
}