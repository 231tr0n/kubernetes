SHELL=/bin/bash

docker-build:
	docker rmi -f test
	docker build -t test .

build:
	go mod tidy
	GOEXPERIMENT=boringcrypto go build -v -o test .

clean:
	rm test
