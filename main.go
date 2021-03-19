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
	client, err := initializeClient()
	if err != nil {
		log.Fatal("Could not create client!", err.Error())
		return
	}
	handleRequests(client)
}

func handleRequests(client *db.Client) {
	requestRouter := mux.NewRouter().StrictSlash(true)

	messageHandler := api.HandleWithClient(client, api.HandleMessage)
	userInfoHandler := api.HandleWithClient(client, api.HandleUserInfo)

	requestRouter.HandleFunc("/message", messageHandler).Methods("POST")
	requestRouter.HandleFunc("/userinfo", userInfoHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", requestRouter))
}

func initializeClient() (*db.Client, error) {
	app, err := firebase.CreateFirebaseApp("key/adminsdk_key.json")
	if err != nil {
		return nil, err
	}
	return firebase.CreateFirestoreClient(app)
}
