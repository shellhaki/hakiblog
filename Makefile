run:
	go run ./cmd

build:
	cd cmd && go build -o ../bin/app

start:
	./bin/app

tidy:
	go mod tidy