package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suits")
}

var _ = Describe("Repository", func() {
	Context("Getting all items", func() {
		Context("Items are in cache", func() {

		})

		Context("Items are not in cache", func() {

		})
	})

	Context("Getting one item", func() {
		Context("Item is in cache", func() {

		})

		Context("Item is not in cache", func() {

		})
	})

	Context("Creating an item", func() {

	})

	Context("Updating an item", func() {

	})

	Context("Deleting an item", func() {

	})
})
