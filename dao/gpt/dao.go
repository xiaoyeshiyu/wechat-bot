package gpt

import (
	"wechat-bot/config"

	"github.com/sashabaranov/go-openai"
)

type GPT struct {
	client *openai.Client
}

func NewGPT() *GPT {
	return &GPT{
		client: openai.NewClient(config.GetApiKey()),
	}
}
