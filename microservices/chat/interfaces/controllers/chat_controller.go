package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type chatController struct {
	chatInteractor interactor.ChatInteractor
}

func NewChatController(ci interactor.ChatInteractor) pb.ChatServiceServer {
	return &chatController{ci}
}

func (c *chatController) GetRoom(ctx context.Context, in *pb.GetRoomReq) (*pb.Room, error) {
	var id int64
	var pID int64

	switch x := in.GetRoom.(type) {
	case *pb.GetRoomReq_RoomId:
		id = x.RoomId
	case *pb.GetRoomReq_PostId:
		pID = x.PostId
	}

	r, err := c.chatInteractor.GetRoom(ctx, id, pID)
	if err != nil {
		return nil, err
	}
	return convRoomProto(r)
}

func (c *chatController) CreateRoom(ctx context.Context, in *pb.CreateRoomReq) (*pb.Room, error) {
	r := &domain.Room{
		PostID: in.PostId,
		Members: []*domain.Member{
			{UserID: in.UserId},
		},
	}
	if err := c.chatInteractor.CreateRoom(ctx, r); err != nil {
		return nil, err
	}
	return convRoomProto(r)
}

func (c *chatController) IsMember(ctx context.Context, in *pb.IsMemberReq) (*wrappers.BoolValue, error) {
	var rID int64
	var pID int64

	switch x := in.IsMember.(type) {
	case *pb.IsMemberReq_PostId:
		pID = x.PostId
	case *pb.IsMemberReq_RoomId:
		rID = x.RoomId
	}

	isMember, err := c.chatInteractor.IsMember(ctx, rID, pID, in.UserId)
	if err != nil {
		return nil, err
	}

	return &wrapperspb.BoolValue{Value: isMember}, nil
}

func (c *chatController) ListMembers(ctx context.Context, in *pb.ListMembersReq) (*pb.ListMembersRes, error) {
	list, err := c.chatInteractor.ListMembers(ctx, in.RoomId)
	if err != nil {
		return nil, err
	}
	listMProto, err := convListMembersProto(list)
	if err != nil {
		return nil, err
	}
	return &pb.ListMembersRes{Members: listMProto}, nil
}

func (c *chatController) CreateMember(ctx context.Context, in *pb.CreateMemberReq) (*pb.Member, error) {
	m := &domain.Member{RoomID: in.RoomId, UserID: in.UserId}
	if err := c.chatInteractor.CreateMember(ctx, m); err != nil {
		return nil, err
	}
	return convMemberProto(m)
}

func (c *chatController) DeleteMember(ctx context.Context, in *pb.DeleteMemberReq) (*empty.Empty, error) {
	if err := c.chatInteractor.DeleteMember(ctx, in.RoomId, in.UserId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *chatController) ListMessages(ctx context.Context, in *pb.ListMessagesReq) (*pb.ListMessagesRes, error) {
	list, err := c.chatInteractor.ListMessages(ctx, in.RoomId)
	if err != nil {
		return nil, err
	}
	listMProto, err := convListMessagesProto(list)
	if err != nil {
		return nil, err
	}
	return &pb.ListMessagesRes{Messages: listMProto}, nil
}

func (c *chatController) CreateMessage(stream pb.ChatService_CreateMessageServer) error {

	ctx := stream.Context()

	m := &domain.Message{}
	imageBuf := &bytes.Buffer{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch x := req.Data.(type) {
		case *pb.CreateMessageReq_Info:
			fmt.Printf("CreateMessageReq_Info: %#v\n", x.Info)
			m.Body = x.Info.Body
			m.RoomID = x.Info.RoomId
			m.UserID = x.Info.UserId

		case *pb.CreateMessageReq_ImageChunk:
			fmt.Printf("CreateMessageReq_ImageChunk: %#v\n", x.ImageChunk)
			if m.Body != "" {
				status.Error(codes.InvalidArgument, "invalid CreateMessageReqInfo.Body, CreateMessageReqImageChunk.ImageChunk: value must be set either body or image")
			}

			if _, err := imageBuf.Write(x.ImageChunk); err != nil {
				return err
			}
		}
	}

	if err := c.chatInteractor.CreateMessage(ctx, m, imageBuf); err != nil {
		return err
	}

	pbM, err := convMessageProto(m)
	if err != nil {
		return err
	}

	return stream.SendAndClose(pbM)
}

func (c *chatController) StreamMessage(in *pb.StreamMessageReq, stream pb.ChatService_StreamMessageServer) error {
	eg, ctx := errgroup.WithContext(stream.Context())
	msgChan := make(chan *domain.Message)
	go func() {
		eg.Wait()
		close(msgChan)
	}()

	eg.Go(func() error {
		if err := c.chatInteractor.StreamMessage(ctx, in.RoomId, msgChan); err != nil {
			return err
		}
		return nil
	})

	for m := range msgChan {
		mProto, err := convMessageProto(m)
		if err != nil {
			return err
		}
		if err := stream.Send(mProto); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
