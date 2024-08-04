FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY ./backend . 

RUN go mod download
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

CMD ["go", "run", "main.go"]