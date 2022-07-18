package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"microservices-boilerplate/internal/serviceB/service"
	assertion "microservices-boilerplate/test/assertion/serviceB"
	mock "microservices-boilerplate/test/mocks/pkg"
)

func TestServiceB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service B Suits")
}

var _ = Describe("Service B", func() {
	logMock := new(mock.Logger)
	s := service.New(logMock)

	Context("Testing CRUD Operations", func() {
		expectedItems := assertion.ItemArray
		Context("When user looks for all items", func() {
			It("Should return all items from DB", func() {
				resp, err := s.GetAll()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).Should(Equal(expectedItems))
			})
		})

		Context("When user looks for an specific item", func() {
			expectedItem := assertion.NewItemWithID(assertion.SampleID.String())
			It("Should return an item with given ID", func() {
				resp, err := s.GetOneByID(expectedItem.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).Should(Equal(expectedItem))
			})
		})

		Context("When user creates an item", func() {
			itemInput := assertion.NewItemWithoutID()
			It("Should return the created object", func() {
				resp, err := s.Create(itemInput)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp.ID).ShouldNot(BeNil())
			})
		})

		Context("When user updates an item", func() {
			inputItem := assertion.NewItemWithID(assertion.SampleID.String())
			It("Should return nothing", func() {
				err := s.Update(assertion.SampleID, inputItem)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When user deletes an item", func() {
			It("Should return nothing", func() {
				err := s.Delete(assertion.SampleID)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
