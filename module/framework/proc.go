package framework

import (
	"context"
	"strconv"
)

func GetUID(ctx context.Context) string {
	uid, ok := ctx.Value(KeyNameUID).(string)
	if ok {
		return uid
	}
	//TODO grpc
	return ""
}

func NewID() string {
	ret, _ := uidGen.GenID()
	return strconv.Itoa(int(ret))
}
