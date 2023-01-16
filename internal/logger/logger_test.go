package logger

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	assertion "app/internal/test/assertion/pkg"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Suits")
}

var _ = Describe("Log", func() {
	Context("Generating application logs", func() {
		Context("When Debug is disabled", func() {
			logger := NewLogger(false)
			When("Logging Info type", func() {
				It("Should call info.Println", func() {
					logger.Info(context.Background(), assertion.InfoLogMessage, nil)
				})
			})
			When("Logging Warn type", func() {
				It("Should call warn.Println", func() {
					logger.Warn(context.Background(), assertion.WarnLogMessage, nil)
				})
			})
			When("Logging Error type", func() {
				It("Should call info.Println", func() {
					logger.Error(context.Background(), errors.New("err"), assertion.ErrLogMessage, nil)
				})
			})
			When("Logging Debug type", func() {
				It("Should not call debug.Println", func() {
					logger.Debug(context.Background(), assertion.DebugLogMessage, nil)
				})
			})
		})
		Context("When Debug is enabled", func() {
			logger := NewLogger(true)
			When("Logging Debug type", func() {
				It("Should call debug.Println", func() {
					logger.Debug(context.Background(), assertion.DebugLogMessage, nil)
				})
			})
		})
	})
	Context("Getting Log Path", func() {
		When("Requesting the log path", func() {
			It("Should return t", func() {
				path := getLogPath()

				dir := strings.Split(path, "/")
				rootDir := dir[len(dir)-2]
				logDir := dir[len(dir)-1]

				Expect(rootDir).To(BeEquivalentTo(assertion.RootDir))
				Expect(logDir).To(BeEquivalentTo(assertion.LogDir))
			})
		})
	})
})

var _ = Describe("Dir", func() {
	Context("Generating application logs", func() {
		When("When Debug is disabled", func() {
			It("Should not call debug.Println", func() {
				dir := getProjectRootDirectory()
				folders := strings.Split(dir, "/")
				rootDir := folders[len(folders)-1]
				Expect(rootDir).To(BeEquivalentTo(assertion.RootDir))
			})
		})
	})
})

var _ = AfterSuite(func() {
	logFile := fmt.Sprintf("%s/logs/%s", getProjectRootDirectory(), assertion.LogFileName)
	err := os.Remove(logFile)
	Expect(err).NotTo(HaveOccurred())
})
