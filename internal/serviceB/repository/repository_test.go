package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/internal/serviceB/repository"
	commonAssertion "microservices-boilerplate/internal/test/assertion/common"
	errorsAssertion "microservices-boilerplate/internal/test/assertion/errors"
	assertion "microservices-boilerplate/internal/test/assertion/serviceB"
	storageMock "microservices-boilerplate/internal/test/mocks/storage"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suits")
}

var _ = Describe("Repository", func() {
	var (
		cacheMock    *storageMock.Cache
		databaseMock *storageMock.Database
		repo         repository.Repository
	)

	BeforeEach(func() {
		cacheMock = storageMock.NewCache(GinkgoT())
		databaseMock = storageMock.NewDatabase(GinkgoT())
		repo = repository.New(
			&repository.Config{
				Database: databaseMock,
				Cache:    cacheMock,
			},
		)
	})

	Context("Testing CRUD operations", func() {
		Context("Getting all items", func() {
			When("Item is in cache", func() {
				When("Succeeds", func() {
					It("Should return an item from cache", func() {
						expectedItemArr := assertion.ArrayOfItem
						itemBrrInBytes := assertion.ArrayOfItemBInBytes(expectedItemArr)
						cacheMock.On("Get", repository.AllItemsKey).
							Return(itemBrrInBytes, nil).
							Once()

						item, err := repo.GetAll(commonAssertion.EmptyCtx)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(item).To(Equal(expectedItemArr))
					})
				})
				When("Fails", func() {
					It("Should return an error", func() {
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, errorsAssertion.ErrGeneric).
							Once()

						item, err := repo.GetAll(commonAssertion.EmptyCtx)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(errorsAssertion.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
			})
			When("Item is not in cache", func() {
				When("Succeeds", func() {
					It("Should return an item", func() {
						var emptyArr []*domain.ItemB
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", commonAssertion.EmptyCtx, &emptyArr).
							Return(nil).
							Once()
						cacheMock.On("Set", repository.AllItemsKey, emptyArr).
							Return(nil).
							Once()

						_, err := repo.GetAll(commonAssertion.EmptyCtx)

						Expect(err).ShouldNot(HaveOccurred())
					})
				})
				When("Fails to get item from Database", func() {
					It("Should return an error", func() {
						var emptyArr []*domain.ItemB
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", commonAssertion.EmptyCtx, &emptyArr).
							Return(errorsAssertion.ErrGeneric).
							Once()

						item, err := repo.GetAll(commonAssertion.EmptyCtx)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(errorsAssertion.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
				When("Fails to set cache", func() {
					It("Should return an error", func() {
						var emptyArr []*domain.ItemB
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", commonAssertion.EmptyCtx, &emptyArr).
							Return(nil).
							Once()
						cacheMock.On("Set", repository.AllItemsKey, emptyArr).
							Return(errorsAssertion.ErrGeneric).
							Once()

						item, err := repo.GetAll(commonAssertion.EmptyCtx)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(errorsAssertion.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
			})
		})

		Context("Getting one item", func() {
			When("Item is in cache", func() {
				When("Succeeds", func() {
					It("Should return an item from cache", func() {
						idString := assertion.SampleID.String()
						expectedItem := assertion.NewItemWithID(idString)
						itemInBytes := assertion.ItemBInBytes(expectedItem)
						cacheMock.On("Get", idString).
							Return(itemInBytes, nil).
							Once()

						item, err := repo.GetByID(commonAssertion.EmptyCtx, assertion.SampleID)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(item).To(Equal(expectedItem))
					})
				})
				When("Fails", func() {
					It("Should return an error", func() {
						idString := assertion.SampleID.String()
						cacheMock.On("Get", idString).
							Return(nil, errorsAssertion.ErrGeneric).
							Once()

						item, err := repo.GetByID(commonAssertion.EmptyCtx, assertion.SampleID)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(errorsAssertion.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
			})
			When("Item is not in cache", func() {
				When("Succeeds", func() {
					It("Should return an item", func() {
						idString := assertion.SampleID.String()
						expectedItem := assertion.NewItemWithID(idString)
						cacheMock.On("Get", idString).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", commonAssertion.EmptyCtx, assertion.NewItemWithID(idString)).
							Return(nil).
							Once()
						cacheMock.On("Set", idString, expectedItem).
							Return(nil).
							Once()

						item, err := repo.GetByID(commonAssertion.EmptyCtx, assertion.SampleID)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(item).To(Equal(expectedItem))
					})
				})
				When("Fails to get item from Database", func() {
					It("Should return an error", func() {
						idString := assertion.SampleID.String()
						cacheMock.On("Get", idString).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", commonAssertion.EmptyCtx, assertion.NewItemWithID(idString)).
							Return(errorsAssertion.ErrGeneric).
							Once()

						item, err := repo.GetByID(commonAssertion.EmptyCtx, assertion.SampleID)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(errorsAssertion.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
				When("Fails to set cache", func() {
					It("Should return an error", func() {
						idString := assertion.SampleID.String()
						expectedItem := assertion.NewItemWithID(idString)
						cacheMock.On("Get", idString).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", commonAssertion.EmptyCtx, assertion.NewItemWithID(idString)).
							Return(nil).
							Once()
						cacheMock.On("Set", idString, expectedItem).
							Return(errorsAssertion.ErrGeneric).
							Once()

						item, err := repo.GetByID(commonAssertion.EmptyCtx, assertion.SampleID)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(errorsAssertion.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
			})
		})

		Context("Creating an item", func() {
			When("Succeeds", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					expectedItem := assertion.NewItemFromInput(inputItem)
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(nil).
						Once()
					databaseMock.On("Create", commonAssertion.EmptyCtx, inputItem).
						Return(nil).
						Once()

					item, err := repo.Insert(commonAssertion.EmptyCtx, inputItem)

					Expect(err).ShouldNot(HaveOccurred())
					Expect(item.ID).NotTo(BeNil())
					Expect(item).To(Equal(expectedItem))
				})
			})
			When("Fails to insert item", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(nil).
						Once()
					databaseMock.On("Create", commonAssertion.EmptyCtx, inputItem).
						Return(errorsAssertion.ErrGeneric).
						Once()

					item, err := repo.Insert(commonAssertion.EmptyCtx, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
					Expect(item).To(BeNil())
				})
			})
			When("Fail to remove all cached items", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(errorsAssertion.ErrGeneric).
						Once()

					item, err := repo.Insert(commonAssertion.EmptyCtx, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
					Expect(item).To(BeNil())
				})
			})
		})

		Context("Updating an item", func() {
			When("Succeeds", func() {
				It("Should return nothing", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(nil).
						Once()
					databaseMock.On("Update", commonAssertion.EmptyCtx, assertion.SampleID, inputItem).
						Return(nil).
						Once()

					err := repo.Update(commonAssertion.EmptyCtx, assertion.SampleID, inputItem)

					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Fail to update item on DB", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(nil).
						Once()
					databaseMock.On("Update", commonAssertion.EmptyCtx, assertion.SampleID, inputItem).
						Return(errorsAssertion.ErrGeneric).
						Once()

					err := repo.Update(commonAssertion.EmptyCtx, assertion.SampleID, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
			When("Fail to remove cached item", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(errorsAssertion.ErrGeneric).
						Once()

					err := repo.Update(commonAssertion.EmptyCtx, assertion.SampleID, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
			When("Fail to remove all cached items", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(errorsAssertion.ErrGeneric).
						Once()

					err := repo.Update(commonAssertion.EmptyCtx, assertion.SampleID, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
		})

		Context("Deleting an item", func() {
			When("Succeeds", func() {
				It("Should return nothing", func() {
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(nil).
						Once()
					databaseMock.On("Delete", commonAssertion.EmptyCtx, assertion.SampleID, domain.ItemB{}).
						Return(nil).
						Once()

					err := repo.Remove(commonAssertion.EmptyCtx, assertion.SampleID)

					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Fail to remove cached item", func() {
				It("Should return an error", func() {
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(errorsAssertion.ErrGeneric).
						Once()

					err := repo.Remove(commonAssertion.EmptyCtx, assertion.SampleID)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
			When("Fail to remove all cached items", func() {
				It("Should return an error", func() {
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(errorsAssertion.ErrGeneric).
						Once()

					err := repo.Remove(commonAssertion.EmptyCtx, assertion.SampleID)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
			When("Fail to delete item from DB", func() {
				It("Should return an error", func() {
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(nil).
						Once()
					databaseMock.On("Delete", commonAssertion.EmptyCtx, assertion.SampleID, domain.ItemB{}).
						Return(errorsAssertion.ErrGeneric).
						Once()

					err := repo.Remove(commonAssertion.EmptyCtx, assertion.SampleID)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(errorsAssertion.ErrGeneric))
				})
			})
		})
	})
})
