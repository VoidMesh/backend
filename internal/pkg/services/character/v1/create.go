package character

import (
	"context"

	v1 "github.com/VoidMesh/backend/internal/pkg/api/character/v1"
)

func (s *character) Create(ctx context.Context, req *v1.CreateRequest) (resp *v1.CreateResponse, err error) {
	return &v1.CreateResponse{}, nil
}
