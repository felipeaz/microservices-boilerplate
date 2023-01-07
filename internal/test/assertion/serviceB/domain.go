package serviceB

import (
	"encoding/json"

	"app/internal/serviceB/domain"
)

func ItemBInBytes(item *domain.ItemB) []byte {
	b, _ := json.Marshal(item)
	return b
}

func ArrayOfItemBInBytes(arr []*domain.ItemB) []byte {
	b, _ := json.Marshal(ArrayOfItem)
	return b
}
