BUILD_NAME = nice-test
GO=$(shell which go)


.PHONY: build
build: 
	$(GO) build -o cmd/$(BUILD_NAME) cmd/main.go

.PHONY: run
run: build
	cd cmd && ./nice-test all
