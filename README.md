gRPC

```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

```bash
protoc --proto_path=proto --go_out=. --go-grpc_out=. Count.proto Hello.proto Average.proto Conversation.proto
```

```bash
go build -o server -tags=server && ./server
```

```bash
# unary
go build -o client -tags=client && ./client Mikaar

# server streaming
go build -o client -tags=client && ./client count

# client streaming
go build -o client -tags=client && ./client avg

# bidirectional streaming
go build -o client -tags=client && ./client talk
```