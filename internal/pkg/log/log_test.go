package log

import (
	"fmt"
	"io"
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
	var (
		logFilePath string
		logFile     io.Writer
	)

	BeforeEach(func() {
		logFilePath = fmt.Sprintf("%s/%s", GetLogPath(), assertion.LogFileName)
	})

	Context("Generating application logs", func() {
		Context("When Debug is disabled", func() {
			logger := NewLogger(os.Stdout, false)
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
			logger := NewLogger(os.Stdout, true)
			When("Logging Debug type", func() {
				It("Should call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)
				})
			})
		})
		Context("When logging to a file", func() {
			logFile = NewLogFile(assertion.LogTime, GetLogPath())
			logger := NewLogger(logFile, false)
			When("Logging Info type", func() {
				It("Should call info.Println", func() {
					logger.Info(assertion.InfoLogMessage)
					Expect(logFile).NotTo(BeNil())
					Expect(logFilePath).To(BeAnExistingFile())
				})
			})
			When("Logging Warn type", func() {
				It("Should call warn.Println", func() {
					logger.Warn(assertion.WarnLogMessage)

					Expect(logFilePath).To(BeAnExistingFile())
				})
			})
			When("Logging Error type", func() {
				It("Should call info.Println", func() {
					logger.Error(assertion.ErrLogMessage)

					Expect(logFilePath).To(BeAnExistingFile())
				})
			})
			When("Logging Debug type", func() {
				It("Should not call debug.Println", func() {
					logger.Debug(assertion.DebugLogMessage)

					Expect(logFilePath).To(BeAnExistingFile())
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
		Context("Creating Log File", func() {
			When("Creating a new file", func() {
				It("Should return t", func() {
					f := NewLogFile(assertion.LogTime, GetLogPath())

					Expect(f).ToNot(BeNil())
					Expect(logFilePath).To(BeAnExistingFile())
				})
			})
		})
		Context("Fail to initialize Log Path", func() {
			When("Creating a log file", func() {
				It("Should log Fatal and exit", func() {
					f := NewLogFile(assertion.LogTime, "log.go")

					Expect(f).To(Equal(os.Stdout))
				})
			})
		})
	})
})

var _ = AfterSuite(func() {
	logFile := fmt.Sprintf("%s/%s", GetLogPath(), assertion.LogFileName)
	err := os.Remove(logFile)
	Expect(err).NotTo(HaveOccurred())
})
