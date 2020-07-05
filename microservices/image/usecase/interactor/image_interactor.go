package interactor

import (
	"context"

	"github.com/ezio1119/fishapp-image/models"
	"github.com/ezio1119/fishapp-image/usecase/repo"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type imageInteractor struct {
	db                *gorm.DB
	imageUploaderRepo repo.ImageUploaderRepo
}

func NewImageInteractor(db *gorm.DB, ur repo.ImageUploaderRepo) *imageInteractor {
	return &imageInteractor{db, ur}
}

type ImageInteractor interface {
	ListImagesByOwnerID(ctx context.Context, ownerType models.OwnerType, ownerID int64) ([]*models.Image, error)
	BatchCreateImages(ctx context.Context, images []*models.Image) error
	BatchDeleteImages(ctx context.Context, id []int64) error
	BatchDeleteImagesByOwnerIDs(ctx context.Context, ownerType models.OwnerType, ownerIDs []int64) error
	DeleteImagesByOwnerID(ctx context.Context, ownerType models.OwnerType, ownerID int64) error
}

func (i *imageInteractor) ListImagesByOwnerID(ctx context.Context, ownerType models.OwnerType, ownerID int64) ([]*models.Image, error) {
	images := []*models.Image{}
	if err := i.db.Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (i *imageInteractor) BatchCreateImages(ctx context.Context, images []*models.Image) error {
	for _, img := range images {
		img.Name = uuid.New().String()

		if err := resizeImage(img); err != nil {
			return err
		}

		if err := i.db.Create(img).Error; err != nil {
			return err
		}

		if err := i.imageUploaderRepo.UploadImage(ctx, img.Buf, img.Name); err != nil {
			return err
		}
	}

	return nil

}

func (i *imageInteractor) BatchDeleteImages(ctx context.Context, ids []int64) error {
	images := []*models.Image{}

	if err := i.db.Where("id IN (?)", ids).Find(images).Error; err != nil {
		return nil
	}

	for _, img := range images {
		if err := i.imageUploaderRepo.DeleteUploadedImage(ctx, img.Name); err != nil {
			return err
		}
	}

	return i.db.Where("id IN (?)", ids).Delete(&models.Image{}).Error
}

func (i *imageInteractor) BatchDeleteImagesByOwnerIDs(ctx context.Context, ownerType models.OwnerType, ownerIDs []int64) error {
	return i.db.Where("owner_id = IN(?) AND owner_type = ?", ownerIDs, ownerType).Delete(&models.Image{}).Error
}

func (i *imageInteractor) DeleteImagesByOwnerID(ctx context.Context, ownerType models.OwnerType, ownerID int64) error {
	images := []*models.Image{}

	if err := i.db.Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).Find(&images).Error; err != nil {
		return nil
	}

	for _, img := range images {
		if err := i.imageUploaderRepo.DeleteUploadedImage(ctx, img.Name); err != nil {
			return err
		}
	}

	return i.db.Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).Delete(&models.Image{}).Error
}
