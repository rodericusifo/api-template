.PHONY: gen-proto
gen-proto:
	protoc \
	--proto_path=internal/proto \
    --go_out=internal/proto --go_opt=paths=source_relative \
    --go-grpc_out=internal/proto --go-grpc_opt=paths=source_relative \
    internal/proto/book/book.proto