package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stuart-pollock/go-raytracer/math"
)

var _ = Describe("Vector", func() {
	Context("GetVector", func() {
		Describe("Return a zero-Vector (0, 0, 0)", func() {
			foo := GetVector()

			Expect("Returned Vector should not be nil", foo).NotTo(BeNil())
		})
	})
})
