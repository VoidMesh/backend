package account

import (
	"context"
	"fmt"

	connect "connectrpc.com/connect"
	v1 "github.com/VoidMesh/backend/api/gen/go/account/v1"
	"github.com/VoidMesh/backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// Create creates a new account with the provided email and password.
// It hashes the password, checks if the account already exists, and then creates the account.
// The function returns the created account information in the response.
// If there is an error during any step, it returns the error.
func (s *AccountServer) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (resp *connect.Response[v1.CreateResponse], err error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Msg.Password), BCRYPT_COST)
	if err != nil {
		return nil, err
	}

	// Check if account already exists
	exists, err := db.New(s.DB).CheckAccountExistsByEmail(ctx, req.Msg.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAccountAlreadyExists
	}

	// Create account
	account, err := db.New(s.DB).CreateAccount(
		ctx,
		db.CreateAccountParams{
			Email:        req.Msg.Email,
			PasswordHash: string(hashedPassword),
		},
	)
	if err != nil {
		return nil, err
	}

	// Return response
	resp = connect.NewResponse(&v1.CreateResponse{
		Id:        fmt.Sprintf("%x", account.ID.Bytes),
		Email:     account.Email,
		IsActive:  false,
		CreatedAt: account.CreatedAt.Time.String(),
		UpdatedAt: account.UpdatedAt.Time.String(),
	})
	return resp, nil
}
