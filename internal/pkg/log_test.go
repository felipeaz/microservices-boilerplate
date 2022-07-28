package pkg_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microservices-boilerplate/internal/pkg"
	assertion "microservices-boilerplate/internal/test/assertion/pkg"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suits")
}

var _ = Describe("Log", func() {
	Context("Generating application logs", func() {
		Context("When Debug is disabled", func() {
			logger := pkg.NewLogger(false)
			When("Logging Info type", func() {
				It("Should call info.Println", func() {
					logger.Info(assertion.InfoLogMessage)
				})
			})
			When("Logging Warn type", func() {
				It("Should call warn.Println", func() {
					logger.Warn(assertion.WarnLogMessage)
				})
			})
			When("Logging Error type", func() {
				It("Should call info.Println", func() {
					logger.Error(assertion.ErrLogMessage)
				})
			})
			When("Logging Debug type", func() {
				It("Should not call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)
				})
			})
		})
		Context("When Debug is enabled", func() {
			logger := pkg.NewLogger(true)
			When("Logging Debug type", func() {
				It("Should call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)
				})
			})
		})
	})
})
