SHELL=/bin/bash

docker-build:
	docker build -t test .

build:
	go mod tidy
	go build -o test .

clean:
	rm test
