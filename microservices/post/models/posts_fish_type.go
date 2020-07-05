package models

import "time"

type PostsFishType struct {
	ID         int64
	PostID     int64
	FishTypeID int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func ConvPostsFishTypeIDs(fTypes []*PostsFishType) []int64 {
	fIDs := make([]int64, len(fTypes))
	for i, fType := range fTypes {
		fIDs[i] = fType.FishTypeID
	}
	return fIDs
}

func ConvPostsFishTypes(fIDs []int64) []*PostsFishType {
	fTypes := make([]*PostsFishType, len(fIDs))
	for i, fID := range fIDs {
		fTypes[i] = &PostsFishType{FishTypeID: fID}
	}
	return fTypes
}
