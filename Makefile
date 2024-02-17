GRPC_OUT_PATH_GO = ./internal/microservice-api/proto

generate_grpc_go:
	@echo "Generating gRPC code for Go"
	mkdir -p $(GRPC_OUT_PATH_GO)
	protoc -I api/ api/*.proto --go_out=$(GRPC_OUT_PATH_GO) --go_opt=paths=source_relative --go-grpc_out=$(GRPC_OUT_PATH_GO) --go-grpc_opt=paths=source_relative
