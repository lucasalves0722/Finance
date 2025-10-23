# ----------------------------------------------------------------------
# Estágio 1: Build (Compilação da Aplicação Go)
# ----------------------------------------------------------------------
# Usamos o Go 1.23, conforme exigido pelos erros de dependência
FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia TODOS os arquivos de código-fonte (incluindo go.mod, handlers, database, models)
COPY . .

# Baixa as dependências. O go.mod e go.sum estão presentes no COPY . . acima.
RUN go mod download

# Compila a aplicação. Usamos -mod=readonly para evitar conflito com o diretório 'vendor' local.
RUN CGO_ENABLED=0 GOOS=linux go build -a -mod=readonly -installsuffix cgo -o app ./main.go


# ----------------------------------------------------------------------
# Estágio 2: Final (Imagem de Produção Leve)
# ----------------------------------------------------------------------
FROM alpine:latest

# Define o fuso horário (Opcional, mas boa prática)
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Copia o binário 'app' compilado e o executa
WORKDIR /root/
COPY --from=builder /app/app .

# Comando de execução: Roda o binário compilado
CMD ["./app"]
