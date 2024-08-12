SHELL=/bin/bash

docker-build:
	$(MAKE) clean
	$(MAKE) build
	docker rmi -f trial
	docker build -t trial .

build:
	$(MAKE) clean
	go mod tidy
	GOEXPERIMENT=boringcrypto go build -v -o trial .

run:
	go mod tidy
	GOEXPERIMENT=boringcrypto go run .

clean:
	rm -rf ./trial
