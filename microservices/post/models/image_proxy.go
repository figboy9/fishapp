package models

import (
	"bytes"

	"github.com/ezio1119/fishapp-post/pb"
)

type BatchCreateImagesProxy struct {
	ImageInfo *pb.ImageInfo
	Buf       *bytes.Buffer
}
