package gpt

import (
	"context"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	v35 "github.com/zhanmengao/aihelp/module/gpt/internal/v35"
)

// IGpt 一个gpt封装
type IGpt interface {
	SendMessage(ctx context.Context, defaultMsg, userContent string) (retMsg []*pb.GptMessage, err error)
}

func NewGpt(ctx context.Context, uid string, sessionID string) IGpt {
	return &v35.TGpt35{}
}
