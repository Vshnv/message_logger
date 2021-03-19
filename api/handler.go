package api

import (
	"firebase.google.com/go/db"
	"net/http"
)

type RouteHandler = func(w http.ResponseWriter, r *http.Request)
type ClientRouteHandler = func(w http.ResponseWriter, r *http.Request, client *db.Client)
type DatabaseRouteHandler = func(w http.ResponseWriter, r *http.Request, database string)

func HandleWithFirebaseClient(client *db.Client, handler ClientRouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, client)
	}
}

func HandleWithSql(databasePath string, handler DatabaseRouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, databasePath)
	}
}
