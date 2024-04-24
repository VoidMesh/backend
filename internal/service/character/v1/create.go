package character

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/VoidMesh/backend/api/gen/go/character/v1"
)

func (s *CharacterServer) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (resp *connect.Response[v1.CreateResponse], err error) {
	return connect.NewResponse(&v1.CreateResponse{}), nil
}
