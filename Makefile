.PHONY: all vendor release

all: build

vendor:
	@go mod tidy
	@go mod vendor
	@go mod download

build: build-report

build-report:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ~/bin/shortcut-report cmd/report/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ~/bin/shortcut-report-by-team cmd/report-by-team/main.go

fmt:
	go fmt ./...
	go vet ./...

run:
	go run cmd/report/main.go

play:
	go run cmd/playground/main.go

