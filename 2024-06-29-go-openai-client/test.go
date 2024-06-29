package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// 获取 OpenAI API 密钥
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY 环境变量未设置")
		return
	}

	// 获取 OpenAI API 基础 URL
	baseUrl := os.Getenv("OPENAI_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://api.openai.com/v1" // 使用默认基础 URL
	}

	// 创建 OpenAI API 客户端
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl // 设置自定义基础 URL
	client := openai.NewClientWithConfig(config)

	// 使用客户端发送请求
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleUser, Content: "你好！"},},
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印响应
	fmt.Println(response.Choices[0].Message.Content)
}