build:
	go build -v -ldflags="-s -w" .

.PHONY: test
test:
	go test -v ./...

.PHONY: vet
vet: ## check go code
	@go vet ./...

dep-upgrade:
	go get -u
	go mod tidy
