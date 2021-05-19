package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stuart-pollock/go-raytracer/raymath"
)

var _ = Describe("Matrix", func() {
	Context("GetMatrix", func() {
		Describe("Return an identity Matrix", func() {
			foo := GetPoint()

			Expect("Returned Matrix should not be nil", foo).NotTo(BeNil())
		})
	})
})
