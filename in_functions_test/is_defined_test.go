package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
)

var _ = Describe("is_defined function", func() {
	It("should return that a nil value is not defined", func() {
		c, err := in_functions.IsDefined(nil)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non empty string is defined", func() {
		c, err := in_functions.IsDefined("this is a non empty string")
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non empty string is defined", func() {
		c, err := in_functions.IsDefined("")
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a zero valued integer is not defined", func() {
		c, err := in_functions.IsDefined(0)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non zero valued integer is defined (positive)", func() {
		c, err := in_functions.IsDefined(1)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non zero valued integer is defined (negative)", func() {
		c, err := in_functions.IsDefined(-1)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a zero valued float is not defined", func() {
		c, err := in_functions.IsDefined(0.0)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non zero valued float is defined (positive)", func() {
		c, err := in_functions.IsDefined(0.000001)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non zero valued float is defined (negative)", func() {
		c, err := in_functions.IsDefined(-0.000001)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return an param wrong type error", func() {
		_, err := in_functions.IsDefined(map[string]interface{}{
			"str": "value",
		})
		Expect(err).NotTo(BeNil())
		Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
	})

	It("should return a param count error (too few arguments)", func() {
		_, err := in_functions.IsDefined()
		Expect(err).NotTo(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
	})

	It("should return a param count error (too many arguments)", func() {
		_, err := in_functions.IsDefined(1, 2)
		Expect(err).NotTo(BeNil())
		Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
	})
})
