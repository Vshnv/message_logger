package api

import (
	"firebase.google.com/go/db"
	"net/http"
)

func HandleWithClient(client *db.Client, handler ClientRouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, client)
	}
}
