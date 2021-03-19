package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	_ "firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"fmt"
	"google.golang.org/api/option"
)

var Ctx = context.Background()

func CreateFirebaseApp(keyPath string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(keyPath)
	config := &firebase.Config{
		ProjectID:   "onionchat-7094e",
		DatabaseURL: "https://onionchat-7094e-default-rtdb.firebaseio.com/",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

//
func CreateFirestoreClient(app *firebase.App) (*db.Client, error) {
	client, err := app.Database(Ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing db client: %v", err)
	}
	return client, nil
}
