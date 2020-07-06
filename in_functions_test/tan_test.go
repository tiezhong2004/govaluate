package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
	"math"
)

var _ = Describe("tan functions", func() {
	Describe("tan", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Tan(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tan(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Tan(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tan(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Tan("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Tan()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Tan(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("atan", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Atan(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Atan(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Atan("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Atan()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Atan(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("atan2", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Atan2(0.1363122, 0.7362629)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan2(0.1363122, 0.7362629)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Atan2(3, 1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan2(3, 1)))
		})

		It("should return an param wrong type error (first param)", func() {
			_, err := in_functions.Atan2("invalid param value", 0.73)
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should return an param wrong type error (second param)", func() {
			_, err := in_functions.Atan2(0.13, "invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Atan2()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too few arguments 2)", func() {
			c, err := in_functions.Atan2(1)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Atan2(1, 2, 3)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("tanh", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Tanh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tanh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Tanh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tanh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Tanh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Tanh()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Tanh(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("atanh", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Atanh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atanh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Atanh(0.12345)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atanh(0.12345)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Atanh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Atanh()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Atanh(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
