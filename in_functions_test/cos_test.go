package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate"
	"math"
)

var _ = Describe("cos functions", func() {
	Describe("cos", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Cos(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cos(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Cos(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cos(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Cos("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Cos()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Cos(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("acos", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Acos(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acos(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Acos(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acos(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Acos("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Acos()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Acos(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("cosh", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Cosh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cosh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Cosh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cosh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Cosh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Cosh()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Cosh(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("acosh", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Acosh(3.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acosh(3.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Acosh(6)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acosh(6)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Acosh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Acosh()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Acosh(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
