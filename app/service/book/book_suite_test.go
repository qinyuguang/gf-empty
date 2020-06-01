package book_test

import (
	"testing"

	_ "gf-empty/boot/testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Book Suite")
}
