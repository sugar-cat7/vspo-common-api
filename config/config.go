package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	"google.golang.org/api/option"
)

func addMockDataToFirestore(ctx context.Context, client *firestore.Client, collectionName string, dataSlice []interface{}) {
	for _, data := range dataSlice {
		dataValue := reflect.ValueOf(data)

		// Check if dataValue is a pointer, if not, then it's directly a struct.
		var docID string
		if dataValue.Kind() == reflect.Ptr {
			docID = dataValue.Elem().FieldByName("ID").String()
		} else {
			docID = dataValue.FieldByName("ID").String()
		}

		if docID == "" {
			log.Fatalf("Failed to get ID from the struct")
			return
		}
		_, err := client.Collection(collectionName).Doc(docID).Set(ctx, data)
		if err != nil {
			log.Fatalf("Failed adding data with ID %s to collection %s: %v", docID, collectionName, err)
		}
	}
}

func setupEmulator(ctx context.Context, client *firestore.Client) {
	// Mock video data
	videos := []interface{}{
		factories.NewVideo("videoID1"),
		factories.NewVideo("videoID2"),
		// ... additional mock videos
	}

	// Mock clip data
	clips := []interface{}{
		factories.NewClip("clipID1"),
		factories.NewClip("clipID2"),
		// ... additional mock clips
	}

	// Mock channel data
	channels := []interface{}{
		factories.NewChannel("channelID1"),
		factories.NewChannel("channelID2"),
		// ... additional mock channels
	}

	// Mock liveStream data
	liveStreams := []interface{}{
		factories.NewLiveStream("liveStreamID1"),
		factories.NewLiveStream("liveStreamID2"),
		// ... additional mock liveStreams
	}

	addMockDataToFirestore(ctx, client, "songs", videos)
	addMockDataToFirestore(ctx, client, "clips", clips)
	addMockDataToFirestore(ctx, client, "channels", channels)
	addMockDataToFirestore(ctx, client, "livestreams", liveStreams)
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
		if err != nil {
			log.Fatalf("error getting Firestore client: %v\n", err)
		}
	}

	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}

	return client, err
}
