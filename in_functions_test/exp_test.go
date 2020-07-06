package in_functions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tiezhong2004/govaluate/in_functions"
	"math"
)

var _ = Describe("exp functions", func() {
	Describe("exp", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Exp(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Exp(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Exp("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Exp()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Exp(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("exp2", func() {
		It("should get correct value with a float param", func() {
			c, err := in_functions.Exp2(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp2(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := in_functions.Exp2(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp2(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := in_functions.Exp2("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(in_functions.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := in_functions.Exp2()
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := in_functions.Exp2(1, 2)
			Expect(c).To(BeNil())
			Expect(in_functions.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
