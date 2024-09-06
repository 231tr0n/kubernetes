SHELL=/bin/bash

.PHONY: docker-build
docker-build:
	$(MAKE) clean
	$(MAKE) build
	docker rmi -f trial
	docker build -t trial .

.PHONY: docker-run
docker-run:
	$(MAKE) docker-build
	docker run --rm -p 8080:8080 trial

.PHONY: build
build:
	$(MAKE) clean
	go mod tidy
	GOEXPERIMENT=boringcrypto go build -v -o trial .

.PHONY: rate-limit-test
rate-limit-test:
	go run utils/rateLimit.go http://localhost:8080/test 1000 100

.PHONY: run
run:
	go mod tidy
	PORT=:8080 GOEXPERIMENT=boringcrypto go run .

.PHONY: clean
clean:
	rm -rf ./trial
