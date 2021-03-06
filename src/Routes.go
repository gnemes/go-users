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
	// Get App middlewares
	jsonApiHeaderMiddleware := container.Get("JsonApiHeaderMiddleware").(*middleware.JsonApiHeaderMiddleware)
	trimSlashMiddleware := container.Get("TrimSlashMiddleware").(*middleware.TrimSlashMiddleware)

	s.Use(trimSlashMiddleware.Execute)
	s.Use(jsonApiHeaderMiddleware.Execute)

	// Users router
	usersRouter := s.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/{id}", fetchHandleFunc(container, "GetUserControllerHttp")).Methods("GET")
	
	return s
}

func fetchHandleFunc(container di.Container, controller string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requestContainer, _ := container.SubContainer()
		defer requestContainer.Delete()

		// Get Request middlewares
		credentialsMiddleware := requestContainer.Get("CredentialsMiddleware").(*middleware.CredentialsMiddleware)
		requestIDMiddleware := requestContainer.Get("RequestIDMiddleware").(*middleware.RequestIDMiddleware)
		requestQueryParserMiddleware := requestContainer.Get("RequestQueryParserMiddleware").(*middleware.RequestQueryParserMiddleware)
		requestedEntityIDMiddleware := requestContainer.Get("RequestedEntityIDMiddleware").(*middleware.RequestedEntityIDMiddleware)
		
		// Get controller
		controllerInstance := requestContainer.Get(controller).(*controllerhttp.Get)

		// Compose middlewares + controller handler
		handler := requestQueryParserMiddleware.Execute(
						requestedEntityIDMiddleware.Execute(
							requestIDMiddleware.Execute(
								credentialsMiddleware.Execute(
									controllerInstance.Execute,
								),
							),
						),
					)

		handler(w, r)
	}
}