APP_NAME=shtlk-fiber
MAIN_PKG=./cmd/api

run:
	go run $(MAIN_PKG)/main.go

build:
	go build -o bin/$(APP_NAME) $(MAIN_PKG)/main.go

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...
