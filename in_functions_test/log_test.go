package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
	"math"
)

var _ = Describe("log functions", func() {
	Describe("log", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Log(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Log(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Log("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Log()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Log(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("log10", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Log10(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log10(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Log10(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log10(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Log10("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Log10()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Log10(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
