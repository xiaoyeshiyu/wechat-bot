package gpt

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

func (g *GPT) CreateChatCompletion(ctx context.Context, content string) (string, error) {
	resp, err := g.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return resp.Choices[0].Message.Content, nil
}
