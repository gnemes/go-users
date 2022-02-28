package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func Listen() {
	s := &Server{
		Router: mux.NewRouter().StrictSlash(true),
	}

	log.Fatal(http.ListenAndServe(":8081", s.Routes()))
}

func (s *Server) Routes() *mux.Router {
	// s.Router.Use(middleware.ApplicationJsonMiddleware)
	// s.Router.Use(middleware.RequestIDMiddleware)

	// Health check
	//s.Router.HandleFunc("/healthcheck", healthactions.Check).Methods("GET", "POST")

	return s.Router
}