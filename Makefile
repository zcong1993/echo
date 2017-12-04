generate:
	@go generate ./...

build: generate
	@echo "====> Build echo"
	@sh -c ./build.sh
