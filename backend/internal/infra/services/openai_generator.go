package services

import (
	"context"
	"encoding/json"
	"errors"

	openai "github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type OpenAIGenerator struct {
	client *openai.Client
	model  string
}

func NewOpenAIGenerator(apiKey string, model string) *OpenAIGenerator {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return &OpenAIGenerator{
		client: &client,
		model:  model,
	}
}

func (g *OpenAIGenerator) GenerateSummary(ctx context.Context, rawPostContent string) (string, error) {
	if rawPostContent == "" {
		return "", errors.New("raw post content cannot be empty")
	}

	prompt := "Please provide a concise summary of the following post:\n\n" + rawPostContent

	params := openai.ChatCompletionNewParams{
		Model: g.model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You summarize content."),
			openai.UserMessage(prompt),
		},
	}

	resp, err := g.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

func (g *OpenAIGenerator) GenerateTags(ctx context.Context, rawPostContent string) ([]string, error) {
	if rawPostContent == "" {
		return nil, errors.New("raw post content cannot be empty")
	}

	prompt := "Given the following post content, extract a list of 3 to 6 relevant tags (single words):\n\n" +
		rawPostContent +
		"\n\nProvide the tags as a JSON array of strings, for example: [\"go\", \"database\", \"web\"] without quote marks"

	params := openai.ChatCompletionNewParams{
		Model: g.model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You extract tags from content."),
			openai.UserMessage(prompt),
		},
	}

	resp, err := g.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("no response from OpenAI")
	}

	content := resp.Choices[0].Message.Content

	var tags []string
	if err := json.Unmarshal([]byte(content), &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
