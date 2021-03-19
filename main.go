package main

import (
	"database/sql"
	"firebase.google.com/go/db"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vshnv/messagelogger/api"
	"github.com/vshnv/messagelogger/firebase"
	"log"
	"net/http"
)

const MSG_DATABASE_PATH = "./db/message_log.db"

func main() {
	client, err := initializeFirebaseClient("key/admin_sdk_key.json")
	if err != nil {
		log.Fatal("Could not create client!", err.Error())
		return
	}
	err = initializeDatabaseTables()
	if err != nil {
		log.Fatal("Could not init db!", err.Error())
		return
	}
	handleRequests(client)
}

func handleRequests(client *db.Client) {
	requestRouter := mux.NewRouter().StrictSlash(true)

	messagehandler := api.HandleWithSql(MSG_DATABASE_PATH, api.HandleMessage)
	userInfoHandler := api.HandleWithFirebaseClient(client, api.HandleUserInfo)

	requestRouter.HandleFunc("/message", messagehandler).Methods("POST")
	requestRouter.HandleFunc("/userinfo", userInfoHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", requestRouter))
}

func initializeFirebaseClient(keyPath string) (*db.Client, error) {
	app, err := firebase.CreateFirebaseApp(keyPath)
	if err != nil {
		return nil, err
	}
	return firebase.CreateFirestoreClient(app)
}

func initializeDatabaseTables() error {
	database, err := sql.Open("sqlite3", "./db/message_log.db")
	if err != nil {
		fmt.Println("Error connecting to database!")
		return err
	}
	_, err = database.Exec("CREATE TABLE IF NOT EXISTS messages (from_user varchar(32), to_user varchar(32), content text, epoch INTEGER)")
	if err != nil {
		fmt.Println("Error executing table setup!")
		return err
	}
	return nil
}
