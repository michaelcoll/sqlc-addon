build:
	go build -v -ldflags="-s -w -X 'github.com/michaelcoll/sqlc-addon/cmd.version=v0.0.0'" .

.PHONY: test
test:
	go test -v ./...

dep-upgrade:
	go get -u
	go mod tidy
