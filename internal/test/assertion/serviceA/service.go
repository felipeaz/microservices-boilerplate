package serviceA

import (
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"

	"microservices-boilerplate/internal/serviceA/domain"
)

var (
	ArrayOfItem = []*domain.ItemA{
		NewItemWithID("481da253-2dda-46e5-9963-58611eb72d7b"),
		NewItemWithID("2a2acd06-c4ce-4bce-aaf9-09a379f02cf8"),
		NewItemWithID("667f4eda-6825-445c-bf45-289f3b64b02b"),
		NewItemWithID("4740de96-9068-4a3a-bdc6-132ad7c58bae"),
	}

	SampleID        = uuid.FromStringOrNil("15664c2f-d5bf-4922-8d19-39c6886bce90")
	InvalidIDString = "15664c2f"

	Ctx = context.Background()

	ErrGeneric      = errors.New("generic error")
	ErrNotFound     = gorm.ErrRecordNotFound
	ErrCreatingUUID = errors.New("failed to create UUID from string")
)

func NewItemWithID(id string) *domain.ItemA {
	return &domain.ItemA{
		ID: uuid.FromStringOrNil(id),
	}
}

func NewItemFromInput(input *domain.ItemA) *domain.ItemA {
	newItem := input
	if len(input.ID.Bytes()) == 0 {
		newItem.ID = uuid.NewV4()
	}
	return newItem
}

func NewItemWithoutID() *domain.ItemA {
	return &domain.ItemA{}
}

func NewErrIncorrectIDLength(id string) error {
	return fmt.Errorf("uuid: incorrect UUID length: %s", id)
}
