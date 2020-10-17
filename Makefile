pb.go:
	protoc -I . \
		--go_out . \
		--go_opt plugins=grpc \
		--go_opt paths=source_relative proto/user/user.proto
