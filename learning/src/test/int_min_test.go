package test_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/leozhao0709/learning/src/test"
)

var _ = Describe("IntMin", func() {

	BeforeEach(func() {

	})

	It("should return correct value", func() {
		Expect(IntMin(2, -2)).To(Equal(-2))
	})
})
