package service

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	payload "github.com/sinulingga23/go-pos/payload/third-party"
	"google.golang.org/api/option"
)

func TestCloudStorageService_AddFile_Success(t *testing.T) {
	storageCredentialFile := os.Getenv("STORAGE_CREDENTIAL_FILE")
	client, err := storage.NewClient(context.TODO(), option.WithCredentialsFile(storageCredentialFile))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer client.Close()

	cloudStorageService := NewCloudStorageService(client)

	bucketName := os.Getenv("BUCKET_NAME")
	urlImage, err := cloudStorageService.AddFile(context.TODO(), payload.CloudStorage{
		BucketName: bucketName,
		FileName: fmt.Sprintf("%v-%v.txt", time.Now().Unix(), "test"),
		FileBytes: []byte("hello world!"),
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(urlImage) == 0 {
		t.Fatalf("got %q want %q", "empty", "not empty")
	}
}