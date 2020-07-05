package interactor

import (
	"time"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/golang/protobuf/ptypes"
)

func convPostProto(p *models.Post) (*pb.Post, error) {
	cAt, err := ptypes.TimestampProto(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	uAt, err := ptypes.TimestampProto(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	mAt, err := ptypes.TimestampProto(p.MeetingAt)
	if err != nil {
		return nil, err
	}

	return &pb.Post{
		Id:                p.ID,
		Title:             p.Title,
		Content:           p.Content,
		FishingSpotTypeId: p.FishingSpotTypeID,
		FishTypeIds:       models.ConvPostsFishTypeIDs(p.PostsFishTypes),
		PrefectureId:      p.PrefectureID,
		MeetingPlaceId:    p.MeetingPlaceID,
		MeetingAt:         mAt,
		MaxApply:          p.MaxApply,
		UserId:            p.UserID,
		CreatedAt:         cAt,
		UpdatedAt:         uAt,
	}, nil

}

func convListPostsProto(list []*models.Post) ([]*pb.Post, error) {
	listP := make([]*pb.Post, len(list))
	for i, p := range list {
		pProto, err := convPostProto(p)
		if err != nil {
			return nil, err
		}
		listP[i] = pProto
	}
	return listP, nil
}

func convApplyPostProto(a *models.ApplyPost) (*pb.ApplyPost, error) {
	cAt, err := ptypes.TimestampProto(a.CreatedAt)
	if err != nil {
		return nil, err
	}
	uAt, err := ptypes.TimestampProto(a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.ApplyPost{
		Id:        a.ID,
		PostId:    a.PostID,
		UserId:    a.UserID,
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}

func convListApplyPostsProto(list []*models.ApplyPost) ([]*pb.ApplyPost, error) {
	listA := make([]*pb.ApplyPost, len(list))
	for i, a := range list {
		aP, err := convApplyPostProto(a)
		if err != nil {
			return nil, err
		}
		listA[i] = aP
	}
	return listA, nil
}

func convPostFilter(f *pb.ListPostsReq_Filter) (*models.PostFilter, error) {
	postF := &models.PostFilter{CanApply: f.CanApply, FishTypeIDs: f.FishTypeIds}

	if f.MeetingAtFrom != nil {
		mAtFrom, err := ptypes.Timestamp(f.MeetingAtFrom)
		if err != nil {
			return nil, err
		}
		postF.MeetingAtFrom = mAtFrom.In(time.Local)
	}

	if f.MeetingAtTo != nil {
		mAtTo, err := ptypes.Timestamp(f.MeetingAtTo)
		if err != nil {
			return nil, err
		}
		postF.MeetingAtTo = mAtTo.In(time.Local)
	}

	switch f.OrderBy {
	case pb.ListPostsReq_Filter_ASC:
		postF.OrderBy = models.OrderByAsc
	case pb.ListPostsReq_Filter_DESC:
		postF.OrderBy = models.OrderByDesc
	}
	switch f.SortBy {
	case pb.ListPostsReq_Filter_CREATED_AT:
		postF.SortBy = models.SortByID
	case pb.ListPostsReq_Filter_MEETING_AT:
		postF.SortBy = models.SortByMeetingAt
	}
	return postF, nil
}
