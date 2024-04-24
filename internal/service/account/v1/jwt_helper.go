package account

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a JWT token for the account
func generateJWT(accountIdBytes [16]byte, duration time.Duration) (string, *jwt.Token, error) {

	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%x", accountIdBytes),
		Issuer:    "voidmesh_backend",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", nil, err
	}

	return tokenString, token, nil
}
