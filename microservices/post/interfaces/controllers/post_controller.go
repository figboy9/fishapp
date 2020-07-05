package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/ezio1119/fishapp-post/usecase/interactor"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postController struct {
	postInteractor interactor.PostInteractor
}

func NewPostController(pu interactor.PostInteractor) *postController {
	return &postController{pu}
}

func (c *postController) GetPost(ctx context.Context, in *pb.GetPostReq) (*pb.Post, error) {
	p, err := c.postInteractor.GetPost(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return convPostProto(p)
}

func (c *postController) ListPosts(ctx context.Context, in *pb.ListPostsReq) (*pb.ListPostsRes, error) {
	f, err := convPostFilter(in.Filter)
	if err != nil {
		return nil, err
	}
	list, nextToken, err := c.postInteractor.ListPosts(ctx, &models.Post{
		FishingSpotTypeID: in.Filter.FishingSpotTypeId,
		PrefectureID:      in.Filter.PrefectureId,
		UserID:            in.Filter.UserId,
	}, in.PageSize, in.PageToken, f)

	if err != nil {
		return nil, err
	}

	listProto, err := convListPostsProto(list)
	if err != nil {
		return nil, err
	}

	return &pb.ListPostsRes{Posts: listProto, NextPageToken: nextToken}, nil
}

func (c *postController) CreatePost(stream pb.PostService_CreatePostServer) error {
	ctx := stream.Context()
	p := &models.Post{}

	imageBufs := []*bytes.Buffer{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch x := req.Data.(type) {
		case *pb.CreatePostReq_Info:

			mAt, err := ptypes.Timestamp(x.Info.MeetingAt)
			if err != nil {
				return err
			}

			p.Title = x.Info.Title
			p.Content = x.Info.Content
			p.FishingSpotTypeID = x.Info.FishingSpotTypeId
			p.PostsFishTypes = models.ConvPostsFishTypes(x.Info.FishTypeIds)
			p.PrefectureID = x.Info.PrefectureId
			p.MeetingPlaceID = x.Info.MeetingPlaceId
			p.MeetingAt = mAt
			p.MaxApply = x.Info.MaxApply
			p.UserID = x.Info.UserId

		case *pb.CreatePostReq_ImageChunk:

			if len(imageBufs) == 0 {
				imageBufs = append(imageBufs, &bytes.Buffer{})
			}

			lastBuf := imageBufs[len(imageBufs)-1]

			if _, err := lastBuf.Write(x.ImageChunk); err != nil {
				return err
			}

		case *pb.CreatePostReq_NextImageSignal:

			imageBufs = append(imageBufs, &bytes.Buffer{})

		default:
			return fmt.Errorf("CreatePostReq.Request has unexpected type %T", x)
		}
	}

	sagaID, err := c.postInteractor.CreatePost(ctx, p, imageBufs)
	if err != nil {
		return err
	}

	pProto, err := convPostProto(p)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.CreatePostRes{
		Post:   pProto,
		SagaId: sagaID,
	})
}

func (c *postController) UpdatePost(stream pb.PostService_UpdatePostServer) error {
	ctx := stream.Context()
	p := &models.Post{}

	imageBufs := []*bytes.Buffer{}
	dltImageIDs := []int64{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch x := req.Data.(type) {
		case *pb.UpdatePostReq_Info:

			mAt, err := ptypes.Timestamp(x.Info.MeetingAt)
			if err != nil {
				return err
			}

			p.ID = x.Info.Id
			p.Title = x.Info.Title
			p.Content = x.Info.Content
			p.FishingSpotTypeID = x.Info.FishingSpotTypeId
			p.PostsFishTypes = models.ConvPostsFishTypes(x.Info.FishTypeIds)
			p.PrefectureID = x.Info.PrefectureId
			p.MeetingPlaceID = x.Info.MeetingPlaceId
			p.MeetingAt = mAt
			p.MaxApply = x.Info.MaxApply
			dltImageIDs = x.Info.ImageIdsToDelete

		case *pb.UpdatePostReq_ImageChunk:

			if len(imageBufs) == 0 {
				imageBufs = append(imageBufs, &bytes.Buffer{})
			}

			lastBuf := imageBufs[len(imageBufs)-1]

			if _, err := lastBuf.Write(x.ImageChunk); err != nil {
				return err
			}

		case *pb.UpdatePostReq_NextImageSignal:

			imageBufs = append(imageBufs, &bytes.Buffer{})

		default:
			return fmt.Errorf("CreatePostReq.Request has unexpected type %T", x)
		}
	}

	fmt.Println(len(imageBufs))

	if err := c.postInteractor.UpdatePost(ctx, p, imageBufs, dltImageIDs); err != nil {
		return err
	}

	pProto, err := convPostProto(p)
	if err != nil {
		return err
	}

	return stream.SendAndClose(pProto)
}

func (c *postController) DeletePost(ctx context.Context, in *pb.DeletePostReq) (*empty.Empty, error) {
	if err := c.postInteractor.DeletePost(ctx, in.Id); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *postController) GetApplyPost(ctx context.Context, in *pb.GetApplyPostReq) (*pb.ApplyPost, error) {
	a, err := c.postInteractor.GetApplyPost(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return convApplyPostProto(a)
}

func (c *postController) ListApplyPosts(ctx context.Context, in *pb.ListApplyPostsReq) (*pb.ListApplyPostsRes, error) {
	if (in.Filter.UserId == 0 && in.Filter.PostId == 0) || (in.Filter.UserId != 0 && in.Filter.PostId != 0) {
		return nil, status.Error(codes.InvalidArgument, "invalid ListApplyPostsReq.Filter.PostId, ListApplyPostsReq.Filter.UserId: value must be set either user_id or post_id")
	}
	list, err := c.postInteractor.ListApplyPosts(ctx, &models.ApplyPost{
		UserID: in.Filter.UserId,
		PostID: in.Filter.PostId,
	})
	if err != nil {
		return nil, err
	}
	listProto, err := convListApplyPostsProto(list)
	if err != nil {
		return nil, err
	}
	return &pb.ListApplyPostsRes{ApplyPosts: listProto}, nil
}

func (c *postController) BatchGetApplyPostsByPostIDs(ctx context.Context, in *pb.BatchGetApplyPostsByPostIDsReq) (*pb.BatchGetApplyPostsByPostIDsRes, error) {
	list, err := c.postInteractor.BatchGetApplyPostsByPostIDs(ctx, in.PostIds)
	if err != nil {
		return nil, err
	}
	listProto, err := convListApplyPostsProto(list)
	if err != nil {
		return nil, err
	}
	return &pb.BatchGetApplyPostsByPostIDsRes{ApplyPosts: listProto}, nil
}

func (c *postController) CreateApplyPost(ctx context.Context, in *pb.CreateApplyPostReq) (*pb.ApplyPost, error) {
	a := &models.ApplyPost{
		PostID: in.PostId,
		UserID: in.UserId,
	}
	err := c.postInteractor.CreateApplyPost(ctx, a)
	if err != nil {
		return nil, err
	}
	return convApplyPostProto(a)
}

func (c *postController) DeleteApplyPost(ctx context.Context, in *pb.DeleteApplyPostReq) (*empty.Empty, error) {
	if err := c.postInteractor.DeleteApplyPost(ctx, in.Id); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
