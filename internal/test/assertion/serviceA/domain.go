package serviceA

import (
	"encoding/json"

	"app/internal/serviceA/domain"
)

func ItemAInBytes(item *domain.ItemA) []byte {
	b, _ := json.Marshal(item)
	return b
}

func ArrayOfItemAInBytes(arr []*domain.ItemA) []byte {
	b, _ := json.Marshal(arr)
	return b
}
