package main

import (
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/controller"
	"github.com/erikrios/my-story-dummy/service"
	cfs "github.com/erikrios/my-story-dummy/util/fs"
	_ "github.com/erikrios/my-story-dummy/validation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var fs cfs.FS = cfs.NewLocalFS()

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if !fs.IsExists("data") {
		if err := fs.CreateDir("data"); err != nil {
			log.Fatal(err)
		}
	}

	if !fs.IsExists("data/.gitkeep") {
		if err := fs.CreateFile("data/.gitkeep"); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	helloController := controller.NewHelloController()
	customController := controller.NewCustomController()

	helloController.Route(r)
	customController.Route(r)

	apiRouter := chi.NewRouter()

	groupService := service.NewGroupServiceImpl(fs)

	groupController := controller.NewGroupController(groupService)
	groupController.Route(apiRouter)

	r.Mount("/api/v1", apiRouter)

	port := ":3000"
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
