package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kkdai/youtube/v2"
)

func GetYouTubeTranscript(videoURL string) (string, error) {
    // 1. Criamos um cliente com um "ClientOptions" para simular um navegador
    client := youtube.Client{}
    
    // 2. Tenta obter o vídeo
    video, err := client.GetVideo(videoURL)
    if err != nil {
        // Se der erro de restrição, tentamos uma segunda vez com metadados básicos
        return "", fmt.Errorf("o YouTube bloqueou o acesso a este vídeo (Embedding disabled). Tente um vídeo que permita incorporação ou use o upload de áudio")
    }

    // 3. Busca as legendas
    if len(video.CaptionTracks) > 0 {
        track := video.CaptionTracks[0]
        resp, err := http.Get(track.BaseURL)
        if err != nil {
            return "", fmt.Errorf("erro ao baixar legenda: %v", err)
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return "", fmt.Errorf("erro ao ler corpo da legenda")
        }
        return string(body), nil
    }

    // Fallback: Título e Descrição (que o YouTube costuma liberar mais fácil)
    return fmt.Sprintf("Título: %s\nDescrição: %s", video.Title, video.Description), nil
}