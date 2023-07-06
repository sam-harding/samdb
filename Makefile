run_local:
	@go run .

proto_build:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/commands.proto
	@echo "Proto Build Finished."