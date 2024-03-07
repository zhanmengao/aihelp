package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zhanmengao/aihelp/config"
	"github.com/zhanmengao/aihelp/global"
	"io/ioutil"
	"log/slog"
	"net/http"
	"time"
)

type OpenAIRequest struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
	Model       string  `json:"model"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// https_proxy=http://127.0.0.1:64578 http_proxy=http://127.0.0.1:64578 all_proxy=socks5://127.0.0.1:64579
func main() {
	global.Flags()
	if err := config.Init(global.ConfigFile); err != nil {
		panic(err)
	}
	gptURL := "https://api.openai.com/v1/completions"
	// 输入提示
	prompt := "Once upon a time,"

	// 设置请求结构体
	requestData := OpenAIRequest{
		Prompt:      prompt,
		MaxTokens:   50,
		Temperature: 0.7,
		//Model:       "davinci-002",
		Model: "gpt-3.5-turbo",
	}

	// 将请求数据转换为 JSON 格式
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding request JSON:", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", gptURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// 设置 HTTP 请求标头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GptConfig.ApiKey)

	// 发送 HTTP 请求
	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 解析响应体
	var responseData OpenAIResponse
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("Error decoding response JSON:", err)
		return
	}
	slog.Info("ret %s", string(responseBody))

	// 输出生成的文本
	for _, choice := range responseData.Choices {
		fmt.Println(choice.Text)
	}
}
