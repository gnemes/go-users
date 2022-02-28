package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	di "github.com/gnemes/go-users/Infrastructure/Services/Di"
	middleware "github.com/gnemes/go-users/Infrastructure/Middleware"
	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
)

type Server struct {
	Router *mux.Router
}

type RouteClousure func(txName string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)

func IdentityClousure(txName string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return f
}

func Listen() {
	routesClousure := IdentityClousure

	s := &Server{
		Router: mux.NewRouter().StrictSlash(true),
	}

	log.Fatal(http.ListenAndServe(":8081", s.Routes(routesClousure)))
}

func (s *Server) Routes(cl RouteClousure) http.Handler {
	s.Router.Use(middleware.ApplicationJsonMiddleware)

	usersRouter := s.Router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/{id}", cl("[GET]/users/{id}", di.GetInstance().Get("GetUserControllerHttp").(*controllerhttp.Get).Execute)).Methods("GET")

	return middleware.TrimSlashMiddleware(s.Router)
}