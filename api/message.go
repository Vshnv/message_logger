package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
	From    string `json:"from"`
	To      string `json:"to"`
	Epoch   int64  `json:"time"`
}

func HandleMessage(w http.ResponseWriter, r *http.Request, databasePath string) {
	fmt.Println("Endpoint Hit: message")

	var p Message
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	go insertMessage(databasePath, &p)
	fmt.Println("Message added")
}

func insertMessage(databasePath string, m *Message) {
	database, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		fmt.Println("Error connecting to database!")
		return
	}
	defer closeDatabase(database)
	statement, err := database.Prepare("INSERT INTO messages (from_user, to_user, content, epoch) VALUES (?,?,?,?)")
	if err != nil {
		fmt.Print("Error preparing message insertion statement.", err)
		return
	}
	_, err = statement.Exec(m.From, m.To, m.Content, m.Epoch)
	if err != nil {
		fmt.Print("Error during message.", err)
		return
	}
}

func closeDatabase(database *sql.DB) {
	if database == nil {
		return
	}
	_ = database.Close()
}
