package account

import (
	"context"

	connect "connectrpc.com/connect"
	v1 "github.com/VoidMesh/backend/api/gen/go/account/v1"
	"github.com/VoidMesh/backend/internal/db"
)

func (s *AccountServer) DeleteRefreshToken(ctx context.Context, req *connect.Request[v1.DeleteRefreshTokenRequest]) (resp *connect.Response[v1.DeleteRefreshTokenResponse], err error) {
	// Delete account session by refresh token
	err = db.New(s.DB).DeleteAccountSession(ctx, req.Msg.RefreshToken)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteRefreshTokenResponse{}), nil
}
