package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/VoidMesh/backend/internal/pkg/api/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a listener on TCP port for gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	defer lis.Close()

	// Create a new request multiplexer for grpc-gateway
	mux := runtime.NewServeMux()

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	// Register reflection service
	reflection.Register(grpcServer)

	// Register health check service
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	// Register V1 services
	service.RegisterV1gRPC(ctx, grpcServer)
	service.RegisterV1HTTP(ctx, mux)

	// Serve the gRPC server
	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Serve the gRPC-Gateway server
	go func() {
		log.Printf("gateway listening at %v", ":9000")
		if err := http.ListenAndServe(":9000", mux); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Block forever (or until an interrupt signal is received)
	select {}
}
