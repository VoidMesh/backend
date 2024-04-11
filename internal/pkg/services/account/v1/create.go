package account

import (
	"context"
	"fmt"

	"github.com/VoidMesh/backend/internal/db"
	v1 "github.com/VoidMesh/backend/internal/pkg/api/account/v1"
	"golang.org/x/crypto/bcrypt"
)

// Create creates a new account with the provided email and password.
// It hashes the password, checks if the account already exists, and then creates the account.
// The function returns the created account information in the response.
// If there is an error during any step, it returns the error.
func (s *account) Create(ctx context.Context, req *v1.CreateRequest) (resp *v1.CreateResponse, err error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), BCRYPT_COST)
	if err != nil {
		return nil, err
	}

	// Check if account already exists
	exists, err := db.New(s.db).CheckAccountExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAccountAlreadyExists
	}

	// Create account
	account, err := db.New(s.db).CreateAccount(
		ctx,
		db.CreateAccountParams{
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
		},
	)
	if err != nil {
		return nil, err
	}

	// Return response
	resp = &v1.CreateResponse{
		Id:        fmt.Sprintf("%x", account.ID.Bytes),
		Email:     account.Email,
		IsActive:  false,
		CreatedAt: account.CreatedAt.Time.String(),
		UpdatedAt: account.UpdatedAt.Time.String(),
	}
	return resp, nil
}
