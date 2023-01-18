package http_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	httpService "app/internal/http"
	assertionErrors "app/internal/test/assertion/errors"
)

func TestError(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suits")
}

var _ = Describe("Error", func() {
	var (
		httpError httpService.Error
	)

	BeforeEach(func() {
		httpError = httpService.NewHttpError()
	})

	Context("Getting status code from error", func() {
		When("A record is not found", func() {
			It("Should return status not found", func() {
				status := httpError.GetStatus(gorm.ErrRecordNotFound)

				Expect(status).To(Equal(http.StatusNotFound))
			})
		})
		When("User sent request with missing required value", func() {
			It("Should return status bad request", func() {
				status := httpError.GetStatus(gorm.ErrPrimaryKeyRequired)

				Expect(status).To(Equal(http.StatusBadRequest))
			})
		})
		When("Error is not mapped", func() {
			It("Should return status internal server error", func() {
				status := httpError.GetStatus(assertionErrors.ErrGeneric)

				Expect(status).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
