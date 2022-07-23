package serviceA

import (
	"encoding/json"
)

func ItemAInBytes(id string) []byte {
	b, _ := json.Marshal(NewItemWithID(id))
	return b
}

func ArrayOfItemAInBytes() []byte {
	b, _ := json.Marshal(ArrayOfItem)
	return b
}
