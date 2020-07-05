package interactor

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
	return &pb.Room{
		Id:        r.ID,
		PostId:    r.PostID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}
