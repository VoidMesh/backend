package server

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/backend/src/api/v1/inventory"
	"github.com/google/uuid"
)

// TODO: Add a database to store inventories
var StoredInventories = []*inventory.Inventory{
	{
		Id:          &inventory.UUID{Value: uuid.NewString()},
		CharacterId: &character.UUID{Value: "172C0434-CFCB-459E-9501-0168269D324F"},
		Slots: []*inventory.Slot{
			{Resource: StoredResources[0], Quantity: 3738},
			{Resource: StoredResources[1], Quantity: 1413},
			{Resource: StoredResources[2], Quantity: 54326},
			{Resource: StoredResources[3], Quantity: 8374},
		},
	},
	{
		Id:          &inventory.UUID{Value: uuid.NewString()},
		CharacterId: &character.UUID{Value: "C278FE2C-26C4-44E3-9BB0-6658147F59D5"},
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
		inv, err := GetInventoryByCharacterUuid(in.CharacterId)
		return &inventory.ReadResponse{Inventory: inv}, err
	}

	// If character ID is provided, return the inventory with the given character ID
	if in.CharacterId != nil {
		inv, err := GetInventoryByCharacterUuid(in.CharacterId)
		return &inventory.ReadResponse{Inventory: inv}, err
	}

	return nil, fmt.Errorf("Inventory not found")
}

func GetInventoryByCharacterUuid(uuid *character.UUID) (*inventory.Inventory, error) {
	for _, inv := range StoredInventories {
		if inv.CharacterId.Value == uuid.Value {
			return inv, nil
		}
	}
	return nil, errors.New("Inventory not found")
}

func GetInventoryByUuid(uuid *inventory.UUID) (*inventory.Inventory, error) {
	for _, inv := range StoredInventories {
		if inv.Id.Value == uuid.Value {
			return inv, nil
		}
	}
	return nil, errors.New("Inventory not found")
}
