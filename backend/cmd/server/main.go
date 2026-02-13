package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Daniel-jacodev/flashcard-generator/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Função auxiliar para limpar as tags XML das legendas do YouTube
func cleanXMLTags(input string) string {
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllString(input, " ")
}

func main() {
	// Tenta carregar o .env apenas localmente
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado (comum em produção)")
	}

	r := gin.Default()

	// Configuração do CORS para permitir que a Vercel acesse o Koyeb
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	r.POST("/generate", func(c *gin.Context) {
		var textoExtraido string
		var serviceErr error
		var filePath string

		youtubeURL := c.PostForm("url")

		if youtubeURL != "" {
			log.Println("Processando link do YouTube via Extração de Texto:", youtubeURL)
			
			// Nova função que pega legendas em vez de baixar áudio
			textoExtraido, serviceErr = services.GetYouTubeTranscript(youtubeURL)
			
			// Limpa as tags <text> do XML do YouTube para economizar tokens na Groq
			if serviceErr == nil {
				textoExtraido = cleanXMLTags(textoExtraido)
			}

		} else {
			// Lógica para Upload de Arquivo (PDF ou Áudio)
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Envie um arquivo ou URL do YouTube"})
				return
			}

			// Define o caminho para salvar temporariamente
			filePath = filepath.Join("../../uploads/", file.Filename)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo no servidor"})
				return
			}

			// Garante que o arquivo será deletado após o processamento
			defer os.Remove(filePath)

			ext := strings.ToLower(filepath.Ext(file.Filename))
			if ext == ".pdf" {
				textoExtraido, serviceErr = services.ReadPdf(filePath)
			} else {
				// Para arquivos de áudio, ainda usamos o Whisper
				textoExtraido, serviceErr = services.TranscribeAudio(filePath)
			}
		}

		// Verifica se houve erro em qualquer um dos serviços de extração
		if serviceErr != nil {
			log.Println("ERRO NO PROCESSAMENTO:", serviceErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao extrair conteúdo: " + serviceErr.Error()})
			return
		}

		// Envia o texto final para a IA gerar os flashcards
		flashcards, err := services.GenerateFlashcards(textoExtraido)
		if err != nil {
			log.Println("ERRO NA IA:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar flashcards com a IA"})
			return
		}

		// Retorno de sucesso
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"flashcards": flashcards,
		})
	})

	// Rota de check básico para ver se o servidor está vivo
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Porta configurada via variável de ambiente (Koyeb usa PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	r.Run(":" + port)
}