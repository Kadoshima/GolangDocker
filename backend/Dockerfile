FROM golang:1.19.1-alpine3.16

WORKDIR /app

# air のインストール (hotreload)
RUN go install github.com/cosmtrek/air@latest

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

# アプリケーションのコピー
COPY ./backend .

# ホットリロードのために air を実行
CMD ["air"]

# CMD ["go", "run", "./cmd/api"]