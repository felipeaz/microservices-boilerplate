package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"microservices-boilerplate/internal/serviceB/service"
	commonAssertion "microservices-boilerplate/internal/test/assertion/common"
	errorsAssertion "microservices-boilerplate/internal/test/assertion/errors"
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
		s = service.New(
			&service.Config{
				Log:        logMock,
				Repository: repoMock,
			},
		)
	})

	Context("Testing CRUD Operations", func() {
		Context("Getting All items", func() {
			When("Request succeeds", func() {
				It("Should return all items from DB", func() {
					expectedItems := assertion.ArrayOfItem
					repoMock.On("GetAll", commonAssertion.EmptyCtx).
						Return(expectedItems, nil).
						Once()

					resp, err := s.GetAll(commonAssertion.EmptyCtx)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItems))
				})
			})
			When("DB is empty", func() {
				It("Should an empty array", func() {
					repoMock.On("GetAll", commonAssertion.EmptyCtx).
						Return(nil, nil).
						Once()

					resp, err := s.GetAll(commonAssertion.EmptyCtx)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(BeNil())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					repoMock.On("GetAll", commonAssertion.EmptyCtx).
						Return(nil, errorsAssertion.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, errorsAssertion.ErrGeneric).
						Once()

					resp, err := s.GetAll(commonAssertion.EmptyCtx)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
					Expect(resp).To(BeNil())
				})
			})
		})

		Context("Getting one item by ID", func() {
			When("Request succeeds", func() {
				It("Should return an item with given ID", func() {
					idString := assertion.SampleID.String()
					expectedItem := assertion.NewItemWithID(idString)
					repoMock.On("GetByID", commonAssertion.EmptyCtx, assertion.SampleID).
						Return(expectedItem, nil).
						Once()

					resp, err := s.GetOneByID(commonAssertion.EmptyCtx, assertion.SampleID.String())

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItem))
				})
			})
			When("Item is not found", func() {
				It("Should return a not found error", func() {
					repoMock.On("GetByID", commonAssertion.EmptyCtx, assertion.SampleID).
						Return(nil, errorsAssertion.ErrNotFound).
						Once()
					logMock.On("Error", mock.Anything, assertion.SampleID, errorsAssertion.ErrNotFound).
						Once()

					resp, err := s.GetOneByID(commonAssertion.EmptyCtx, assertion.SampleID.String())

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrNotFound))
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

					resp, err := s.GetOneByID(commonAssertion.EmptyCtx, assertion.InvalidIDString)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrCreatingUUID))
					Expect(resp).To(BeNil())
				})
			})
		})

		Context("Creating an item", func() {
			When("Request succeeds", func() {
				It("Should return the created object", func() {
					itemInput := assertion.NewItemWithoutID()
					expectedItem := assertion.NewItemFromInput(itemInput)
					repoMock.On("Insert", commonAssertion.EmptyCtx, itemInput).
						Return(expectedItem, nil).
						Once()

					resp, err := s.Create(commonAssertion.EmptyCtx, itemInput)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp).To(Equal(expectedItem))
					Expect(resp.ID).NotTo(BeEmpty())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					itemInput := assertion.NewItemWithoutID()
					repoMock.On("Insert", commonAssertion.EmptyCtx, itemInput).
						Return(nil, errorsAssertion.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, itemInput, errorsAssertion.ErrGeneric).
						Return().
						Once()

					resp, err := s.Create(commonAssertion.EmptyCtx, itemInput)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
					Expect(resp).To(BeNil())
				})
			})
		})

		Context("Updating an item", func() {
			When("Request succeeds", func() {
				It("Should return nothing", func() {
					idString := assertion.SampleID.String()
					inputItem := assertion.NewItemWithID(idString)
					repoMock.On("Update", commonAssertion.EmptyCtx, assertion.SampleID, inputItem).
						Return(nil).
						Once()

					err := s.Update(commonAssertion.EmptyCtx, idString, inputItem)
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					idString := assertion.SampleID.String()
					inputItem := assertion.NewItemWithID(idString)
					repoMock.On("Update", commonAssertion.EmptyCtx, assertion.SampleID, inputItem).
						Return(errorsAssertion.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, assertion.SampleID, inputItem, errorsAssertion.ErrGeneric).
						Return().
						Once()

					err := s.Update(commonAssertion.EmptyCtx, idString, inputItem)
					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
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

					err := s.Update(commonAssertion.EmptyCtx, assertion.InvalidIDString, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrCreatingUUID))
				})
			})
		})

		Context("Deleting an item", func() {
			When("Request succeeds", func() {
				It("Should return nothing", func() {
					repoMock.On("Remove", commonAssertion.EmptyCtx, assertion.SampleID).
						Return(nil).
						Once()

					err := s.Delete(commonAssertion.EmptyCtx, assertion.SampleID.String())
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Request fails", func() {
				It("Should return an error", func() {
					repoMock.On("Remove", commonAssertion.EmptyCtx, assertion.SampleID).
						Return(errorsAssertion.ErrGeneric).
						Once()
					logMock.On("Error", mock.Anything, assertion.SampleID, errorsAssertion.ErrGeneric).
						Return().
						Once()

					err := s.Delete(commonAssertion.EmptyCtx, assertion.SampleID.String())
					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
			When("Fails to parse UUID from string", func() {
				It("Should return an error", func() {
					logMock.On(
						"Error",
						mock.Anything,
						assertion.NewErrIncorrectIDLength(assertion.InvalidIDString),
					).Once()

					err := s.Delete(commonAssertion.EmptyCtx, assertion.InvalidIDString)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrCreatingUUID))
				})
			})
		})
	})
})
