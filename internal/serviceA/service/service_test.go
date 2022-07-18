package service

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	assertion "microservices-boilerplate/testing/assertion/serviceA"
	mock "microservices-boilerplate/testing/mocks/pkg"
	"testing"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServiceA Suits")
}

var _ = Describe("Testing ServiceA Service", func() {
	logMock := new(mock.Logger)
	service := New(logMock)

	Context("Testing CRUD Operations", func() {
		Context("When user looks for all items", func() {
			It("Should return all items from DB", func() {
				resp, err := service.GetAll()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).Should(Equal(assertion.ItemArray))
			})
		})

		Context("When user looks for an specific item", func() {
			expectedItem := assertion.NewItemWithID(assertion.SampleID.String())
			It("Should return an item with given ID", func() {
				resp, err := service.GetOneByID(expectedItem.ID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).Should(Equal(expectedItem))
			})
		})

		Context("When user creates an item", func() {
			itemInput := assertion.NewItemWithoutID()
			It("Should return the created object", func() {
				resp, err := service.Create(itemInput)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp.ID).ShouldNot(BeNil())
			})
		})

		Context("When user updates an item", func() {
			inputItem := assertion.NewItemWithID(assertion.SampleID.String())
			It("Should return nothing", func() {
				err := service.Update(assertion.SampleID, inputItem)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When user deletes an item", func() {
			It("Should return nothing", func() {
				err := service.Delete(assertion.SampleID)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
