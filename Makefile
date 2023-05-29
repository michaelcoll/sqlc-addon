build:
	go build -v -ldflags="-s -w" .

.PHONY: test
test:
	go test -vet=all ./...

dep-upgrade:
	go get -u
	go mod tidy
