package domain

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const (
	FailedToUnmarshal = "failed to unmarshal data to ItemA: %v"
)

type ItemA struct {
	ID uuid.UUID `json:"id"`
}

func NewFromBytes(b []byte) (*ItemA, error) {
	var item *ItemA
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf(FailedToUnmarshal, err)
	}
	return item, nil
}

func NewArrayFromBytes(b []byte) ([]*ItemA, error) {
	var item []*ItemA
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf(FailedToUnmarshal, err)
	}
	return item, nil
}
