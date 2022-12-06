## mocks: generate mocks
mocks:
	@echo Generating mocks
	@go install github.com/vektra/mockery/v2
	@go generate ./...
