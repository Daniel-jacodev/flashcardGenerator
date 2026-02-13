# Estágio 1: Build do Go
FROM golang:1.21-bookworm AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário (ajuste o caminho se o seu main estiver em outro lugar)
RUN go build -o main ./cmd/server/main.go

# Estágio 2: Execução (Imagem leve)
FROM debian:bookworm-slim

# Instala yt-dlp e ffmpeg
RUN apt-get update && apt-get install -y \
    yt-dlp \
    ffmpeg \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Cria a pasta de uploads (importante!)
RUN mkdir -p uploads

# Copia o binário do estágio de build
COPY --from=builder /app/main .

# Copia o .env se necessário (ou use as variáveis do Render)
# COPY .env .

EXPOSE 8080

CMD ["./main"]