### for cross compile
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build