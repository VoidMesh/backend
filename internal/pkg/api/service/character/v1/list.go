package character

import (
	"context"

	v1 "github.com/VoidMesh/backend/pkg/api/character/v1"
)

func (s *character) List(ctx context.Context, req *v1.ListRequest) (resp *v1.ListResponse, err error) {
	return &v1.ListResponse{}, nil
}
