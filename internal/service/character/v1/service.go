package character

import (
	"github.com/VoidMesh/backend/api/gen/go/character/v1/characterv1connect"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CharacterServer struct {
	characterv1connect.UnimplementedCharacterServiceHandler
	DB *pgxpool.Pool
}
