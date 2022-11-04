package main

import (
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	helloController := controller.NewHelloController()

	helloController.Route(r)

	port := ":3000"
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
