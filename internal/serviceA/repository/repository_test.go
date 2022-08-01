package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/serviceA/repository"
	assertionErr "microservices-boilerplate/internal/test/assertion/errors"
	assertion "microservices-boilerplate/internal/test/assertion/serviceA"
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
		repo = repository.New(databaseMock, cacheMock)
	})

	Context("Testing CRUD operations", func() {
		Context("Getting all items", func() {
			When("Item is in cache", func() {
				When("Succeeds", func() {
					It("Should return an item from cache", func() {
						expectedItemArr := assertion.ArrayOfItem
						itemArrInBytes := assertion.ArrayOfItemAInBytes(expectedItemArr)
						cacheMock.On("Get", repository.AllItemsKey).
							Return(itemArrInBytes, nil).
							Once()

						item, err := repo.GetAll(assertion.Ctx)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(item).To(Equal(expectedItemArr))
					})
				})
				When("Fails", func() {
					It("Should return an error", func() {
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, assertionErr.ErrGeneric).
							Once()

						item, err := repo.GetAll(assertion.Ctx)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(assertionErr.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
			})
			When("Item is not in cache", func() {
				When("Succeeds", func() {
					It("Should return an item", func() {
						var emptyArr []*domain.ItemA
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", &emptyArr).
							Return(nil).
							Once()
						cacheMock.On("Set", repository.AllItemsKey, emptyArr).
							Return(nil).
							Once()

						_, err := repo.GetAll(assertion.Ctx)

						Expect(err).ShouldNot(HaveOccurred())
					})
				})
				When("Fails to get item from Database", func() {
					It("Should return an error", func() {
						var emptyArr []*domain.ItemA
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", &emptyArr).
							Return(assertionErr.ErrGeneric).
							Once()

						item, err := repo.GetAll(assertion.Ctx)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(assertionErr.ErrGeneric))
						Expect(item).To(BeNil())
					})
				})
				When("Fails to set cache", func() {
					It("Should return an error", func() {
						var emptyArr []*domain.ItemA
						cacheMock.On("Get", repository.AllItemsKey).
							Return(nil, nil).
							Once()
						databaseMock.On("Select", &emptyArr).
							Return(nil).
							Once()
						cacheMock.On("Set", repository.AllItemsKey, emptyArr).
							Return(assertionErr.ErrGeneric).
							Once()

						item, err := repo.GetAll(assertion.Ctx)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(assertionErr.ErrGeneric))
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
						itemInBytes := assertion.ItemAInBytes(expectedItem)
						cacheMock.On("Get", idString).
							Return(itemInBytes, nil).
							Once()

						item, err := repo.GetByID(assertion.Ctx, assertion.SampleID)

						Expect(err).ShouldNot(HaveOccurred())
						Expect(item).To(Equal(expectedItem))
					})
				})
				When("Fails", func() {
					It("Should return an error", func() {
						idString := assertion.SampleID.String()
						cacheMock.On("Get", idString).
							Return(nil, assertionErr.ErrGeneric).
							Once()

						item, err := repo.GetByID(assertion.Ctx, assertion.SampleID)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(assertionErr.ErrGeneric))
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
						databaseMock.On("Select", assertion.NewItemWithID(idString)).
							Return(nil).
							Once()
						cacheMock.On("Set", idString, expectedItem).
							Return(nil).
							Once()

						item, err := repo.GetByID(assertion.Ctx, assertion.SampleID)

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
						databaseMock.On("Select", assertion.NewItemWithID(idString)).
							Return(assertionErr.ErrGeneric).
							Once()

						item, err := repo.GetByID(assertion.Ctx, assertion.SampleID)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(assertionErr.ErrGeneric))
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
						databaseMock.On("Select", assertion.NewItemWithID(idString)).
							Return(nil).
							Once()
						cacheMock.On("Set", idString, expectedItem).
							Return(assertionErr.ErrGeneric).
							Once()

						item, err := repo.GetByID(assertion.Ctx, assertion.SampleID)

						Expect(err).Should(HaveOccurred())
						Expect(err).To(Equal(assertionErr.ErrGeneric))
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
					databaseMock.On("Create", inputItem).
						Return(nil).
						Once()

					item, err := repo.Insert(assertion.Ctx, inputItem)

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
					databaseMock.On("Create", inputItem).
						Return(assertionErr.ErrGeneric).
						Once()

					item, err := repo.Insert(assertion.Ctx, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
					Expect(item).To(BeNil())
				})
			})
			When("Fail to remove all cached items", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(assertionErr.ErrGeneric).
						Once()

					item, err := repo.Insert(assertion.Ctx, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
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
					databaseMock.On("Update", assertion.SampleID, inputItem).
						Return(nil).
						Once()

					err := repo.Update(assertion.Ctx, assertion.SampleID, inputItem)

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
					databaseMock.On("Update", assertion.SampleID, inputItem).
						Return(assertionErr.ErrGeneric).
						Once()

					err := repo.Update(assertion.Ctx, assertion.SampleID, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
				})
			})
			When("Fail to remove cached item", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(assertionErr.ErrGeneric).
						Once()

					err := repo.Update(assertion.Ctx, assertion.SampleID, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
				})
			})
			When("Fail to remove all cached items", func() {
				It("Should return an error", func() {
					inputItem := assertion.NewItemWithID(assertion.SampleID.String())
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(assertionErr.ErrGeneric).
						Once()

					err := repo.Update(assertion.Ctx, assertion.SampleID, inputItem)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
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
					databaseMock.On("Delete", assertion.SampleID, domain.ItemA{}).
						Return(nil).
						Once()

					err := repo.Remove(assertion.Ctx, assertion.SampleID)

					Expect(err).ShouldNot(HaveOccurred())
				})
			})
			When("Fail to remove cached item", func() {
				It("Should return an error", func() {
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(assertionErr.ErrGeneric).
						Once()

					err := repo.Remove(assertion.Ctx, assertion.SampleID)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
				})
			})
			When("Fail to remove all cached items", func() {
				It("Should return an error", func() {
					cacheMock.On("Remove", assertion.SampleID.String()).
						Return(nil).
						Once()
					cacheMock.On("Remove", repository.AllItemsKey).
						Return(assertionErr.ErrGeneric).
						Once()

					err := repo.Remove(assertion.Ctx, assertion.SampleID)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
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
					databaseMock.On("Delete", assertion.SampleID, domain.ItemA{}).
						Return(assertionErr.ErrGeneric).
						Once()

					err := repo.Remove(assertion.Ctx, assertion.SampleID)

					Expect(err).Should(HaveOccurred())
					Expect(err).To(Equal(assertionErr.ErrGeneric))
				})
			})
		})
	})
})
