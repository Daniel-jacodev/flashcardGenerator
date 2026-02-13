package services

import (
	"fmt"
	"os"
	"os/exec"
)

func DownloadYoutubeAudio(url string) (string, error) {
	finalPath := "../../uploads/yt_audio.mp3"


	os.Remove(finalPath)

	cmd := exec.Command("yt-dlp", 
    "--extractor-args", "youtube:player_client=ios,web", // Simula clientes mais comuns
    "--force-overwrites",
    "--no-playlist",
    "-x", 
    "--audio-format", "mp3",
    "-o", "../../uploads/yt_audio.%(ext)s", 
    url,
)


	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("falha no yt-dlp: %v | log: %s", err, string(output))
	}

	if _, err := os.Stat(finalPath); os.IsNotExist(err) {
		return "", fmt.Errorf("arquivo mp3 não foi gerado pelo yt-dlp")
	}

	return finalPath, nil
}