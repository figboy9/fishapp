package repo

import (
	"bytes"
	"context"
	"io"

	"github.com/ezio1119/fishapp-chat/conf"
	"github.com/ezio1119/fishapp-chat/pb"
)

type imageRepo struct {
	client pb.ImageServiceClient
}

func NewImageRepository(c pb.ImageServiceClient) *imageRepo {
	return &imageRepo{c}
}

func (r *imageRepo) BatchCreateImages(ctx context.Context, mID int64, imageBufs []*bytes.Buffer) error {
	stream, err := r.client.BatchCreateImages(ctx)
	if err != nil {
		return err
	}

	for _, imageBuf := range imageBufs {
		req := &pb.BatchCreateImagesReq{
			Data: &pb.BatchCreateImagesReq_Info{
				Info: &pb.ImageInfo{
					OwnerId:   mID,
					OwnerType: pb.OwnerType_MESSAGE,
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

func (r *imageRepo) BatchDeleteImagesByMessageIDs(ctx context.Context, ids []int64) error {
	_, err := r.client.BatchDeleteImagesByOwnerIDs(ctx, &pb.BatchDeleteImagesByOwnerIDsReq{
		OwnerIds:  ids,
		OwnerType: pb.OwnerType_MESSAGE},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *imageRepo) DeleteImagesByMessageID(ctx context.Context, mID int64) error {
	if _, err := r.client.DeleteImagesByOwnerID(ctx, &pb.DeleteImagesByOwnerIDReq{
		OwnerId:   mID,
		OwnerType: pb.OwnerType_MESSAGE,
	}); err != nil {
		return err
	}

	return nil
}
