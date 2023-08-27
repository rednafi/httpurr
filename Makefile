.PHONY: lint
lint:
	@echo "Running lint"
	@golangci-lint run --fix
	@go mod tidy


.PHONY: lint-check
lint-check:
	@echo "Checking lint"
	@golangci-lint run


.PHONY: test
test:
	@echo "Running tests"
	@go test -v ./...


.PHONY: clean
clean:
	@echo "Cleaning up"
	@go clean
	@rm -rf ./bin


.PHONY: init
init:
	@echo "Initializing project"
	@go mod download
	@go mod tidy
	@go test -v ./...
