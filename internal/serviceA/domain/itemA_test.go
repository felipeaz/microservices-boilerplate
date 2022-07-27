package domain_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"microservices-boilerplate/internal/serviceA/domain"
	assertion "microservices-boilerplate/internal/test/assertion/serviceA"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain")
}

var _ = Describe("Domain", func() {
	Context("Creating Instances", func() {
		When("Creating a single instance from Bytes", func() {
			It("Should return a single item", func() {
				itemID := assertion.SampleID.String()
				expect := assertion.NewItemWithID(itemID)

				item, err := domain.NewFromBytes(assertion.ItemAInBytes(expect))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(item).To(Equal(expect))
			})
			It("Should fail to unmarshal item", func() {
				item, err := domain.NewFromBytes(nil)
				Expect(err).Should(HaveOccurred())
				Expect(item).To(BeNil())
			})
		})

		When("Creating an Array of instances from Bytes", func() {
			It("Should return an array of item", func() {
				expect := assertion.ArrayOfItem

				itemArr, err := domain.NewArrayFromBytes(assertion.ArrayOfItemAInBytes(expect))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(itemArr).To(Equal(expect))
			})
			It("Should fail to unmarshal item", func() {
				itemArr, err := domain.NewArrayFromBytes(nil)
				Expect(err).Should(HaveOccurred())
				Expect(itemArr).To(BeNil())
			})
		})
	})
})
