package dir

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	assertion "microservices-boilerplate/internal/test/assertion/pkg"
)

func TestDir(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dir Suits")
}

var _ = Describe("Dir", func() {
	Context("Generating application logs", func() {
		When("When Debug is disabled", func() {
			It("Should not call debug.Println", func() {
				dir := GetProjectRootDirectory()
				folders := strings.Split(dir, "/")
				rootDir := folders[len(folders)-1]
				Expect(rootDir).To(BeEquivalentTo(assertion.RootDir))
			})
		})
	})
})
