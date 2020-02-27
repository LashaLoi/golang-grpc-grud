# Generate proto file

    protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/user.proto

# Run client

    go build cmd/client/main.go

# Run server

    go build cmd/server/main.go
