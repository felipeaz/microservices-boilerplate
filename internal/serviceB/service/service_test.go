package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microservices-boilerplate/internal/serviceB/service"
	assertion "microservices-boilerplate/internal/test/assertion/serviceB"
	pkgMock "microservices-boilerplate/internal/test/mocks/pkg"
	repositoryMock "microservices-boilerplate/internal/test/mocks/serviceB/repository"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suits")
}

var _ = Describe("Service", func() {
	logMock := new(pkgMock.Logger)
	repoMock := new(repositoryMock.Repository)
	s := service.New(logMock, repoMock)

	Context("Testing CRUD Operations", func() {

		Context("Getting All items", func() {
			When("Request succeeds", func() {
				expectedItems := assertion.ArrayOfItem
				It("Should return all items from DB", func() {
					resp, err := s.GetAll(assertion.Ctx)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItems))
				})
			})
		})

		Context("Getting one item by ID", func() {
			When("Request succeeds", func() {
				idString := assertion.SampleID.String()
				expectedItem := assertion.NewItemWithID(idString)
				It("Should return an item with given ID", func() {
					resp, err := s.GetOneByID(assertion.Ctx, expectedItem.ID.String())
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItem))
				})
			})
		})

		Context("Creating an item", func() {
			When("Request succeeds", func() {
				itemInput := assertion.NewItemWithoutID()
				It("Should return the created object", func() {
					resp, err := s.Create(assertion.Ctx, itemInput)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp.ID).NotTo(BeNil())
				})
			})
		})

		Context("Updating an item", func() {
			When("Request succeeds", func() {
				idString := assertion.SampleID.String()
				inputItem := assertion.NewItemWithID(idString)
				It("Should return nothing", func() {
					err := s.Update(assertion.Ctx, idString, inputItem)
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Context("Deleting an item", func() {
			When("Request succeeds", func() {
				It("Should return nothing", func() {
					err := s.Delete(assertion.Ctx, assertion.SampleID.String())
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
