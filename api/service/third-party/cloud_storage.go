package service

import (
	"context"

	payload "github.com/sinulingga23/go-pos/payload/third-party"
)

type CloudStorageService interface {
	AddFile(ctx context.Context, cloudStorage payload.CloudStorage) (string, error)
}