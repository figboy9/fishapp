package controllers

import (
	"errors"

	"github.com/ezio1119/fishapp-image/models"
	"github.com/ezio1119/fishapp-image/pb"
	"github.com/golang/protobuf/ptypes"
)

func convImageProto(i *models.Image) (*pb.Image, error) {
	cAt, err := ptypes.TimestampProto(i.CreatedAt)
	if err != nil {
		return nil, err
	}
	uAt, err := ptypes.TimestampProto(i.UpdatedAt)
	if err != nil {
		return nil, err
	}
	ownerP, err := convOwnerTypeProto(i.OwnerType)
	if err != nil {
		return nil, err
	}
	return &pb.Image{
		Id:        i.ID,
		Name:      i.Name,
		OwnerId:   i.OwnerID,
		OwnerType: ownerP,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convListImageProto(list []*models.Image) ([]*pb.Image, error) {
	listI := make([]*pb.Image, len(list))
	for i, image := range list {
		iProto, err := convImageProto(image)
		if err != nil {
			return nil, err
		}
		listI[i] = iProto
	}
	return listI, nil
}

func convOwnerTypeProto(o models.OwnerType) (pb.OwnerType, error) {
	switch o {
	case models.POST:
		return pb.OwnerType_POST, nil
	case models.USER:
		return pb.OwnerType_USER, nil
	case models.MESSAGE:
		return pb.OwnerType_MESSAGE, nil
	default:
		return pb.OwnerType_OWNERTYPE_UNSPECIFIED, errors.New("convproto unexpected owner_type")
	}
}

func convOwnerType(o pb.OwnerType) (models.OwnerType, error) {
	switch o {
	case pb.OwnerType_POST:
		return models.POST, nil
	case pb.OwnerType_USER:
		return models.USER, nil
	case pb.OwnerType_MESSAGE:
		return models.MESSAGE, nil
	default:
		return 0, errors.New("convmodels unexpected owner_type")
	}
}
