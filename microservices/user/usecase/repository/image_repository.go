package repository

import (
	"bytes"
	"context"
)

type ImageRepository interface {
	BatchCreateImages(ctx context.Context, userID int64, imageBufs []*bytes.Buffer) error
	BatchDeleteImages(ctx context.Context, ids []int64) error
	DeleteImagesByUserID(ctx context.Context, userID int64) error
}
