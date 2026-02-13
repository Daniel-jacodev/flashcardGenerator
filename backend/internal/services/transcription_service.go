package services

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func TranscribeAudio(filePath string) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.groq.com/openai/v1"
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateTranscription(
		context.Background(),
		openai.AudioRequest{
			Model:    "whisper-large-v3",
			FilePath: filePath,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Text, nil
}