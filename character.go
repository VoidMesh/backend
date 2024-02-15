package main

import (
	"context"
	"log"

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
	Characters[in.Character.Name] = in.Character
	return &character.CreateResponse{Character: Characters[in.Character.Name]}, nil
}

func (s *CharacterServer) Read(ctx context.Context, in *character.ReadRequest) (*character.ReadResponse, error) {
	log.Printf("Received: %v", in.GetCharacter())
	log.Printf("Getting character named: %v", in.Character.Name)
	return &character.ReadResponse{Character: Characters[in.Character.Name]}, nil
}
