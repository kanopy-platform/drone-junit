CMD_NAME=drone-junit
RUN ?= .*
PKG ?= ./...

.PHONY:
tidy:
	go mod tidy
	go mod verify

.PHONY: test
test: tidy ## Run tests in local environment
	go test -short -cover -run=$(RUN) $(PKG)

.PHONY: build
build:
	mkdir -p build
	go build -o build/$(CMD_NAME) ./cmd/$(CMD_NAME)/
