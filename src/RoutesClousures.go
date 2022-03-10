package main

import (
	"net/http"

	// "github.com/newrelic/go-agent/v3/newrelic"
)

type RouteClousure func(txName string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)

// do nothing
func IdentityClousure(txName string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return f
}

// do nothing
/*
func NewRelicClousure(app *newrelic.Application) func(txName string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(txName string, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			txn := app.StartTransaction(txName)
			defer txn.End()
			f(w, r)
			txn.SetWebRequestHTTP(r)
			txn.SetWebResponse(w)
		}
	}
}
*/