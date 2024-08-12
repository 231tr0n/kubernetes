SHELL=/bin/bash

docker-build:
	docker build -t test .

build:
	go mod tidy
	GOEXPERIMENT=boringcrypto go build -o test .

clean:
	rm test
