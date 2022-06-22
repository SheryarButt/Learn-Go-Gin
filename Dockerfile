FROM golang:latest

LABEL maintainer="Sheryar <m.sheryarbutt@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
ENV PORT 5000
RUN go build -o app

RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./app"]