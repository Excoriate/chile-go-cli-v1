fmt:
	gofmt -s -w .
test:
	go test ./... -cover
lint:
	golangci-lint run
build:
	go build -v -o chilecli main.go
