package server

import (
	"context"
	"errors"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/google/uuid"
)

// TODO: Add a database to store accounts
var StoredAccounts = map[string]*account.Account{
	// TODO: Remove this test account
	"void-mesh@alyx.pink": {
		Id:    &account.UUID{Value: uuid.NewString()},
		Email: "void-mesh@alyx.pink",
	},
}

type AccountServer struct {
	account.UnimplementedAccountSvcServer
}

func (s *AccountServer) Create(ctx context.Context, in *account.CreateRequest) (*account.CreateResponse, error) {
	log.Printf("Creating account using email: %s", in.Email)

	if _, ok := StoredAccounts[in.Email]; ok {
		return nil, errors.New("Account already exists")
	}

	newAccount := account.Account{
		Id:    &account.UUID{Value: uuid.NewString()},
		Email: in.Email,
	}
	StoredAccounts[newAccount.Email] = &newAccount
	return &account.CreateResponse{Account: &newAccount}, nil
}

func (s *AccountServer) Authenticate(ctx context.Context, in *account.AuthenticateRequest) (*account.AuthenticateResponse, error) {
	log.Printf("Authenticating: %v", in.Email)

	if _, ok := StoredAccounts[in.Email]; !ok {
		return &account.AuthenticateResponse{
			Account: nil,
		}, errors.New("Account not found")
	}
	log.Printf("ID for '%s': %s", in.Email, StoredAccounts[in.Email].Id.Value)

	return &account.AuthenticateResponse{
		Account: StoredAccounts[in.Email],
	}, nil
}
