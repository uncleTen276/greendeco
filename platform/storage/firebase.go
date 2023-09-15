package storage

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"github.com/sekke276/greendeco.git/pkg/configs"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type firebaseApp struct {
	*firebase.App
}

var defaultApp = &firebaseApp{}

// Connect to firebase return app
func (s *firebaseApp) connectFirebaseStorage() error {
	cfg := configs.AppConfig()
	ctx := context.Background()

	ptr, err := json.Marshal(cfg.Storage.Firebase)
	if err != nil {
		return err
	}

	scope := []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/datastore",
		"https://www.googleapis.com/auth/devstorage.full_control",
		"https://www.googleapis.com/auth/firebase",
		"https://www.googleapis.com/auth/identitytoolkit",
		"https://www.googleapis.com/auth/userinfo.email",
	}

	credential, err := google.CredentialsFromJSON(ctx, ptr, scope...)
	if err != nil {
		return err
	}

	bucketConfig := &firebase.Config{
		StorageBucket: cfg.Storage.Bucket,
	}
	opt := option.WithCredentials(credential)
	s.App, err = firebase.NewApp(ctx, bucketConfig, opt)

	if err != nil {
		return err
	}

	return nil
}

func GetStorage(ctx context.Context) (*storage.Client, error) {
	storage, err := defaultApp.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func ConnectStorage() error {
	return defaultApp.connectFirebaseStorage()
}
