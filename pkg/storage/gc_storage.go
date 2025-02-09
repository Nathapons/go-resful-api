package storage

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

type CloudStorage interface {
	UploadFile() error
	DownloadFile() error
	DeleteFile() error
}

type cloudStorage struct {
	client *storage.Client
}

func NewGcpStrorage() CloudStorage {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create GCP client: %v", err)
	}

	return &cloudStorage{client: client}
}

func (gcps *cloudStorage) UploadFile() error {
	return nil
}

func (gcps *cloudStorage) DownloadFile() error {
	return nil
}

func (gcps *cloudStorage) DeleteFile() error {
	return nil
}
