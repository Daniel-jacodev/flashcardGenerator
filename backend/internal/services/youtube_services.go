package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kkdai/youtube/v2"
)

func GetYouTubeTranscript(videoURL string) (string, error) {
	client := youtube.Client{}
	
	// 1. Pega informações do vídeo
	video, err := client.GetVideo(videoURL)
	if err != nil {
		return "", fmt.Errorf("erro ao obter vídeo: %v", err)
	}

	// 2. Tenta encontrar legendas (Captions)
	if len(video.CaptionTracks) > 0 {
		// Pega a primeira legenda disponível (geralmente a automática ou em inglês/português)
		track := video.CaptionTracks[0]
		
		resp, err := http.Get(track.BaseURL)
		if err != nil {
			return "", fmt.Errorf("erro ao baixar legenda: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("erro ao ler corpo da legenda: %v", err)
		}

		// O retorno é um XML. Para a Groq, podemos mandar o XML bruto 
		// ou limpar as tags. A Groq é esperta o suficiente para ler o XML.
		return string(body), nil
	}

	// 3. Fallback: Se não tiver legenda, usamos a descrição e o título
	// Isso garante que a IA tenha algum contexto para trabalhar
	contexto := fmt.Sprintf("Título: %s\n\nDescrição: %s", video.Title, video.Description)
	return contexto, nil
}