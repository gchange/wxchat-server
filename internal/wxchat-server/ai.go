package wxchatserver

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type AI struct {
	client *openai.Client
	config *AIConfig
}

func NewAI(config *AIConfig) (*AI, error) {
	client := openai.NewClient(
		option.WithBaseURL(config.BaseUrl),
		option.WithAPIKey(config.Key),
	)
	return &AI{
		client: client,
		config: config,
	}, nil
}

func (ai *AI) ChatCompletions(ctx context.Context, message string) (content string, err error) {
	chatCompletion, err := ai.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		}),
		Model: openai.F(openai.ChatModel(ai.config.Model)),
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}

func (s *Server) GetAIClient(model string) (c *AI, err error) {
	v, ok := s.ai.Load(model)
	if !ok {
		return nil, fmt.Errorf("AI model %s not found.", model)
	}
	ai, ok := v.(*AI)
	if !ok {
		return nil, fmt.Errorf("AI model %s is not *AI.", model)
	}
	return ai, nil
}
