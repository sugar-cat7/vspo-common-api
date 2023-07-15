package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// Firestore is the client to Firestore.
// var Firestore *firestore.Client

func Config() (client *firestore.Client, err error) {
	var Firestore *firestore.Client
	ctx := context.Background()

	conf := &firebase.Config{
		ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
	}

	// Generate a properly escaped JSON string for the private key.
	privateKey, _ := json.Marshal(strings.Replace(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n", -1))

	opt := option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "` + os.Getenv("FIREBASE_PROJECT_ID") + `",
		"private_key_id": "` + os.Getenv("FIREBASE_PRIVATE_KEY_ID") + `",
		"private_key": ` + string(privateKey) + `,
		"client_email": "` + os.Getenv("FIREBASE_CLIENT_EMAIL") + `",
		"client_id": "` + os.Getenv("FIREBASE_CLIENT_ID") + `",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "` + os.Getenv("FIREBASE_CLIENT_CERT_URL") + `"
	}`))

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	Firestore, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	return Firestore, err
}
