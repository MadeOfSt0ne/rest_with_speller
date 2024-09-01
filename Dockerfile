FROM golang:1.22.2
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /rest_speller cmd/main.go
CMD ["/rest_speller"]