package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/go-raytracer/raymath"
)

var _ = Describe("Point", func() {
	Context("GetPoint", func() {
		It("returns a Point with four non-nil co-ordinates", func() {
			result := GetPoint(-12.2, 54, 0.001)

			Expect(result.XYZ).NotTo(BeNil())
			Expect(len(result.XYZ)).To(Equal(4))
		})

		It("returns a Point at (-12.2, 54, 0.001)", func() {
			result := GetPoint(-12.2, 54, 0.001)

			expected := Point{}
			expected.XYZ = make([]float64, 4)
			expected.XYZ[0] = -12.2
			expected.XYZ[1] = 54
			expected.XYZ[2] = 0.001
			expected.XYZ[3] = 1

			Expect(result).To(Equal(expected))
		})
	})

	Context("Add", func() {
		It("adds the Vector (5.321, -3333.221, 92) to the Point (2, 4, 3)", func() {
			result := GetPoint(2, 4, 3)
			result.Add(GetVector(5.321, -3333.221, 92))

			expected := GetPoint(7.321, -3329.221, 95)

			for i := 0; i < len(expected.XYZ); i++ {
				Expect(result.XYZ[i]).Should(BeNumerically("~", expected.XYZ[i], EPSILON))
			}
		})
	})

	Context("Subtract", func() {
		It("subtracts the Point (5.321, -3333.221, 92) from the Point (2, 4, 3)", func() {
			p := GetPoint(2, 4, 3)
			result := p.Subtract(GetPoint(5.321, -3333.221, 92))

			expected := GetVector(-3.321, 3337.221, -89)

			for i := 0; i < len(expected.XYZ); i++ {
				Expect(result.XYZ[i]).Should(BeNumerically("~", expected.XYZ[i], EPSILON))
			}
		})
	})

	Context("MatMult", func() {
		It("premultiplies the Point (2, 4, 3) by the identity matrix", func() {
			result := GetPoint(2, 4, 3)
			result.MatMult(GetIdentityMatrix())

			expected := GetPoint(2, 4, 3)

			Expect(result).To(Equal(expected))
		})

		It("premultiplies the Point (2, 4, 3) by the a non-identity matrix", func() {
			result := GetPoint(2, 4, 3)
			result.MatMult(GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12)))

			expected := GetPoint(
				2*1+4*2+3*3+10,
				2*4+4*5+3*6+11,
				2*7+4*8+3*9+12,
			)

			Expect(result).To(Equal(expected))
		})
	})
})
