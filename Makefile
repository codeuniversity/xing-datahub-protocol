proto:
	protoc --go_out=. protocol.proto

test:
	go test ./...
