package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
	"math"
)

var _ = Describe("sin functions", func() {
	Describe("sin", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Sin(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sin(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Sin(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sin(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Sin("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Sin()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Sin(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("asin", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Asin(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asin(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Asin(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asin(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Asin("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Asin()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Asin(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("sinh", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Sinh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sinh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Sinh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sinh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Sinh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Sinh()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Sinh(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("asinh", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Asinh(3.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asinh(3.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Asinh(6)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asinh(6)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Asinh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Asinh()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Asinh(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
