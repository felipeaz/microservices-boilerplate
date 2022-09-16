package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	httpErr "microservices-boilerplate/internal/http"
	errorsAssertion "microservices-boilerplate/internal/test/assertion/errors"
	assertion "microservices-boilerplate/internal/test/assertion/serviceA"
	serviceMocks "microservices-boilerplate/internal/test/mocks/serviceA/service"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suits")
}

func ginCtxParam(key, value string) gin.Param {
	return gin.Param{
		Key:   key,
		Value: value,
	}
}

var _ = Describe("Handler", func() {
	var (
		router      *gin.Engine
		w           *httptest.ResponseRecorder
		ginCtx      *gin.Context
		serviceMock *serviceMocks.Service
		config      *Config
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		w = httptest.NewRecorder()
		ginCtx, router = gin.CreateTestContext(w)
		serviceMock = serviceMocks.NewService(GinkgoT())
		config = &Config{
			Service:   serviceMock,
			HttpError: httpErr.NewHttpError(),
			Router:    router,
		}
	})

	Context("CRUD Operations", func() {
		Context("GET", func() {
			When("Succeed", func() {
				It("Return an array of item from DB", func() {
					arrayOfItemAInBytes := assertion.ArrayOfItemAInBytes(assertion.ArrayOfItem)
					serviceMock.On("GetAll", ginCtx).
						Return(assertion.ArrayOfItem, nil)

					New(config)

					request, err := http.NewRequestWithContext(ginCtx, http.MethodGet, "/api/v1/a-items", nil)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					respInBytes, err := ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusOK))
					Expect(respInBytes).To(Equal(arrayOfItemAInBytes))
				})
			})
			When("Fails", func() {
				It("Return an Internal Server Error", func() {
					serviceMock.On("GetAll", ginCtx).
						Return(nil, errorsAssertion.ErrGeneric)

					New(config)

					request, err := http.NewRequestWithContext(ginCtx, http.MethodGet, "/api/v1/a-items", nil)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).NotTo(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusInternalServerError))
				})
			})
		})

		Context("FIND", func() {
			When("Succeed", func() {
				It("Return an item from DB", func() {
					itemID := assertion.SampleID.String()
					item := assertion.NewItemWithID(itemID)
					itemAInBytes := assertion.ItemAInBytes(item)
					serviceMock.On("GetOneByID", ginCtx, itemID).
						Return(item, nil)
					ginCtx.Params = []gin.Param{
						ginCtxParam("id", itemID),
					}

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodGet,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						nil,
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					respInBytes, err := ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusOK))
					Expect(respInBytes).To(Equal(itemAInBytes))
				})
			})
			When("Fails", func() {
				It("Return a Not Found error", func() {
					itemID := assertion.SampleID.String()
					serviceMock.On("GetOneByID", ginCtx, itemID).
						Return(nil, errorsAssertion.ErrNotFound)
					ginCtx.Params = []gin.Param{
						ginCtxParam("id", itemID),
					}

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodGet,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						nil,
					)

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).NotTo(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusNotFound))
				})
			})
		})

		Context("CREATE", func() {
			When("Succeed", func() {
				It("Creates a new item", func() {
					itemInput := assertion.NewItemWithoutID()
					inputInBytes := assertion.ItemAInBytes(itemInput)
					expectedOutput := *itemInput
					expectedOutput.ID = assertion.SampleID
					serviceMock.On("Create", ginCtx, itemInput).
						Return(&expectedOutput, nil)

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodPost,
						"/api/v1/a-items",
						bytes.NewBuffer(inputInBytes),
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					respInBytes, err := ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusOK))
					Expect(respInBytes).To(Equal(assertion.ItemAInBytes(&expectedOutput)))
				})
			})
			When("Fails", func() {
				It("Return Bad Request when fails to Bind Input JSON", func() {
					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodPost,
						"/api/v1/a-items",
						nil,
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusBadRequest))
				})
				It("Return Internal Server Error when fails to create item", func() {
					itemInput := assertion.NewItemWithoutID()
					inputInBytes := assertion.ItemAInBytes(itemInput)
					expectedOutput := *itemInput
					expectedOutput.ID = assertion.SampleID
					serviceMock.On("Create", ginCtx, itemInput).
						Return(nil, errorsAssertion.ErrCreatingUUID)

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodPost,
						"/api/v1/a-items",
						bytes.NewBuffer(inputInBytes),
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusInternalServerError))
				})
			})
		})

		Context("UPDATE", func() {
			When("Succeed", func() {
				It("Return no content", func() {
					itemID := assertion.SampleID.String()
					itemInput := assertion.NewItemWithoutID()
					inputInBytes := assertion.ItemAInBytes(itemInput)
					serviceMock.On("Update", ginCtx, itemID, itemInput).
						Return(nil)

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodPut,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						bytes.NewBuffer(inputInBytes),
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusNoContent))
				})
			})
			When("Fails", func() {
				It("Return a Bad Request Error when fails to Bind Input JSON", func() {
					itemID := assertion.SampleID.String()
					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodPut,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						nil,
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusBadRequest))
				})
				It("Return a Not Found Error when fails to Update item", func() {
					itemID := assertion.SampleID.String()
					itemInput := assertion.NewItemWithoutID()
					inputInBytes := assertion.ItemAInBytes(itemInput)
					serviceMock.On("Update", ginCtx, itemID, itemInput).
						Return(errorsAssertion.ErrNotFound)

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodPut,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						bytes.NewBuffer(inputInBytes),
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusNotFound))
				})
			})
		})

		Context("DELETE", func() {
			When("Succeed", func() {
				It("Return an item from DB", func() {
					itemID := assertion.SampleID.String()
					serviceMock.On("Delete", ginCtx, itemID).
						Return(nil)
					ginCtx.Params = []gin.Param{
						ginCtxParam("id", itemID),
					}

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodDelete,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						nil,
					)
					Expect(err).ToNot(HaveOccurred())

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).ToNot(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusNoContent))
				})
			})
			When("Fails", func() {
				It("Return a Not Found error", func() {
					itemID := assertion.SampleID.String()
					serviceMock.On("Delete", ginCtx, itemID).
						Return(errorsAssertion.ErrNotFound)
					ginCtx.Params = []gin.Param{
						ginCtxParam("id", itemID),
					}

					New(config)

					request, err := http.NewRequestWithContext(
						ginCtx,
						http.MethodDelete,
						fmt.Sprintf("/api/v1/a-items/%s", itemID),
						nil,
					)

					router.ServeHTTP(w, request)

					_, err = ioutil.ReadAll(w.Body)
					Expect(err).NotTo(HaveOccurred())

					Expect(w.Code).To(Equal(http.StatusNotFound))
				})
			})
		})
	})
})
