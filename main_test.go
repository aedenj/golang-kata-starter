package main_test

import (
	. "github.com/aedenj/golang-kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scenario", func() {
	BeforeEach(func() {

	})

	It("true is true", func() {
		Expect(Kata()).To(Equal(true))
	})
})
