export PATH=$PATH:$HOME/go/bin
export PATH=\$PATH:/usr/local/go/bin
protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/user.proto
