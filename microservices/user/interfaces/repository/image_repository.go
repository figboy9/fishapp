package repository

import (
	"bytes"
	"context"
	"io"

	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/pb"
)

type imageRepository struct {
	client pb.ImageServiceClient
}

func NewImageRepository(c pb.ImageServiceClient) *imageRepository {
	return &imageRepository{c}
}

func (r *imageRepository) BatchCreateImages(ctx context.Context, uID int64, imageBufs []*bytes.Buffer) error {
	stream, err := r.client.BatchCreateImages(ctx)
	if err != nil {
		return err
	}

	for _, imageBuf := range imageBufs {
		req := &pb.BatchCreateImagesReq{
			Data: &pb.BatchCreateImagesReq_Info{
				Info: &pb.ImageInfo{
					OwnerId:   uID,
					OwnerType: pb.OwnerType_USER,
				},
			}}

		if err := stream.Send(req); err != nil {
			return err
		}

		for {
			buf := make([]byte, conf.C.Sv.ImageChunkSize)
			n, err := imageBuf.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			req = &pb.BatchCreateImagesReq{
				Data: &pb.BatchCreateImagesReq_Chunk{
					Chunk: buf[:n],
				}}

			if err = stream.Send(req); err != nil {
				return err
			}
		}

	}

	if _, err := stream.CloseAndRecv(); err != nil {
		return err
	}

	return nil
}

func (r *imageRepository) BatchDeleteImages(ctx context.Context, ids []int64) error {
	if _, err := r.client.BatchDeleteImages(ctx, &pb.BatchDeleteImagesReq{Ids: ids}); err != nil {
		return err
	}

	return nil
}

func (r *imageRepository) DeleteImagesByUserID(ctx context.Context, uID int64) error {
	if _, err := r.client.DeleteImagesByOwnerID(ctx, &pb.DeleteImagesByOwnerIDReq{
		OwnerId:   uID,
		OwnerType: pb.OwnerType_USER,
	}); err != nil {
		return err
	}

	return nil
}
