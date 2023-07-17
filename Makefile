build:
	go build -o bin/main main.go

run:
	go run main.go

vendor: 
	go mod tidy
	go mod vendor

test: vendor
	cd task; go test .

all: vendor run