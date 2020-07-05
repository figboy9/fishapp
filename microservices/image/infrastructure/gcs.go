package infrastructure

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/ezio1119/fishapp-image/conf"
)

func NewGCSClient(ctx context.Context) (*storage.Client, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	if _, err = c.Bucket(conf.C.Gcs.BucketName).Attrs(ctx); err != nil {
		return nil, err
	}

	return c, nil
}
