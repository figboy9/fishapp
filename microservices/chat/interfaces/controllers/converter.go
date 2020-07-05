package controllers

import (
	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/pb"
	"github.com/golang/protobuf/ptypes"
)

func convRoomProto(r *domain.Room) (*pb.Room, error) {
	uAt, err := ptypes.TimestampProto(r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	cAt, err := ptypes.TimestampProto(r.CreatedAt)
	if err != nil {
		return nil, err
	}

	memP, err := convListMembersProto(r.Members)
	if err != nil {
		return nil, err
	}

	mP, err := convListMessagesProto(r.Messages)
	if err != nil {
		return nil, err
	}

	return &pb.Room{
		Id:        r.ID,
		PostId:    r.PostID,
		Members:   memP,
		Messages:  mP,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convMemberProto(m *domain.Member) (*pb.Member, error) {
	uAt, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	cAt, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.Member{
		Id:        m.ID,
		RoomId:    m.RoomID,
		UserId:    m.UserID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convListMembersProto(list []*domain.Member) ([]*pb.Member, error) {
	listM := make([]*pb.Member, len(list))
	for i, m := range list {
		mProto, err := convMemberProto(m)
		if err != nil {
			return nil, err
		}
		listM[i] = mProto
	}
	return listM, nil
}

func convMessageProto(m *domain.Message) (*pb.Message, error) {
	uAt, err := ptypes.TimestampProto(m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	cAt, err := ptypes.TimestampProto(m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.Message{
		Id:        m.ID,
		Body:      m.Body,
		RoomId:    m.RoomID,
		UserId:    m.UserID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convListMessagesProto(list []*domain.Message) ([]*pb.Message, error) {
	listM := make([]*pb.Message, len(list))
	for i, m := range list {
		mProto, err := convMessageProto(m)
		if err != nil {
			return nil, err
		}
		listM[i] = mProto
	}
	return listM, nil
}
