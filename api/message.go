package api

import (
	"encoding/json"
	"firebase.google.com/go/db"
	"fmt"
	"github.com/vshnv/messagelogger/firebase"
	"net/http"
)

type RouteHandler = func(w http.ResponseWriter, r *http.Request)
type ClientRouteHandler = func(w http.ResponseWriter, r *http.Request, client *db.Client)

type Message struct {
	Content string `json:"content"`
	From    string `json:"from"`
	To      string `json:"to"`
	Epoch   int64  `json:"time"`
}

func HandleMessage(w http.ResponseWriter, r *http.Request, client *db.Client) {
	fmt.Println("Endpoint Hit: message")
	var p Message
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	ref, err := client.NewRef("messages").Push(firebase.Ctx, p)
	fmt.Println("Message: ", ref)
}
