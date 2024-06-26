package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/VoidMesh/backend/internal/db"
	v1 "github.com/VoidMesh/backend/pkg/api/account/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (s *account) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (resp *v1.AuthenticateResponse, err error) {
	// Get account by email
	account, err := db.New(s.db).GetAccountByEmail(ctx, req.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrAccountWrongCredentials
	}
	if err != nil {
		return nil, err
	}

	// Account is verified once they clicked the verification link
	// TODO: Send verification email
	if !account.IsVerified.Bool {
		return nil, ErrAccountNotVerified
	}

	// Compare password hashes
	err = bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(req.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, ErrAccountWrongCredentials
	}
	if err != nil {
		return nil, err
	}

	//
	// JWT Tokens generation
	//

	// Generate long-live refresh token
	refreshToken, refreshTokenJWT, err := generateJWT(account.ID.Bytes, JWT_REFRESH_DURATION)
	if err != nil {
		return nil, ErrGeneratingRefreshToken
	}

	// Generate short-lived access token
	accessToken, _, err := generateJWT(account.ID.Bytes, JWT_ACCESS_TOKEN_DURATION)
	if err != nil {
		return nil, ErrGeneratingAccessToken
	}

	// Get issued and expiration time of the refresh token
	refreshTokenIssuedAt, err := refreshTokenJWT.Claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}
	refreshTokenExpiresAt, err := refreshTokenJWT.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	db.New(s.db).CreateAccountSession(ctx, db.CreateAccountSessionParams{
		AccountID:    account.ID,
		RefreshToken: refreshToken,
		UserAgent:    pgtype.Text{String: req.UserAgent, Valid: true},
		IpAddress:    pgtype.Text{String: req.IpAddress, Valid: true},
		IssuedAt:     pgtype.Timestamp{Time: refreshTokenIssuedAt.Time, Valid: true},
		ExpiresAt:    pgtype.Timestamp{Time: refreshTokenExpiresAt.Time, Valid: true},
	})

	// Return response
	resp = &v1.AuthenticateResponse{
		Id:           fmt.Sprintf("%x", account.ID.Bytes),
		IsActive:     account.IsActive.Bool,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    account.CreatedAt.Time.String(),
		UpdatedAt:    account.UpdatedAt.Time.String(),
	}
	return resp, nil
}
