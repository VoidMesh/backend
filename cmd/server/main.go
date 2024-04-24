package main

import (
	"context"
	"log"
	"net/http"

	"github.com/VoidMesh/backend/api/gen/go/account/v1/accountv1connect"
	"github.com/VoidMesh/backend/api/gen/go/character/v1/characterv1connect"

	"github.com/VoidMesh/backend/internal/service/account/v1"
	"github.com/VoidMesh/backend/internal/service/character/v1"

	"github.com/VoidMesh/backend/internal/db"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// func oldmain() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// Create a listener on TCP port for gRPC server
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v\n", err)
// 	}
// 	defer lis.Close()

// 	// Create a new gRPC server
// 	grpcServer := grpc.NewServer()
// 	defer grpcServer.GracefulStop()

// 	// Register reflection service
// 	reflection.Register(grpcServer)

// 	// Register health check service
// 	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

// 	// Register V1 services
// 	service.RegisterV1(ctx, grpcServer)

// 	// Serve the gRPC server
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

func main() {
	ctx := context.Background()

	// Initialize database connection
	db, err := db.NewDatabase(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(serviceNames()...)))
	mux.Handle(grpcreflect.NewHandlerV1(grpcreflect.NewStaticReflector(serviceNames()...)))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(grpcreflect.NewStaticReflector(serviceNames()...)))

	mux.Handle(accountv1connect.NewAccountServiceHandler(&account.AccountServer{DB: db}))
	mux.Handle(characterv1connect.NewCharacterServiceHandler(&character.CharacterServer{DB: db}))

	http.ListenAndServe(
		"0.0.0.0:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func serviceNames() []string {
	return []string{
		accountv1connect.AccountServiceName,
		characterv1connect.CharacterServiceName,
	}
}
