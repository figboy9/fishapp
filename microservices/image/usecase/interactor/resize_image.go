package interactor

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/disintegration/imaging"
	"github.com/ezio1119/fishapp-image/conf"
	"github.com/ezio1119/fishapp-image/models"
)

func resizeImage(i *models.Image) error {
	img, t, err := image.Decode(i.Buf)
	if err != nil {
		return err
	}

	nrgba := imaging.Fit(img, conf.C.Sv.ImageWidth, conf.C.Sv.ImageHeight, imaging.Lanczos)

	switch t {
	case "jpeg":
		if err := jpeg.Encode(i.Buf, nrgba, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			return err
		}
		i.Name = i.Name + ".jpg"
	case "png":
		if err := png.Encode(i.Buf, nrgba); err != nil {
			return err
		}
		i.Name = i.Name + ".png"
	case "gif":
		if err := gif.Encode(i.Buf, nrgba, nil); err != nil {
			return err
		}
		i.Name = i.Name + ".gif"
	}

	return nil
}
