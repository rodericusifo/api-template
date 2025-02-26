.PHONY: proto-auth-service
proto-auth-service:
	protoc \
		--proto_path=external/client/auth-service/http/grpc/proto \
		--go_out=external/client/auth-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/auth-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/auth-service/http/grpc/proto/auth/auth.proto \
		external/client/auth-service/http/grpc/proto/user/user.proto
.PHONY: proto-author-service
proto-author-service:
	protoc \
		--proto_path=external/client/author-service/http/grpc/proto \
		--go_out=external/client/author-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/author-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/author-service/http/grpc/proto/author/author.proto
.PHONY: proto-category-service
proto-category-service:
	protoc \
		--proto_path=external/client/category-service/http/grpc/proto \
		--go_out=external/client/category-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/category-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/category-service/http/grpc/proto/category/category.proto
.PHONY: proto-book-service
proto-book-service:
	protoc \
		--proto_path=external/client/book-service/http/grpc/proto \
		--go_out=external/client/book-service/http/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=external/client/book-service/http/grpc/proto --go-grpc_opt=paths=source_relative \
		external/client/book-service/http/grpc/proto/book/book.proto

.PHONY: gen-proto
gen-proto: proto-auth-service proto-author-service proto-category-service proto-book-service