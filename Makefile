.PHONY: build

build:
	@go mod tidy
	@mkdir -p build
	@go build -o build/hedge_cli

clean:
	@rm -fr build
