default: build
	./bin/server

build:
	go build -o ./bin/server .

run:
	./bin/server
