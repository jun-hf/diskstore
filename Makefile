build:
	go build -o bin/diskstore
run: build
	./bin/diskstore