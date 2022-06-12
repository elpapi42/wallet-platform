build:
	go build -o bin/application source/infrastructure/main.go

start:
	bin/application

dev:
	air

make test:
	go test -v ./tests/
