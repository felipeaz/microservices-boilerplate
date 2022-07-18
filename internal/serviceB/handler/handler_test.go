package handler

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServiceB Suits")
}

var _ = Describe("Testing ServiceB Handler", func() {
	Context("CRUD Operations", func() {
		Context("When user looks for all items", func() {
			It("Should return all items from DB", func() {})
		})

		Context("When user looks for an specific item", func() {
			It("Should return an item with given ID", func() {})
		})

		Context("When user creates an item", func() {
			It("Should return the created object", func() {})
		})

		Context("When user updates an item", func() {
			It("Should return nothing", func() {})
		})

		Context("When user deletes an item", func() {
			It("Should return nothing", func() {})
		})
	})
})
