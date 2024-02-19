package server

import (
	"context"
	"fmt"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/google/uuid"
)

// TODO: Add a database to store characters
var alyxId = StoredAccounts["void-mesh@alyx.pink"].Id
var StoredCharacters = map[*account.UUID][]*character.Character{
	{Value: alyxId.Value}: {
		{
			Id:        &character.UUID{Value: uuid.NewString()},
			AccountId: &account.UUID{Value: alyxId.Value},
			Name:      "Example 1",
		},
		{
			Id:        &character.UUID{Value: uuid.NewString()},
			AccountId: &account.UUID{Value: alyxId.Value},
			Name:      "Example 2",
		},
	},
}

type CharacterServer struct {
	character.UnimplementedCharacterSvcServer
}

func (s *CharacterServer) Create(ctx context.Context, in *character.CreateRequest) (*character.CreateResponse, error) {
	log.Printf("Received: %v", in.GetCharacter())
	log.Printf("Creating character: %v", in.Character.Name)

	newCharacter := in.Character
	newCharacter.Id = &character.UUID{Value: uuid.NewString()}
	accountId := &account.UUID{Value: in.Character.AccountId.Value}

	StoredCharacters[accountId] = append(StoredCharacters[accountId], newCharacter)

	return &character.CreateResponse{Character: newCharacter}, nil
}

func (s *CharacterServer) List(ctx context.Context, in *character.ListRequest) (*character.ListResponse, error) {
	log.Printf("Listing characters for account '%s'", in.AccountId)

	c := StoredCharacters[in.AccountId]
	log.Printf("Listing characters '%v'", c)

	if c == nil {
		return nil, fmt.Errorf("Characters not found")
	}

	return &character.ListResponse{Characters: c}, nil
}
