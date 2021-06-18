package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRaymath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Raymath Suite")
}
