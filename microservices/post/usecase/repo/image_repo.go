package repo

import (
	"bytes"
	"context"
)

type ImageRepo interface {
	BatchCreateImages(ctx context.Context, postID int64, imageBufs []*bytes.Buffer) error
	BatchDeleteImages(ctx context.Context, ids []int64) error
	DeleteImagesByPostID(ctx context.Context, postID int64) error
}
