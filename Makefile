.PHONY: build run clean

build:
	go build -o repo .

run: build
	./repo

clean:
	rm -f repo
	go clean
