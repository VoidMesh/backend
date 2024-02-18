package server

import (
	"context"
	"errors"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/google/uuid"
)

// TODO: Add a database to store accounts
var Accounts = map[string]*account.Account{
	// TODO: Remove this test account
	"void-mesh@alyx.pink": {
		Id:    uuid.NewString(),
		Email: "void-mesh@alyx.pink",
		Characters: []*character.Character{
			{Id: uuid.NewString(), Name: "Example 1"},
			{Id: uuid.NewString(), Name: "Example 2"},
		},
	},
}

type AccountServer struct {
	account.UnimplementedAccountSvcServer
}

func (s *AccountServer) Create(ctx context.Context, in *account.CreateRequest) (*account.CreateResponse, error) {
	log.Printf("Creating account using email: %s", in.Email)

	if _, ok := Accounts[in.Email]; ok {
		return nil, errors.New("Account already exists")
	}

	newAccount := account.Account{
		Id:    uuid.NewString(),
		Email: in.Email,
	}
	Accounts[newAccount.Email] = &newAccount
	return &account.CreateResponse{Account: &newAccount}, nil
}

func (s *AccountServer) Authenticate(ctx context.Context, in *account.AuthenticateRequest) (*account.AuthenticateResponse, error) {
	log.Printf("Authenticating: %v", in.Email)

	if _, ok := Accounts[in.Email]; !ok {
		return &account.AuthenticateResponse{
			Account: nil,
		}, errors.New("Account not found")
	}

	return &account.AuthenticateResponse{
		Account: Accounts[in.Email],
	}, nil
}

func (s *AccountServer) Get(ctx context.Context, in *account.GetRequest) (*account.GetResponse, error) {
	log.Printf("Getting account: %v", in.Account.Email)

	if _, ok := Accounts[in.Account.Email]; !ok {
		return &account.GetResponse{
			Account: nil,
		}, nil
	}

	return &account.GetResponse{Account: Accounts[in.Account.Email]}, nil
}
