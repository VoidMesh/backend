package server

import (
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/VoidMesh/backend/src/api/v1/resource"
)

type ResourceServer struct {
	resource.UnimplementedResourceSvcServer
}

var Resources = []*resource.Resource{
	{Id: "1", Name: "Tritium", Description: "A rare and valuable resource."},
	{Id: "2", Name: "Titanium", Description: "A common and useful resource."},
}

func (s *ResourceServer) List(ctx context.Context, in *resource.ListRequest) (*resource.ListReponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Listing resources: %v", Resources)
	return &resource.ListReponse{Resources: Resources}, nil
}

func (s *ResourceServer) Get(ctx context.Context, in *resource.GetRequest) (*resource.GetResponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Getting resource: %v", in.Resource.Name)
	for _, r := range Resources {
		if r.Name == in.Resource.Name {
			return &resource.GetResponse{Resource: r}, nil
		}
	}
	return nil, errors.New("Resource not found")
}

func (s *ResourceServer) Obtain(ctx context.Context, in *resource.ObtainRequest) (*resource.ObtainResponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Obtaining resource: %v", in.Resource.Name)
	for _, r := range Resources {
		if r.Name == in.Resource.Name {
			amount := (rand.Float32() * float32(rand.Intn(150)))
			return &resource.ObtainResponse{Resource: r, Amount: amount}, nil
		}
	}
	return nil, errors.New("Resource not found")
}
