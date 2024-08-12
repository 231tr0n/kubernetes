SHELL=/bin/bash

docker-build:
	$(MAKE) clean
	$(MAKE) build
	docker rmi -f trial
	docker build -t trial .

docker-run:
	$(MAKE) docker-build
	docker run --rm -p 8080:8080 trial

build:
	$(MAKE) clean
	go mod tidy
	GOEXPERIMENT=boringcrypto go build -v -o trial .

run:
	go mod tidy
	PORT=:8080 GOEXPERIMENT=boringcrypto go run .

clean:
	rm -rf ./trial
