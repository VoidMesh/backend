package server

import (
	"context"
	"errors"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/google/uuid"
)

// TODO: Add a database to store accounts
var alyxAccountID = &account.UUID{Value: "06BD77BA-E52B-4287-A138-F0154752C701"}
var StoredAccounts = []*account.Account{
	// TODO: Remove this test account
	{
		Id:    alyxAccountID,
		Email: "void-mesh@alyx.pink",
	},
}

type AccountServer struct {
	account.UnimplementedAccountSvcServer
}

func (s *AccountServer) Create(ctx context.Context, in *account.CreateRequest) (*account.CreateResponse, error) {
	log.Printf("Creating account using email: %s", in.Email)

	if _, err := GetAccountByEmail(in.Email); err != nil {
		return nil, errors.New("Account already exists")
	}

	newAccount := account.Account{
		Id:    &account.UUID{Value: uuid.NewString()},
		Email: in.Email,
	}
	StoredAccounts = append(StoredAccounts, &newAccount)
	return &account.CreateResponse{Account: &newAccount}, nil
}

func (s *AccountServer) Authenticate(ctx context.Context, in *account.AuthenticateRequest) (*account.AuthenticateResponse, error) {
	log.Printf("Authenticating: %v", in.Email)
	var acc *account.Account

	acc, err := GetAccountByEmail(in.Email)

	if err != nil {
		return &account.AuthenticateResponse{Account: nil}, err
	}

	return &account.AuthenticateResponse{Account: acc}, nil
}

func GetAccountByEmail(email string) (*account.Account, error) {
	for _, acc := range StoredAccounts {
		if acc.Email == email {
			return acc, nil
		}
	}
	return nil, errors.New("Account not found")
}
