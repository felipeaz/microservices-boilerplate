package serviceB

import (
	"encoding/json"
)

func ItemBInBytes(id string) []byte {
	b, _ := json.Marshal(NewItemWithID(id))
	return b
}

func ArrayOfItemBInBytes() []byte {
	b, _ := json.Marshal(ArrayOfItem)
	return b
}
