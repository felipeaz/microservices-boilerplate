package log

import (
	"fmt"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	assertion "microservices-boilerplate/internal/test/assertion/pkg"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Suits")
}

var _ = Describe("Log", func() {
	logFile := fmt.Sprintf("%s/%s", GetLogPath(), assertion.LogFile)
	Context("Generating application logs", func() {
		Context("When Debug is disabled", func() {
			logger := NewLogger(assertion.LogTime, false)
			When("Logging Info type", func() {
				It("Should call info.Println", func() {
					logger.Info(assertion.InfoLogMessage)

					Expect(logFile).To(BeAnExistingFile())
				})
			})
			When("Logging Warn type", func() {
				It("Should call warn.Println", func() {
					logger.Warn(assertion.WarnLogMessage)

					Expect(logFile).To(BeAnExistingFile())
				})
			})
			When("Logging Error type", func() {
				It("Should call info.Println", func() {
					logger.Error(assertion.ErrLogMessage)

					Expect(logFile).To(BeAnExistingFile())
				})
			})
			When("Logging Debug type", func() {
				It("Should not call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)

					Expect(logFile).To(BeAnExistingFile())
				})
			})
		})
		Context("When Debug is enabled", func() {
			logger := NewLogger(assertion.LogTime, true)
			When("Logging Debug type", func() {
				It("Should call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)

					Expect(logFile).To(BeAnExistingFile())
				})
			})
		})
		Context("Getting Log Path", func() {
			When("Requesting the log path", func() {
				It("Should return t", func() {
					path := GetLogPath()

					dir := strings.Split(path, "/")
					rootDir := dir[len(dir)-2]
					logDir := dir[len(dir)-1]

					Expect(rootDir).To(BeEquivalentTo(assertion.RootDir))
					Expect(logDir).To(BeEquivalentTo(assertion.LogDir))
				})
			})
		})
	})
})

var _ = AfterSuite(func() {
	logFile := fmt.Sprintf("%s/%s", GetLogPath(), assertion.LogFile)
	err := os.Remove(logFile)

	Expect(err).NotTo(HaveOccurred())
})
