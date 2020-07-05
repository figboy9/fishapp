package repo

import (
	"context"
	"io"
)

type ImageUploaderRepo interface {
	UploadImage(ctx context.Context, image io.Reader, objName string) error
	DeleteUploadedImage(ctx context.Context, objName string) error
}
