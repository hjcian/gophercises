FROM golang:latest

WORKDIR /app

COPY ./ /app
RUN go mod download
# version 1
# ENTRYPOINT go run commands/runserver.go

# version 2
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build commands/runserver.go" --command=./runserver