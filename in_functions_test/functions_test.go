package in_functions_test

import (
	"github.com/jamillosantos/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"testing"
)

func TestFunctions(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "in_functions")
}
