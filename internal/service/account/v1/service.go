package account

import (
	"errors"
	"os"
	"time"

	"github.com/VoidMesh/backend/api/gen/go/account/v1/accountv1connect"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	BCRYPT_COST               = 14
	JWT_REFRESH_DURATION      = 24 * time.Hour * 365 // Token is valid for 365 days
	JWT_ACCESS_TOKEN_DURATION = 15 * time.Minute     // Token is valid for 15 minutes
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

var (
	ErrAccountAlreadyExists               = errors.New("account already exists")
	ErrAccountNotVerified                 = errors.New("account not verified")
	ErrAccountSessionRefreshTokenNotFound = errors.New("refresh token not foud")
	ErrAccountWrongCredentials            = errors.New("email or password is incorrect")

	ErrGeneratingAccessToken  = errors.New("failed to generate access token")
	ErrGeneratingRefreshToken = errors.New("failed to generate refresh token")

	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrExpiredRefreshToken = errors.New("refresh token expired")
)

type AccountServer struct {
	accountv1connect.UnimplementedAccountServiceHandler
	DB *pgxpool.Pool
}
