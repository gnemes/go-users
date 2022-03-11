package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	di "github.com/sarulabs/di/v2"

	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	middleware "github.com/gnemes/go-users/Infrastructure/Middleware"
)

func listen(container di.Container) {
	Router := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":8081", routes(container, Router)))
}

func routes(container di.Container, s *mux.Router) http.Handler {
	// Get middlewares
	jsonApiHeaderMiddleware := container.Get("JsonApiHeaderMiddleware").(*middleware.JsonApiHeaderMiddleware)
	trimSlashMiddleware := container.Get("TrimSlashMiddleware").(*middleware.TrimSlashMiddleware)

	s.Use(trimSlashMiddleware.TrimSlashMiddleware)
	s.Use(jsonApiHeaderMiddleware.JsonApiHeaderMiddleware)

	// /users router
	usersRouter := s.PathPrefix("/users").Subrouter()

	// GET - /users/{id}
	usersRouter.HandleFunc("/{id}", fetchHandleFunc(container, "GetUserControllerHttp")).Methods("GET")
	
	return s
}

func fetchHandleFunc(container di.Container, controller string) func(w http.ResponseWriter, r *http.Request) {
	requestContainer, _ := container.SubContainer()
	defer requestContainer.Delete()

	credentialsMiddleware := requestContainer.Get("CredentialsMiddleware").(*middleware.CredentialsMiddleware)
	
	controllerInstance := requestContainer.Get(controller).(*controllerhttp.Get)
	return credentialsMiddleware.CredentialsMiddleware(controllerInstance.Execute)
}