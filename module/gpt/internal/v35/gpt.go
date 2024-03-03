package v35

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/zhanmengao/aihelp/config"
)

type TGpt35 struct {
	UID        string
	SessionKey string
}

func (p *TGpt35) SendMessage(ctx context.Context, msg []byte) (err error) {
	//TODO 读出当前会话上下文
	cli := openai.NewClient(config.GptConfig.ApiKey)
	_, err = cli.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:            "",
		Messages:         nil,
		MaxTokens:        0,
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
		Functions:        nil,
		FunctionCall:     nil,
		Tools:            nil,
		ToolChoice:       nil,
	})
	return
}
