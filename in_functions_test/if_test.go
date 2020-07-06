package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
)

var _ = Describe("if functions", func() {
	It("should evaluate a boolean operation (true)", func() {
		c, err := in_functions.If(true, 1, 0)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(1))
	})

	It("should evaluate a boolean operation (false)", func() {
		c, err := in_functions.If(false, 1, 0)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(0))
	})

	It("should get correct value with a float (non-zero)", func() {
		c, err := in_functions.If(0.000001, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())

		c, err = in_functions.If(-0.000001, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should get correct value with a float (zero)", func() {
		c, err := in_functions.If(0.0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())

		c, err = in_functions.If(-0.0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should get correct value with an integer (non-zero)", func() {
		c, err := in_functions.If(1, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())

		c, err = in_functions.If(-1, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should get correct value with an integer (zero)", func() {
		c, err := in_functions.If(0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())

		c, err = in_functions.If(-0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should get correct value with a string (non-empty)", func() {
		c, err := in_functions.If("non empty string", true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should get correct value with a string (empty)", func() {
		c, err := in_functions.If("", true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should get an wrong param count error (too few arguments)", func() {
		c, err := in_functions.If()
		Expect(c).To(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())

		c, err = in_functions.If(0)
		Expect(c).To(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())

		c, err = in_functions.If(0, 0)
		Expect(c).To(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())
	})

	It("should get an wrong param count error (too many arguments)", func() {
		c, err := in_functions.Exp(true, true, true, true)
		Expect(c).To(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())
	})
})
