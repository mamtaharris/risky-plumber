test:
	go test ./... -coverprofile=c.out -covermode=count && go tool cover -func=c.out && go tool cover -html=c.out

server:
	go run main.go server

lint:
	golangci-lint run

mock:
	go generate ./...
