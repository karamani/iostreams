build:
	gofmt -w .
	go tool vet *.go
	go test
	go build