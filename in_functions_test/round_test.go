package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
)

var _ = Describe("round functions", func() {
	Describe("round", func() {
		It("should round to zero", func() {
			c, err := in_functions.Round(0.49999999)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should round to one", func() {
			c, err := in_functions.Round(0.5)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should round to zero", func() {
			c, err := in_functions.Round(-0.4999999)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should round to minus one", func() {
			c, err := in_functions.Round(-0.5)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(-1.0))
		})

		It("should round taking a decimal point in count", func() {
			c, err := in_functions.Round(1.2345678, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.23))
		})

		It("should round taking a decimal point in count (rounding up)", func() {
			c, err := in_functions.Round(1.2385674, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.24))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Round("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Round()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Round(1, 2, 3)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("trunc", func() {
		It("should trunc down the value", func() {
			c, err := in_functions.Trunc(0.49999999)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should round to one", func() {
			c, err := in_functions.Trunc(0.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should return a param type error", func() {
			_, err := in_functions.Trunc("no valid value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Trunc()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Trunc(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("ceil", func() {
		It("should ceil a low float to one", func() {
			c, err := in_functions.Ceil(0.009)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should ceil a high float to one", func() {
			c, err := in_functions.Ceil(0.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should ceil a negative value", func() {
			c, err := in_functions.Ceil(-1.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(-1.0))
		})

		It("should return a param type error", func() {
			_, err := in_functions.Ceil("no valid value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Ceil()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Ceil(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("floor", func() {
		It("should floor a low float to zero", func() {
			c, err := in_functions.Floor(0.009)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should floor a high float to one", func() {
			c, err := in_functions.Floor(0.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should floor a negative value", func() {
			c, err := in_functions.Floor(-1.4)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(-2.0))
		})

		It("should return a param type error", func() {
			_, err := in_functions.Floor("no valid value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Floor()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Floor(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
