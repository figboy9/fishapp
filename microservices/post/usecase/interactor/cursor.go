package interactor

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
)

func extractIDFromPageToken(t string) (int64, error) {
	byteToken, err := base64.StdEncoding.DecodeString(t)
	if err != nil {
		return 0, err
	}
	splitToken := strings.Split(string(byteToken), ":")
	if splitToken[0] != "post" && len(splitToken) != 2 {
		return 0, errors.New("wrong page_token format")
	}
	id, err := strconv.ParseInt(splitToken[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func genPageTokenFromID(i int64) string {
	strID := strconv.FormatInt(i, 10)
	return base64.StdEncoding.EncodeToString([]byte("post:" + strID))
}
