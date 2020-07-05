package models

import (
	"time"
)

type Post struct {
	ID                int64
	Title             string
	Content           string
	FishingSpotTypeID int64
	PostsFishTypes    []*PostsFishType
	PrefectureID      int64
	MeetingPlaceID    string
	MeetingAt         time.Time
	MaxApply          int64
	UserID            int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type orderBy int64

const (
	OrderByAsc orderBy = iota
	OrderByDesc
)

func (o orderBy) String() string {
	switch o {
	case OrderByAsc:
		return "asc"
	case OrderByDesc:
		return "desc"
	}
	return ""
}

type sortBy int64

const (
	SortByID sortBy = iota
	SortByMeetingAt
)

func (s sortBy) String() string {
	switch s {
	case SortByID:
		return "id"
	case SortByMeetingAt:
		return "meeting_at"
	}
	return ""
}

type PostFilter struct {
	MeetingAtFrom time.Time
	MeetingAtTo   time.Time
	OrderBy       orderBy
	SortBy        sortBy
	FishTypeIDs   []int64
	CanApply      bool
}
