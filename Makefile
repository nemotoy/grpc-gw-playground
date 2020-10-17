GOCMD=go
GOTEST=$(GOCMD) test
GOLINT=golangci-lint

all: lint test

test:
	$(GOTEST) -v ./...

test.race:
	$(GOTEST) -v -race ./...

lint:
	$(GOLINT) run -v ./... -c .golangci.yaml

pb.go:
	rm proto/user/user.pb.go
	protoc -I . \
		--go_out . \
		--go_opt plugins=grpc \
		--go_opt paths=source_relative proto/user/user.proto
