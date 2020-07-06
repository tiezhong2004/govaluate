package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
)

var _ = Describe("max, min functions", func() {
	Describe("max", func() {
		It("should resolve a dual max", func() {
			c, err := in_functions.Max(1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(2.0))
		})

		It("should resolve a the max value for 5 values", func() {
			c, err := in_functions.Max(4, 5, 3, 1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(5.0))
		})

		It("should fail with a param type error (string supplied)", func() {
			c, err := in_functions.Max(4, 5, 3, 1, 2, "non number")
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param type error (boolean supplied)", func() {
			c, err := in_functions.Max(4, 5, 3, 1, 2, false)
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param count error", func() {
			c, err := in_functions.Max()
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		})
	})

	Describe("min", func() {
		It("should resolve a dual min", func() {
			c, err := in_functions.Min(2, 1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should resolve a the min value for 5 values", func() {
			c, err := in_functions.Min(4, 5, 3, 1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should fail with a param type error (string supplied)", func() {
			c, err := in_functions.Min(4, 5, 3, 1, 2, "non number")
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param type error (boolean supplied)", func() {
			c, err := in_functions.Min(4, 5, 3, 1, 2, false)
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param count error", func() {
			c, err := in_functions.Min()
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
		})
	})
})
