package repo

import (
	"bytes"
	"context"
)

type ImageRepo interface {
	BatchCreateImages(ctx context.Context, messageID int64, imageBufs []*bytes.Buffer) error
	BatchDeleteImagesByMessageIDs(ctx context.Context, ids []int64) error
	DeleteImagesByMessageID(ctx context.Context, messageID int64) error
}
