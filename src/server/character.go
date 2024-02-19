package server

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/google/uuid"
)

// TODO: Add a database to store characters
var StoredCharacters = []*character.Character{
	{
		Id:        &character.UUID{Value: "172C0434-CFCB-459E-9501-0168269D324F"},
		AccountId: alyxAccountID,
		Name:      "Example 1",
	},
	{
		Id:        &character.UUID{Value: "C278FE2C-26C4-44E3-9BB0-6658147F59D5"},
		AccountId: alyxAccountID,
		Name:      "Example 2",
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

	StoredCharacters = append(StoredCharacters, newCharacter)

	return &character.CreateResponse{Character: newCharacter}, nil
}

func (s *CharacterServer) List(ctx context.Context, in *character.ListRequest) (*character.ListResponse, error) {
	log.Printf("Listing characters for account '%s'", in.AccountId)

	c, err := GetCharacterByAccountUuid(in.AccountId)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	log.Printf("Listing characters '%v'", c)

	return &character.ListResponse{Characters: c}, nil
}

func GetCharacterByAccountUuid(uuid *account.UUID) ([]*character.Character, error) {
	var chars []*character.Character
	for _, char := range StoredCharacters {
		if char.AccountId.Value == uuid.Value {
			chars = append(chars, char)
		}
	}

	if len(chars) == 0 {
		return nil, errors.New("Characters not found")
	}

	return chars, nil
}
