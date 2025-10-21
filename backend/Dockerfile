# Estágio de construção (BUILDER)
# Usamos a versão 1.23, que já corrigimos ser a correta
FROM golang:1.23 AS builder 

# 1. Definir o diretório de trabalho conforme a convenção Go
WORKDIR /go/src/app

# 2. Copiar os arquivos de módulo (para cache)
COPY go.mod .
COPY go.sum .

# 3. Baixar dependências
RUN go mod download

# 4. Copiar todo o código-fonte restante
COPY . .

# 5. Construir o executável
# Usamos o caminho completo do módulo e CGO_ENABLED=0 para um executável estático
# que não terá problemas no estágio FROM scratch
RUN CGO_ENABLED=0 go build -o main -ldflags='-w -s' .


# Estágio Final (PRODUÇÃO)
# Imagem scratch (a menor possível)
FROM scratch 

# Copiar os certificados de Root (necessário para a API Go acessar o DB)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Definir o diretório e copiar o binário
WORKDIR /app
COPY --from=builder /go/src/app/main . 

# Expor a porta e definir o comando de execução
EXPOSE 8080
CMD ["./main"]