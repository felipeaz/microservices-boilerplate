package domain

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const (
	FailedToUnmarshal = "failed to unmarshal data to ItemB: %v"
)

type ItemB struct {
	ID uuid.UUID `json:"id"`
}

func NewFromBytes(b []byte) (*ItemB, error) {
	var item *ItemB
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf(FailedToUnmarshal, err)
	}
	return item, nil
}

func NewArrayFromBytes(b []byte) ([]*ItemB, error) {
	var item []*ItemB
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf(FailedToUnmarshal, err)
	}
	return item, nil
}
