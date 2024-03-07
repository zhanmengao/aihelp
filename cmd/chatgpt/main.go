package main

import (
	"context"
	"github.com/zhanmengao/aihelp/config"
	"github.com/zhanmengao/aihelp/global"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	endpoint = "wss://api.openai.com/v1/engines/davinci/completions"
)

type Request struct {
	Model       string  `json:"model"`
	MaxTokens   int     `json:"max_tokens"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
}

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int    `json:"index"`
		LogProbs     string `json:"logprobs"`
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

func main() {
	global.Flags()
	if err := config.Init(global.ConfigFile); err != nil {
		panic(err)
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5)*time.Second))
	defer cancel()
	header := http.Header{}
	header.Set("Authorization", "Bearer "+config.GptConfig.ApiKey)
	// 建立 WebSocket 连接
	c, _, err := websocket.DefaultDialer.DialContext(ctx, endpoint, header)
	if err != nil {
		log.Fatal("dial:", err)
	} else {
		slog.Info("connect to %s", endpoint)
	}
	defer c.Close()

	done := make(chan struct{})

	// 处理来自 ChatGPT 的消息
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				slog.Error("read:", err)
				return
			}
			slog.Info("Received:", string(message))
		}
	}()

	// 发送消息给 ChatGPT
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				requestData := Request{
					Model:       "davinci",
					MaxTokens:   50,
					Prompt:      "Once upon a time,",
					Temperature: 0.7,
				}

				if err := c.WriteJSON(requestData); err != nil {
					slog.Error("write:", err)
					return
				}
				slog.Info("Sent:", t.String())
			}
		}
	}()
	select {}
}
