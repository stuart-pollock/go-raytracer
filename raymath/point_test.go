package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stuart-pollock/go-raytracer/raymath"
)

var _ = Describe("Point", func() {
	Context("GetPoint", func() {
		Describe("Return a Point object which is at the origin (0, 0, 0)", func() {
			foo := GetPoint()

			Expect("Returned Point should not be nil", foo).NotTo(BeNil())
		})
	})
})
