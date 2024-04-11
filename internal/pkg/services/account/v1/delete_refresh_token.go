package account

import (
	"context"

	"github.com/VoidMesh/backend/internal/db"
	v1 "github.com/VoidMesh/backend/internal/pkg/api/account/v1"
)

func (s *account) DeleteRefreshToken(ctx context.Context, req *v1.DeleteRefreshTokenRequest) (resp *v1.DeleteRefreshTokenResponse, err error) {
	// Delete account session by refresh token
	err = db.New(s.db).DeleteAccountSession(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRefreshTokenResponse{}, nil
}
