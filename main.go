package main

import (
	"firebase.google.com/go/db"
	"github.com/gorilla/mux"
	api "github.com/vshnv/messagelogger/api"
	"github.com/vshnv/messagelogger/firebase"
	"log"
	"net/http"
)

func main() {
	app, err := firebase.CreateFirebaseApp("key/adminsdk_key.json")
	if err != nil {
		log.Fatal("Could not find key!", err.Error())
		return
	}
	client, err := firebase.CreateFirestoreClient(app)
	if err != nil {
		log.Fatal("Could not create client!", err.Error())
		return
	}
	handleRequests(client)
}

func handleRequests(client *db.Client) {
	myRouter := mux.NewRouter().StrictSlash(true)

	messageHandler := api.HandleWithClient(client, api.HandleMessage)
	userInfoHandler := api.HandleWithClient(client, api.HandleUserInfo)

	myRouter.HandleFunc("/message", messageHandler).Methods("POST")
	myRouter.HandleFunc("/userinfo", userInfoHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
