package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/google/uuid"
)

var Characters = map[string]*character.Character{}

type CharacterServer struct {
	character.UnimplementedCharacterSvcServer
}

func (s *CharacterServer) Create(ctx context.Context, in *character.CreateRequest) (*character.CreateResponse, error) {
	log.Printf("Received: %v", in.GetCharacter())
	log.Printf("Creating character: %v", in.Character.Name)

	in.Character.Id = uuid.New().String()
	in.Character.Inventory = &character.Inventory{}
	Characters[in.Character.Id] = in.Character

	return &character.CreateResponse{Character: Characters[in.Character.Id]}, nil
}

func (s *CharacterServer) Read(ctx context.Context, in *character.ReadRequest) (*character.ReadResponse, error) {
	log.Printf("Received character ID: %s", in.Id)

	c := Characters[in.Id]

	if c == nil {
		return nil, fmt.Errorf("Character not found")
	}

	return &character.ReadResponse{Character: c}, nil
}

func (s *CharacterServer) GatherResource(ctx context.Context, in *character.GatherResourceRequest) (*character.GatherResourceResponse, error) {
	c := Characters[in.CharacterId]

	fmt.Printf("Gathering resource for character: %v", c)

	for _, r := range c.Inventory.Resources {
		if r.Resource.Id == in.ResourceToGather.Id {
			r.Amount += rand.Int63n(10) + rand.Int63n(100) + 5
		}
	}

	return &character.GatherResourceResponse{Character: c, Inventory: c.Inventory}, nil
}
