package main

import (
	"log"
	"net"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/backend/src/api/v1/resource"
	"github.com/VoidMesh/backend/src/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	account.RegisterAccountSvcServer(s, &server.AccountServer{})
	character.RegisterCharacterSvcServer(s, &server.CharacterServer{})
	resource.RegisterResourceSvcServer(s, &server.ResourceServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
