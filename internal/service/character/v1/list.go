package character

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/VoidMesh/backend/api/gen/go/character/v1"
)

func (s *CharacterServer) List(ctx context.Context, req *connect.Request[v1.ListRequest]) (resp *connect.Response[v1.ListResponse], err error) {
	return connect.NewResponse(&v1.ListResponse{}), nil
}
