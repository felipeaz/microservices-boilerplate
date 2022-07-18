package serviceA

import (
	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceA/domain"
)

func NewItemWithID(id string) domain.ItemA {
	return domain.ItemA{
		ID: uuid.FromStringOrNil(id),
	}
}

func NewItemWithoutID() domain.ItemA {
	return domain.ItemA{}
}

var ItemArray = []domain.ItemA{
	NewItemWithID("481da253-2dda-46e5-9963-58611eb72d7b"),
	NewItemWithID("2a2acd06-c4ce-4bce-aaf9-09a379f02cf8"),
	NewItemWithID("667f4eda-6825-445c-bf45-289f3b64b02b"),
	NewItemWithID("4740de96-9068-4a3a-bdc6-132ad7c58bae"),
}

var SampleID = uuid.FromStringOrNil("15664c2f-d5bf-4922-8d19-39c6886bce90")
