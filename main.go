package main

import (
	"log"
	"net"

	apb "github.com/VoidMesh/backend/src/api/v1/account"
	cpb "github.com/VoidMesh/backend/src/api/v1/character"
	rpb "github.com/VoidMesh/backend/src/api/v1/resource"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	apb.RegisterAccountSvcServer(s, &AccountServer{})
	cpb.RegisterCharacterSvcServer(s, &CharacterServer{})
	rpb.RegisterResourceSvcServer(s, &ResourceServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
