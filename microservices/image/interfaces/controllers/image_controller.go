package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ezio1119/fishapp-image/conf"
	"github.com/ezio1119/fishapp-image/models"
	"github.com/ezio1119/fishapp-image/pb"
	"github.com/ezio1119/fishapp-image/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
)

type imageController struct {
	imageInteractor interactor.ImageInteractor
}

func NewImageController(i interactor.ImageInteractor) *imageController {
	return &imageController{i}
}

func (c *imageController) ListImagesByOwnerID(ctx context.Context, in *pb.ListImagesByOwnerIDReq) (*pb.ListImagesByOwnerIDRes, error) {
	ctx, cancel := context.WithTimeout(ctx, conf.C.Sv.TimeoutDuration)
	defer cancel()

	o, err := convOwnerType(in.OwnerType)
	if err != nil {
		return nil, err
	}

	list, err := c.imageInteractor.ListImagesByOwnerID(ctx, o, in.OwnerId)
	if err != nil {
		return nil, err
	}

	listP, err := convListImageProto(list)
	if err != nil {
		return nil, err
	}

	return &pb.ListImagesByOwnerIDRes{Images: listP}, nil
}

func (c *imageController) BatchCreateImages(stream pb.ImageService_BatchCreateImagesServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), conf.C.Sv.TimeoutDuration)
	defer cancel()

	images := []*models.Image{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch x := req.Data.(type) {
		case *pb.BatchCreateImagesReq_Info:
			o, err := convOwnerType(x.Info.OwnerType)
			if err != nil {
				return err
			}

			img := &models.Image{
				OwnerID:   x.Info.OwnerId,
				OwnerType: o,
				Buf:       &bytes.Buffer{},
			}
			images = append(images, img)

		case *pb.BatchCreateImagesReq_Chunk:

			lastImg := images[len(images)-1]

			if _, err := lastImg.Buf.Write(x.Chunk); err != nil {
				return err
			}

		default:
			return fmt.Errorf("BatchCreateImages.Data has unexpected type %T", x)
		}
	}

	if err := c.imageInteractor.BatchCreateImages(ctx, images); err != nil {
		return err
	}

	imgsP, err := convListImageProto(images)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.BatchCreateImagesRes{Images: imgsP})
}

func (c *imageController) BatchDeleteImages(ctx context.Context, in *pb.BatchDeleteImagesReq) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, conf.C.Sv.TimeoutDuration)
	defer cancel()

	if err := c.imageInteractor.BatchDeleteImages(ctx, in.Ids); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *imageController) BatchDeleteImagesByOwnerIDs(ctx context.Context, in *pb.BatchDeleteImagesByOwnerIDsReq) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, conf.C.Sv.TimeoutDuration)
	defer cancel()

	o, err := convOwnerType(in.OwnerType)
	if err != nil {
		return nil, err
	}

	if err := c.imageInteractor.BatchDeleteImagesByOwnerIDs(ctx, o, in.OwnerIds); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *imageController) DeleteImagesByOwnerID(ctx context.Context, in *pb.DeleteImagesByOwnerIDReq) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, conf.C.Sv.TimeoutDuration)
	defer cancel()

	o, err := convOwnerType(in.OwnerType)
	if err != nil {
		return nil, err
	}

	if err := c.imageInteractor.DeleteImagesByOwnerID(ctx, o, in.OwnerId); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
