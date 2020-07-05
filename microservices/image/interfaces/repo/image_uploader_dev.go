package repo

import (
	"context"
	"io"
	"os"

	"github.com/ezio1119/fishapp-image/conf"
	"github.com/ezio1119/fishapp-image/usecase/repo"
)

type imageUploaderDevRepo struct{}

func NewImageUploaderDevRepo() repo.ImageUploaderRepo {
	return &imageUploaderDevRepo{}
}

func (r *imageUploaderDevRepo) UploadImage(ctx context.Context, image io.Reader, objName string) error {
	f, err := os.Create(conf.C.Sv.DevImagesPath + objName)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, image)
	return nil
}

func (r *imageUploaderDevRepo) DeleteUploadedImage(ctx context.Context, objName string) error {
	return os.Remove(conf.C.Sv.DevImagesPath + objName)
}
