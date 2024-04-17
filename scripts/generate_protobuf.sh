#!/bin/sh

echo "Generating protobuf files..."

protoc --proto_path=api/protobuf \
       --go_out=internal/pkg/api/ \
       --go_opt=paths=source_relative \
       --go-grpc_out=internal/pkg/api \
       --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=internal/pkg/api/ \
       --grpc-gateway_opt=paths=source_relative \
       --grpc-gateway_opt=generate_unbound_methods=true \
       $(find api/protobuf -name '*.proto' -type f)

echo "Generated protobuf files successfully!"
