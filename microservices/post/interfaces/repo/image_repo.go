package repo

import (
	"bytes"
	"context"
	"io"

	"github.com/ezio1119/fishapp-post/conf"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/ezio1119/fishapp-post/usecase/repo"
)

type imageRepo struct {
	client pb.ImageServiceClient
}

func NewImageRepo(c pb.ImageServiceClient) repo.ImageRepo {
	return &imageRepo{c}
}

func (r *imageRepo) BatchCreateImages(ctx context.Context, pID int64, imageBufs []*bytes.Buffer) error {
	stream, err := r.client.BatchCreateImages(ctx)
	if err != nil {
		return err
	}

	for _, imageBuf := range imageBufs {
		req := &pb.BatchCreateImagesReq{
			Data: &pb.BatchCreateImagesReq_Info{
				Info: &pb.ImageInfo{
					OwnerId:   pID,
					OwnerType: pb.OwnerType_POST,
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

func (r *imageRepo) BatchDeleteImages(ctx context.Context, ids []int64) error {
	if _, err := r.client.BatchDeleteImages(ctx, &pb.BatchDeleteImagesReq{Ids: ids}); err != nil {
		return err
	}

	return nil
}

func (r *imageRepo) DeleteImagesByPostID(ctx context.Context, pID int64) error {
	if _, err := r.client.DeleteImagesByOwnerID(ctx, &pb.DeleteImagesByOwnerIDReq{
		OwnerId:   pID,
		OwnerType: pb.OwnerType_POST,
	}); err != nil {
		return err
	}

	return nil
}
