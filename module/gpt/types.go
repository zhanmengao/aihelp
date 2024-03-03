package gpt

import "context"

// IGpt 一个gpt封装
type IGpt interface {
	SendMessage(ctx context.Context, msg []byte) (err error)
}

func NewGpt(ctx context.Context, uid string, sessionID string) IGpt {
	return nil
}
