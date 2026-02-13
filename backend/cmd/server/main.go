package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Daniel-jacodev/flashcard-generator/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado")
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/generate", func(c *gin.Context) {
		var textoExtraido string
		var serviceErr error
		var filePath string

		
		youtubeURL := c.PostForm("url")

		if youtubeURL != "" {
			log.Println("Processando link do YouTube:", youtubeURL)
			filePath, serviceErr = services.DownloadYoutubeAudio(youtubeURL)
			if serviceErr == nil {

				defer os.Remove(filePath)
				textoExtraido, serviceErr = services.TranscribeAudio(filePath)
			}
		} else {
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Envie um arquivo ou URL"})
				return
			}

			filePath = "../../uploads/" + file.Filename
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo"})
				return
			}

			defer os.Remove(filePath)

			ext := strings.ToLower(filepath.Ext(file.Filename))
			if ext == ".pdf" {
				textoExtraido, serviceErr = services.ReadPdf(filePath)
			} else {
				textoExtraido, serviceErr = services.TranscribeAudio(filePath)
			}
		}

		if serviceErr != nil {
    		log.Println("ERRO NO WHISPER:", serviceErr)
    		c.JSON(500, gin.H{"error": serviceErr.Error()})
   			return
		}
		flashcards, err := services.GenerateFlashcards(textoExtraido)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar flashcards"})
			return
		}


		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"flashcards": flashcards,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}