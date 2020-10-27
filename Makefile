GOCMD=go
GOTEST=$(GOCMD) test
GOLINT=golangci-lint

all: pb.lint lint test

test:
	$(GOTEST) -v ./...

test.race:
	$(GOTEST) -v -race ./...

lint:
	$(GOLINT) run -v ./... -c .golangci.yaml

# pb.go:
# 	protoc -I . \
# 	--go_out . --go_opt paths=source_relative \
# 	--go-grpc_out . --go-grpc_opt paths=source_relative \
# 	proto/user/user.proto

pb.go:
	protoc -I . \
	--go_out . \
	--go_opt plugins=grpc \
	--go_opt paths=source_relative proto/user/user.proto

pb.clean:
	rm proto/user/user.pb.go

pb.lint:
	buf check lint

evans.req:
	echo '{ "id" : 1 }' | evans --proto proto/user/user.proto cli call UserService.GetUser

# For request to server in local by the evans.
# init later, request to the server below commands.
# $ header access_key=1
# $ call GetUser
evans.init:
	evans --port 50051 proto/user/user.proto
