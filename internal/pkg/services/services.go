package services

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	accountv1 "github.com/VoidMesh/backend/internal/pkg/api/account/v1"
	characterv1 "github.com/VoidMesh/backend/internal/pkg/api/character/v1"
	"github.com/VoidMesh/backend/internal/pkg/services/account/v1"
	"github.com/VoidMesh/backend/internal/pkg/services/character/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Services struct {
	Account   accountv1.AccountSvcServer
	Character characterv1.CharacterSvcServer
}

func RegisterV1(ctx context.Context, s *grpc.Server) {
	// Create a new PostgreSQL connection pool
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	accountv1.RegisterAccountSvcServer(s, account.Account(db))
	characterv1.RegisterCharacterSvcServer(s, character.Character(db))
}
