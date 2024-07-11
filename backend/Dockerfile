FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

COPY ./backend .

CMD ["go", "run", "./cmd/api"]