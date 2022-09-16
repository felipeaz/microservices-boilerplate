package serviceB

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceB/domain"
)

var (
	ArrayOfItem = []*domain.ItemB{
		NewItemWithID("481da253-2dda-46e5-9963-58611eb72d7b"),
		NewItemWithID("2a2acd06-c4ce-4bce-aaf9-09a379f02cf8"),
		NewItemWithID("667f4eda-6825-445c-bf45-289f3b64b02b"),
		NewItemWithID("4740de96-9068-4a3a-bdc6-132ad7c58bae"),
	}

	SampleID        = uuid.FromStringOrNil("15664c2f-d5bf-4922-8d19-39c6886bce90")
	InvalidIDString = "15664c2f"
)

func NewItemWithID(id string) *domain.ItemB {
	return &domain.ItemB{
		ID: uuid.FromStringOrNil(id),
	}
}

func NewItemFromInput(input *domain.ItemB) *domain.ItemB {
	newItem := input
	if len(input.ID.Bytes()) == 0 {
		newItem.ID = uuid.NewV4()
	}
	return newItem
}

func NewItemWithoutID() *domain.ItemB {
	return &domain.ItemB{}
}

func NewErrIncorrectIDLength(id string) error {
	return fmt.Errorf("uuid: incorrect UUID length: %s", id)
}
