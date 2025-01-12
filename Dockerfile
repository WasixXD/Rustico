# Etapa 1: Build do backend (Go)
FROM golang:1.23.3-alpine AS backend

RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY ./go /app
COPY go.mod /app
COPY go.sum /app

RUN go mod tidy
RUN CGO_ENABLED=1 go build -o backend .

# Etapa 2: Imagem final
FROM alpine:latest

RUN apk add --no-cache ca-certificates nodejs npm

COPY --from=backend ./app/backend /usr/local/bin/backend

RUN apk add --no-cache sqlite
COPY ./.output /app/.output
COPY ./sql/ /app
RUN sqlite3 -init /app/init.sql /app/project.db .quit
COPY ./go/project.db /usr/local/bin/project.db
EXPOSE 3001 

# Comando para rodar o backend e o frontend
CMD ["sh", "-c", "/usr/local/bin/backend & node /app/.output/server/index.mjs"]
