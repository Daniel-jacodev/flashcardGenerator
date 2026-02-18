# 📇 Gerador de Flashcards com IA (YouTube & PDF)

Este projeto é uma ferramenta completa para estudantes que desejam transformar conteúdos de vídeos do YouTube ou documentos PDF em flashcards de estudo de forma automática. O sistema utiliza **Inteligência Artificial (Groq/Gemini)** para processar o texto e gerar perguntas e respostas inteligentes.

A arquitetura foi redesenhada para utilizar **microsserviços**, garantindo que a extração de transcrições do YouTube funcione sem bloqueios de IP, rodando de forma resiliente no ambiente local.

---

## 🏗️ Arquitetura do Sistema

O projeto é dividido em três partes que se comunicam via APIs REST:

1.  **Frontend (React + Vite):** Interface de usuário moderna para inserção de links e arquivos.
2.  **Backend Orquestrador (Go):** Gerencia as rotas, processa PDFs e comunica-se com a IA.
3.  **Microsserviço de Transcrição (Python):** Especializado em extrair legendas do YouTube de forma segura.

---

## 🚀 Como Rodar o Projeto Localmente

### Pré-requisitos

- **Go** (versão 1.20+)
- **Python** (versão 3.10+)
- **Node.js & npm**
- **Chave de API da Groq** (ou Gemini)

---

### 1️⃣ Microsserviço de Extração (Python)

Este serviço evita os bloqueios de segurança do YouTube ao rodar no seu IP residencial.

```bash
# Navegue até a pasta do microsserviço
cd yt-transcript-service

# Crie e ative o ambiente virtual (venv)
python3 -m venv venv
source venv/bin/activate

# Instale as dependências
pip install flask youtube-transcript-api flask-cors

# Inicie o serviço
python3 app.py
Porta padrão: http://localhost:5000

2️⃣ Backend Orquestrador (Go)
O "cérebro" que coordena a leitura de dados e a geração por IA.

Bash
# Navegue até a pasta do backend
cd backend

# Configure suas variáveis de ambiente no arquivo .env
# Exemplo: GROQ_API_KEY=gsk_your_key_here

# Execute o servidor
go run cmd/server/main.go
Porta padrão: http://localhost:8080

3️⃣ Frontend (React)
A interface visual do usuário.

Bash
# Navegue até a pasta do frontend
cd frontend

# Instale as dependências
npm install

# Inicie o painel de desenvolvimento
npm run dev
Porta padrão: http://localhost:5173

🛠️ Funcionalidades
YouTube to Flashcards: Transcreve automaticamente vídeos através da URL.

PDF to Flashcards: Faz o upload de arquivos PDF e extrai o conteúdo textual.

IA Generativa: Utiliza modelos de linguagem avançados para criar pares de Pergunta/Resposta.

CORS Enabled: Configurado para permitir a comunicação fluida entre os diferentes serviços locais.

📂 Estrutura de Pastas
Plaintext
.
├── backend/
│   ├── cmd/server/main.go          # Servidor principal em Go
│   ├── internal/services/          # Lógica de PDF, IA e YouTube
│   └── uploads/                    # Armazenamento temporário de arquivos
├── frontend/
│   ├── src/App.jsx                 # Interface principal em React
│   └── tailwind.config.js          # Configurações de estilo
└── yt-transcript-service/          # Microsserviço Python (Flask)
```
