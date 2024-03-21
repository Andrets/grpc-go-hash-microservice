PHONY: generate-structs
generate-structs:
				mkdir -p pkg/hash_v1
				protoc --go_out=pkg/hash_v1 --go_opt=paths=source_relative \
				        internal/proto/hash.proto

PHONY: generate
generate:
				mkdir -p pkg/hash_v1
				protoc --go_out=pkg/hash_v1 --go_opt=paths=import \
								--go-grpc_out=pkg/hash_v1 --go_opt=paths=import \
				        internal/proto/hash.proto
				mv pkg/hash_v1/github.com/Andrets/grpc-go-hash-microservice/pkg/hash_v1/* pkg/hash_v1/
				rm -rf pkg/hash_v1/github.com