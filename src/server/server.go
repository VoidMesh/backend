package main

import (
	"context"
	"log"
	"net"

	pb "github.com/VoidMesh/backend/src/api/v1"
	"github.com/google/uuid"

	"google.golang.org/grpc"
)

// server is used to implement greeter.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedCharacterServiceServer
	pb.UnimplementedResourceServiceServer
}

var Characters = map[string]*pb.Character{}

var Resources = []*pb.Resource{
	{Id: "1", Name: "Tritium", Description: "A rare and valuable resource."},
	{Id: "2", Name: "Titanium", Description: "A common and useful resource."},
}

// SayHello implements greeter.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) CreateCharacter(ctx context.Context, in *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {
	log.Printf("Received: %v", in.GetCharacter())
	log.Printf("Creating character: %v", in.Character.Name)
	in.Character.Id = uuid.New().String()
	Characters[in.Character.Name] = in.Character
	return &pb.CreateCharacterResponse{Character: Characters[in.Character.Name]}, nil
}

func (s *server) ReadCharacter(ctx context.Context, in *pb.ReadCharacterRequest) (*pb.ReadCharacterResponse, error) {
	log.Printf("Received: %v", in.GetCharacter())
	log.Printf("Getting character named: %v", in.Character.Name)
	return &pb.ReadCharacterResponse{Character: Characters[in.Character.Name]}, nil
}

func (s *server) ListResources(ctx context.Context, in *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Listing resources: %v", Resources)
	return &pb.ListResourcesResponse{Resources: Resources}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterCharacterServiceServer(s, &server{})
	pb.RegisterResourceServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
