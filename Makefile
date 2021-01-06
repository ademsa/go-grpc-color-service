default: build

clean:
	rm -r ./bin && mkdir ./bin

prepare-dependencies:
	go mod download

build-code:
	go build -o ./bin/go-grpc-color-service ./cmd/main.go

build: clean prepare-dependencies build-code

run-code:
	./bin/go-grpc-color-service

build-and-run: build run-code