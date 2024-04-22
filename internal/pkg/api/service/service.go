package service

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	account_svc_v1 "github.com/VoidMesh/backend/internal/pkg/api/service/account/v1"
	character_svc_v1 "github.com/VoidMesh/backend/internal/pkg/api/service/character/v1"
	account_api_v1 "github.com/VoidMesh/backend/pkg/api/account/v1"
	character_api_v1 "github.com/VoidMesh/backend/pkg/api/character/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Services struct {
	Account   account_api_v1.AccountSvcServer
	Character character_api_v1.CharacterSvcServer
}

func RegisterV1gRPC(ctx context.Context, s *grpc.Server) {
	// Create a new PostgreSQL connection pool
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	account_api_v1.RegisterAccountSvcServer(s, account_svc_v1.Account(db))
	character_api_v1.RegisterCharacterSvcServer(s, character_svc_v1.Character(db))
}

func RegisterV1HTTP(ctx context.Context, s *runtime.ServeMux) {
	opts := []grpc.DialOption{}

	account_api_v1.RegisterAccountSvcHandlerFromEndpoint(ctx, s, "localhost:50051", opts)
	character_api_v1.RegisterCharacterSvcHandlerFromEndpoint(ctx, s, "localhost:50051", opts)
}
