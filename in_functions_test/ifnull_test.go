package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
)

var _ = Describe("ifnull function", func() {
	It("should return the first param", func() {
		c, err := in_functions.IfNull(1, 2)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(1))
	})

	It("should return the second param", func() {
		c, err := in_functions.IfNull(nil, 2)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(2))
	})

	It("should return the thrid param", func() {
		c, err := in_functions.IfNull(nil, nil, 3)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(3))
	})

	It("should return nil", func() {
		c, err := in_functions.IfNull(nil, nil, nil, nil)
		Expect(err).To(BeNil())
		Expect(c).To(BeNil())
	})

	It("should return a param count error (too few params)", func() {
		c, err := in_functions.IfNull()
		Expect(err).NotTo(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		Expect(c).To(BeNil())
	})
})
