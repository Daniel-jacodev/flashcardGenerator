package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kkdai/youtube/v2"
)

func GetYouTubeTranscript(videoURL string) (string, error) {
	// Usamos um cliente padrão. 
	// Dica: Em servidores, o YouTube bloqueia o GetVideo quase sempre.
	client := youtube.Client{}
	
	video, err := client.GetVideo(videoURL)
	if err != nil {
		// Se o YouTube bloqueou o IP ou o Embedding, não damos erro 500.
		// Retornamos um erro específico que a main vai tratar.
		return "", fmt.Errorf("O YouTube bloqueou o acesso do servidor a este link (IP do servidor marcado).")
	}

	// Se ele conseguiu acessar o vídeo, tentamos as legendas
	if len(video.CaptionTracks) > 0 {
		track := video.CaptionTracks[0]
		resp, err := http.Get(track.BaseURL)
		if err != nil {
			return "", fmt.Errorf("Erro ao baixar legenda: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("Erro ao ler corpo da legenda")
		}
		return string(body), nil
	}

	// Se não tem legenda, mas conseguiu o vídeo, usa o que tem
	return fmt.Sprintf("Título: %s. Descrição: %s", video.Title, video.Description), nil
}