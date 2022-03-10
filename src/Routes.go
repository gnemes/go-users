package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	di "github.com/sarulabs/di/v2"

	middleware "github.com/gnemes/go-users/Infrastructure/Middleware"
	// controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
)

func listen(container di.Container) {
	routesClousure := IdentityClousure

	Router := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":8081", routes(container, Router, routesClousure)))
}

func routes(container di.Container, s *mux.Router, cl RouteClousure) http.Handler {
	requestContainer, _ := container.SubContainer()
	defer requestContainer.Delete()

	// Get middlewares
	credentialsMiddleware := requestContainer.Get("CredentialsMiddleware").(*middleware.CredentialsMiddleware)
	jsonApiHeaderMiddleware := requestContainer.Get("JsonApiHeaderMiddleware").(*middleware.JsonApiHeaderMiddleware)
	trimSlashMiddleware := requestContainer.Get("TrimSlashMiddleware").(*middleware.TrimSlashMiddleware)

	s.Use(jsonApiHeaderMiddleware.JsonApiHeaderMiddleware)
	s.Use(credentialsMiddleware.CredentialsMiddleware)

	/*
	usersRouter := s.Router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/{id}", cl("[GET]/users/{id}", di.GetInstance().Get("GetUserControllerHttp").(*controllerhttp.Get).Execute)).Methods("GET")
	*/

	return trimSlashMiddleware.TrimSlashMiddleware(s)
}