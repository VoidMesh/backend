package server

import (
	"context"
	"errors"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/resource"
	"github.com/google/uuid"
)

// TODO: Add a database to store resources
var StoredResources = []*resource.Resource{
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Tritium", Description: "A rare and valuable resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Titanium", Description: "A common and useful resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Gold", Description: "A valuable resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Iron", Description: "A common resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Copper", Description: "A common resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Silver", Description: "A valuable resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Platinum", Description: "A rare and valuable resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Uranium", Description: "A rare and dangerous resource."},
	{Id: &resource.UUID{Value: uuid.NewString()}, Name: "Plutonium", Description: "A rare and dangerous resource."},
}

type ResourceServer struct {
	resource.UnimplementedResourceSvcServer
}

func (s *ResourceServer) List(ctx context.Context, in *resource.ListRequest) (*resource.ListReponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Listing resources: %v", StoredResources)
	return &resource.ListReponse{Resources: StoredResources}, nil
}

func (s *ResourceServer) Get(ctx context.Context, in *resource.GetRequest) (*resource.GetResponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Getting resource: %v", in.Resource.Name)
	for _, r := range StoredResources {
		if r.Name == in.Resource.Name {
			return &resource.GetResponse{Resource: r}, nil
		}
	}
	return nil, errors.New("Resource not found")
}
