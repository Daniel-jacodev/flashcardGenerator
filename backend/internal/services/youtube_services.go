package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetYouTubeTranscript(videoURL string) (string, error) {
	// Dados para o microsserviço Python
	requestBody, _ := json.Marshal(map[string]string{
		"url": videoURL,
	})

	// Chamada para o Python local
	resp, err := http.Post("http://localhost:5000/transcript", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("microsserviço Python está offline: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro no python: %s", result["error"])
	}

	return result["transcript"], nil
}