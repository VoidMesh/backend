package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	connect "connectrpc.com/connect"
	v1 "github.com/VoidMesh/backend/api/gen/go/account/v1"
	"github.com/VoidMesh/backend/internal/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

func (s *AccountServer) RefreshAccessToken(ctx context.Context, req *connect.Request[v1.RefreshAccessTokenRequest]) (resp *connect.Response[v1.RefreshAccessTokenResponse], err error) {
	// Get account session by refresh token
	accountSession, err := db.New(s.DB).GetAccountSessionByRefreshToken(ctx, req.Msg.RefreshToken)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrAccountSessionRefreshTokenNotFound
	}
	if err != nil {
		return nil, err
	}

	// Parse the Refresh Token
	refreshToken, err := jwt.ParseWithClaims(req.Msg.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := refreshToken.Claims.(*jwt.RegisteredClaims); ok && refreshToken.Valid {
		// Check if the Refresh Token is expired or not
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, ErrExpiredRefreshToken
		}

		// Check if the Refresh Token is associated with the correct Account ID
		if !(string(req.Msg.RefreshToken) == accountSession.RefreshToken) && !(claims.Subject == fmt.Sprintf("%x", accountSession.AccountID.Bytes)) {
			return nil, ErrInvalidRefreshToken
		}

		// Generate a new Access Token
		newAccessToken, _, err := generateJWT(accountSession.AccountID.Bytes, JWT_ACCESS_TOKEN_DURATION)
		if err != nil {
			return nil, err
		}

		return connect.NewResponse(&v1.RefreshAccessTokenResponse{AccessToken: newAccessToken}), nil
	} else {
		return nil, ErrInvalidRefreshToken
	}
}
