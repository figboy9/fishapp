package controllers

import (
	"errors"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/pb"
	"github.com/golang/protobuf/ptypes"
)

func convUserProto(u *domain.User) (*pb.User, error) {
	updatedAt, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	pbSex, err := convSexProto(u.Sex)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:           u.ID,
		Email:        u.Email,
		Name:         u.Name,
		Introduction: u.Introduction,
		Sex:          pbSex,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}, nil
}

func convSexProto(s domain.Sex) (pb.Sex, error) {
	switch s {
	case domain.Male:
		return pb.Sex_MALE, nil
	case domain.Female:
		return pb.Sex_FEMALE, nil
	default:
		return 0, errors.New("unexpected domain.Sex type")
	}
}

func convSex(s pb.Sex) (domain.Sex, error) {
	switch s {
	case pb.Sex_MALE:
		return domain.Male, nil
	case pb.Sex_FEMALE:
		return domain.Female, nil
	default:
		return 0, errors.New("unexpected pb.Sex type")
	}
}

func convTokenPairProto(tp *domain.TokenPair) *pb.TokenPair {
	return &pb.TokenPair{
		IdToken:      tp.IDToken,
		RefreshToken: tp.RefreshToken,
	}
}
