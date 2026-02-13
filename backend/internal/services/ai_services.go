package services

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GenerateFlashcards(texto string) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.groq.com/openai/v1"
	client := openai.NewClientWithConfig(config)

	prompt := fmt.Sprintf("Atue como especialista em Anki e ENEM; com base no texto fornecido, gere um arquivo CSV (Frente;Verso) usando ponto e vírgula como separador, criando 10 flashcards de alto nível, sendo 3 fáceis, 4 médios e 3 difíceis , que foquem em interpretação, relações de causa/efeito e aplicações práticas do conteúdo, seguindo o Princípio da Atomicidade, fornecendo apenas o texto bruto, sem blocos de código, sem cabeçalhos, sem aspas e sem explicações. Texto: %s", texto)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "llama-3.3-70b-versatile",
			Messages: []openai.ChatCompletionMessage{{Role: "user", Content: prompt}},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

