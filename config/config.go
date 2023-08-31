package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	"google.golang.org/api/option"
)

func setupEmulator(ctx context.Context, client *firestore.Client) {
	video := factories.NewVideo("videoID1")
	_, err := client.Collection("songs").Doc(video.ID).Set(ctx, video)
	if err != nil {
		log.Fatalf("Failed adding video: %v", err)
	}

	clip := factories.NewClip("clipID1")
	_, err = client.Collection("clips").Doc(clip.ID).Set(ctx, clip)
	if err != nil {
		log.Fatalf("Failed adding clip: %v", err)
	}

	channel := factories.NewChannel("channelID1")
	_, err = client.Collection("channels").Doc(channel.ID).Set(ctx, channel)
	if err != nil {
		log.Fatalf("Failed adding channel: %v", err)
	}
}

func Config() (client *firestore.Client, err error) {
	ctx := context.Background()

	// エミュレータへの接続情報が環境変数に設定されているかチェック
	emulatorHost := os.Getenv("FIRESTORE_EMULATOR_HOST")

	if emulatorHost != "" {
		client, err = firestore.NewClient(ctx, "test-project")
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		setupEmulator(ctx, client)
	} else {
		// 本番環境の設定（以前のコードをそのまま使用）
		conf := &firebase.Config{
			ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
		}

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

		client, err = app.Firestore(ctx)
	}

	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}

	return client, err
}
