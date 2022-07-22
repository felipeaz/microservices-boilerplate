package domain

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type ItemA struct {
	ID uuid.UUID `json:"id"`
}

func NewFromBytes(b []byte) (*ItemA, error) {
	var item *ItemA
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to ItemA: %v", err)
	}
	return item, nil
}

func NewArrayFromBytes(b []byte) ([]*ItemA, error) {
	var item []*ItemA
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to ItemA: %v", err)
	}
	return item, nil
}
