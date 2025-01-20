package storage

import "cloud.google.com/go/storage"

type CloudStorage interface {
	UploadFile() error
	DownloadFile() error
	DeleteFile() error
}

type cloudStorage struct {
	client *storage.Client
}
