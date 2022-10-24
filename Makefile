.PHONY: build

build:
	@go mod tidy
	@go mod vendor
	@mkdir -p build
	@go build -o build/hedge_cli

clean:
	@rm -fr build
