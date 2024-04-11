package character

import (
	v1 "github.com/VoidMesh/backend/internal/pkg/api/character/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

type character struct {
	v1.UnimplementedCharacterSvcServer
	db *pgxpool.Pool
}

func Character(db *pgxpool.Pool) v1.CharacterSvcServer {
	return &character{db: db}
}
