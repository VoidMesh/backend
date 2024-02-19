package server

import (
	"context"
	"fmt"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/backend/src/api/v1/inventory"
	"github.com/google/uuid"
)

// TODO: Add a database to store inventories
var StoredInventories = map[*character.UUID]inventory.Inventory{
	{Value: alyxId.Value}: {
		Id: &inventory.UUID{Value: uuid.NewString()},
		Slots: []*inventory.Slot{
			{Resource: StoredResources[0], Quantity: 3738},
			{Resource: StoredResources[1], Quantity: 1413},
			{Resource: StoredResources[2], Quantity: 54326},
			{Resource: StoredResources[3], Quantity: 8374},
		},
	},
	{Value: alyxId.Value}: {
		Id: &inventory.UUID{Value: uuid.NewString()},
		Slots: []*inventory.Slot{
			{Resource: StoredResources[0], Quantity: 23523},
			{Resource: StoredResources[2], Quantity: 6532},
			{Resource: StoredResources[4], Quantity: 8238},
			{Resource: StoredResources[6], Quantity: 9346},
		},
	},
}

type InventoryServer struct {
	inventory.UnimplementedInventorySvcServer
}

func (s *InventoryServer) Read(ctx context.Context, in *inventory.ReadRequest) (*inventory.ReadResponse, error) {
	log.Printf("Received: %v", in)
	log.Printf("Getting inventories: %v", StoredInventories)

	// If inventory ID is provided, return the inventory with the given ID
	if in.Id != nil {
		for _, inv := range StoredInventories {
			if inv.Id.Value == in.Id.Value {
				return &inventory.ReadResponse{Inventory: &inv}, nil
			}
		}
	}

	// If character ID is provided, return the inventory with the given character ID
	if in.CharacterId != nil {
		for _, inv := range StoredInventories {
			if inv.CharacterId.Value == in.CharacterId.Value {
				return &inventory.ReadResponse{Inventory: &inv}, nil
			}
		}
	}

	return nil, fmt.Errorf("Inventory not found")
}
