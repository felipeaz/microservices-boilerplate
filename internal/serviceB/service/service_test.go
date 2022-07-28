package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"microservices-boilerplate/internal/serviceB/service"
	assertionErrors "microservices-boilerplate/internal/test/assertion/errors"
	assertion "microservices-boilerplate/internal/test/assertion/serviceB"
	pkgMock "microservices-boilerplate/internal/test/mocks/pkg"
	repositoryMock "microservices-boilerplate/internal/test/mocks/serviceB/repository"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suits")
}

var _ = Describe("Service", func() {
	var (
		logMock  *pkgMock.Logger
		repoMock *repositoryMock.Repository
		s        service.Service
	)

	BeforeEach(func() {
		logMock = pkgMock.NewLogger(GinkgoT())
		repoMock = repositoryMock.NewRepository(GinkgoT())
		s = service.New(logMock, repoMock)
	})

	Context("Testing CRUD Operations", func() {
		Context("Getting All items", func() {
			When("Request succeeds", func() {
				It("Should return all items from DB", func() {
					expectedItems := assertion.ArrayOfItem
					repoMock.On("GetAll", assertion.Ctx).
						Return(expectedItems, nil).
						Once()

					resp, err := s.GetAll(assertion.Ctx)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItems))
				})
			})
			When("DB is empty", func() {
				It("Should an empty array", func() {
					repoMock.On("GetAll", assertion.Ctx).
						Return(nil, nil).
						Once()

					resp, err := s.GetAll(assertion.Ctx)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(BeNil())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					repoMock.On("GetAll", assertion.Ctx).
						Return(nil, assertionErrors.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, assertionErrors.ErrGeneric).
						Once()

					resp, err := s.GetAll(assertion.Ctx)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrGeneric))
					Expect(resp).To(BeNil())
				})
			})
		})

		Context("Getting one item by ID", func() {
			When("Request succeeds", func() {
				It("Should return an item with given ID", func() {
					idString := assertion.SampleID.String()
					expectedItem := assertion.NewItemWithID(idString)
					repoMock.On("GetByID", assertion.Ctx, assertion.SampleID).
						Return(expectedItem, nil).
						Once()

					resp, err := s.GetOneByID(assertion.Ctx, assertion.SampleID.String())

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItem))
				})
			})
			When("Item is not found", func() {
				It("Should return a not found error", func() {
					repoMock.On("GetByID", assertion.Ctx, assertion.SampleID).
						Return(nil, assertionErrors.ErrNotFound).
						Once()
					logMock.On("Error", mock.Anything, assertion.SampleID, assertionErrors.ErrNotFound).
						Once()

					resp, err := s.GetOneByID(assertion.Ctx, assertion.SampleID.String())

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrNotFound))
					Expect(resp).To(BeNil())
				})
			})
			When("Fails to parse UUID from string", func() {
				It("Should return an error", func() {
					logMock.On(
						"Error",
						mock.Anything,
						assertion.NewErrIncorrectIDLength(assertion.InvalidIDString),
					).Once()

					resp, err := s.GetOneByID(assertion.Ctx, assertion.InvalidIDString)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrCreatingUUID))
					Expect(resp).To(BeNil())
				})
			})
		})

		Context("Creating an item", func() {
			When("Request succeeds", func() {
				It("Should return the created object", func() {
					itemInput := assertion.NewItemWithoutID()
					expectedItem := assertion.NewItemFromInput(itemInput)
					repoMock.On("Insert", assertion.Ctx, itemInput).
						Return(expectedItem, nil).
						Once()

					resp, err := s.Create(assertion.Ctx, itemInput)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItem))
					Expect(resp.ID).NotTo(BeEmpty())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					itemInput := assertion.NewItemWithoutID()
					repoMock.On("Insert", assertion.Ctx, itemInput).
						Return(nil, assertionErrors.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, itemInput, assertionErrors.ErrGeneric).
						Return().
						Once()

					resp, err := s.Create(assertion.Ctx, itemInput)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrGeneric))
					Expect(resp).To(BeNil())
				})
			})
		})

		Context("Updating an item", func() {
			When("Request succeeds", func() {
				It("Should return nothing", func() {
					idString := assertion.SampleID.String()
					inputItem := assertion.NewItemWithID(idString)
					repoMock.On("Update", assertion.Ctx, assertion.SampleID, inputItem).
						Return(nil).
						Once()

					err := s.Update(assertion.Ctx, idString, inputItem)
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					idString := assertion.SampleID.String()
					inputItem := assertion.NewItemWithID(idString)
					repoMock.On("Update", assertion.Ctx, assertion.SampleID, inputItem).
						Return(assertionErrors.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, assertion.SampleID, inputItem, assertionErrors.ErrGeneric).
						Return().
						Once()

					err := s.Update(assertion.Ctx, idString, inputItem)
					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrGeneric))
				})
			})
			When("Fails to parse UUID from string", func() {
				It("Should return an error", func() {
					idString := assertion.InvalidIDString
					inputItem := assertion.NewItemWithID(idString)
					logMock.On(
						"Error",
						mock.Anything,
						assertion.NewErrIncorrectIDLength(assertion.InvalidIDString),
					).Once()

					err := s.Update(assertion.Ctx, assertion.InvalidIDString, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrCreatingUUID))
				})
			})
		})

		Context("Deleting an item", func() {
			When("Request succeeds", func() {
				It("Should return nothing", func() {
					repoMock.On("Remove", assertion.Ctx, assertion.SampleID).
						Return(nil).
						Once()

					err := s.Delete(assertion.Ctx, assertion.SampleID.String())
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					repoMock.On("Remove", assertion.Ctx, assertion.SampleID).
						Return(assertionErrors.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, assertion.SampleID, assertionErrors.ErrGeneric).
						Return().
						Once()

					err := s.Delete(assertion.Ctx, assertion.SampleID.String())
					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrGeneric))
				})
			})
			When("Fails to parse UUID from string", func() {
				It("Should return an error", func() {
					logMock.On(
						"Error",
						mock.Anything,
						assertion.NewErrIncorrectIDLength(assertion.InvalidIDString),
					).Once()

					err := s.Delete(assertion.Ctx, assertion.InvalidIDString)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErrors.ErrCreatingUUID))
				})
			})
		})
	})
})
