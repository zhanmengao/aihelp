package v35

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/sashabaranov/go-openai"
	"github.com/zhanmengao/aihelp/config"
	"github.com/zhanmengao/aihelp/global"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"io/ioutil"
	"net/http"
)

type TGpt35 struct {
	UID        string
	SessionKey string
}

func (p *TGpt35) SendMessage(ctx context.Context, defaultMsg, userContent string) (retMsg string, err error) {
	//读出当前会话上下文
	dbMsg, _, err := global.PikaDB.GetDBUserMessage(ctx, p.UID, p.SessionKey)
	if err != nil {
		return
	}
	//cli := client.NewV3Client()
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: nil,
		//MaxTokens:        4096,
		//Temperature:      0.5,
		//TopP:             0,
		//N:                1,
		//Stream:           false,
		//Stop:             nil,
		//PresencePenalty:  0,
		//ResponseFormat:   nil,
		//Seed:             nil,
		//FrequencyPenalty: 0,
		//LogitBias:        nil,
		//LogProbs:         false,
		//TopLogProbs:      0,
		//User:             "",
		//Tools:            nil,
		//ToolChoice:       nil,
	}
	for _, msg := range dbMsg.Message {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:         msg.Role,
			Content:      msg.Content,
			MultiContent: nil,
			Name:         "",
			FunctionCall: nil,
			ToolCalls:    nil,
			ToolCallID:   "",
		})
	}
	//如果为空，设置默认msg
	if len(dbMsg.Message) <= 0 {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:         "system",
			Content:      defaultMsg,
			MultiContent: nil,
			Name:         "",
			FunctionCall: nil,
			ToolCalls:    nil,
			ToolCallID:   "",
		})
	}
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:         "user",
		Content:      userContent,
		MultiContent: nil,
		Name:         "",
		FunctionCall: nil,
		ToolCalls:    nil,
		ToolCallID:   "",
	})
	u := fmt.Sprintf("%s/%s", config.GptConfig.V3Host, "v1/chat/completions")
	body, _ := json.Marshal(req)
	httpReq, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.GptConfig.V3ApiKey))
	resp, err := g.Client().Do(httpReq)
	//resp, err := cli.CreateChatCompletion(ctx, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//解析
	rsp := &pb.GptMessageResponse{}
	if err = json.Unmarshal(body, rsp); err != nil {
		return
	}
	rspMsg := rsp.Choices[0].Message
	//回写
	dbMsg.Message = append(dbMsg.Message, rspMsg)
	if err = global.PikaDB.SetDBUserMessage(ctx, dbMsg); err != nil {
		return
	}
	glog.Debugf(ctx, "gpt rsp [%+v]", string(body))
	retMsg = rspMsg.Content
	return
}
