package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/VoidMesh/backend/internal/db"
	v1 "github.com/VoidMesh/backend/pkg/api/account/v1"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

func (s *account) RefreshAccessToken(ctx context.Context, req *v1.RefreshAccessTokenRequest) (resp *v1.RefreshAccessTokenResponse, err error) {
	// Get account session by refresh token
	accountSession, err := db.New(s.db).GetAccountSessionByRefreshToken(ctx, req.RefreshToken)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrAccountSessionRefreshTokenNotFound
	}
	if err != nil {
		return nil, err
	}

	// Parse the Refresh Token
	refreshToken, err := jwt.ParseWithClaims(req.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
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
		if !(string(req.RefreshToken) == accountSession.RefreshToken) && !(claims.Subject == fmt.Sprintf("%x", accountSession.AccountID.Bytes)) {
			return nil, ErrInvalidRefreshToken
		}

		// Generate a new Access Token
		newAccessToken, _, err := generateJWT(accountSession.AccountID.Bytes, JWT_ACCESS_TOKEN_DURATION)
		if err != nil {
			return nil, err
		}

		return &v1.RefreshAccessTokenResponse{AccessToken: newAccessToken}, nil
	} else {
		return nil, ErrInvalidRefreshToken
	}
}
