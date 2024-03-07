package v35

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/zhanmengao/aihelp/module/gpt/client"
)

type TGpt35 struct {
	UID        string
	SessionKey string
}

func (p *TGpt35) SendMessage(ctx context.Context, msg []byte) (err error) {
	//TODO 读出当前会话上下文
	cli := client.NewV3Client()
	resp, err := cli.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:            "",
		Messages:         nil,
		MaxTokens:        4096,
		Temperature:      0,
		TopP:             0,
		N:                0,
		Stream:           false,
		Stop:             nil,
		PresencePenalty:  0,
		ResponseFormat:   nil,
		Seed:             nil,
		FrequencyPenalty: 0,
		LogitBias:        nil,
		LogProbs:         false,
		TopLogProbs:      0,
		User:             "",
		Tools:            nil,
		ToolChoice:       nil,
	})
	if err != nil {
		return
	}
	return
}
