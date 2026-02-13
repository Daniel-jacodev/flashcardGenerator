# Estágio 1: Build do Go
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Copia os arquivos de dependências e baixa
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
RUN go build -o main ./cmd/server/main.go

# Estágio 2: Execução (Imagem ultra leve)
FROM debian:bookworm-slim

# Instala apenas o essencial para conexões HTTPS e certificados
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Cria a pasta de uploads para quando você enviar arquivos PDF/MP3 manualmente
RUN mkdir -p /uploads

# Copia o binário do estágio de build
COPY --from=builder /app/main .

# O Koyeb/Railway injeta as variáveis de ambiente automaticamente, 
# então não precisamos copiar o .env aqui.

EXPOSE 8080

CMD ["./main"]