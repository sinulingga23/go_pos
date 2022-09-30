package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sinulingga23/go-pos/definition"
	payload "github.com/sinulingga23/go-pos/payload/third-party"
)

type cloudStorageService struct {
	client *storage.Client
}

func NewCloudStorageService(client *storage.Client) *cloudStorageService {
	return &cloudStorageService{client: client}
}

func (c cloudStorageService) AddFile(ctx context.Context, cloudStorage payload.CloudStorage) (string, error) {
	if len(strings.Trim(cloudStorage.BucketName, " ")) == 0 || len(strings.Trim(cloudStorage.FileName, " ")) == 0 ||
		len(cloudStorage.FileBytes) == 0 {
		return "", definition.ErrBadRequest
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

	wc := c.client.Bucket(cloudStorage.BucketName).Object(cloudStorage.FileName).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0
	
	buff := bytes.NewBuffer(cloudStorage.FileBytes)
	if _, err := io.Copy(wc, buff); err != nil {
		log.Printf("[CloudStorageService][io.Copy]: %v\n", err.Error())
		return "", definition.ErrInternalServer
	}

	// Data can continue to be added to the file until the writer is closed
	if err := wc.Close(); err != nil {
		log.Printf("[CloudStorageService][wc.Close]: %v\n", err.Error())
		return "", definition.ErrInternalServer
	}

	return fmt.Sprintf("%s/%s/%s", "https://storage.googleapis.com", cloudStorage.BucketName, cloudStorage.FileName), nil
}