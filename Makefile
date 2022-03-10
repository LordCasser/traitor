

.PHONY: build
build:
	CGO_ENABLED=0 go build ./cmd/traitor

.PHONY: pack
prepare:
	sudo apt-get update && sudo apt-get install -y i686-linux-gnu-gcc aarch64-linux-gnu-gcc
	go run ./prepare.go

.PHONY: install
install:
	CGO_ENABLED=0 go install -ldflags "-X traitor/version.Version=`git describe --tags`" ./cmd/traitor

.PHONY: test
test:
	go test ./... -race -cover
