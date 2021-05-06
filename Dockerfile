FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build cmd/main.go
WORKDIR /dist
RUN cp /app/main .
EXPOSE 8000

CMD ["/dist/main"]