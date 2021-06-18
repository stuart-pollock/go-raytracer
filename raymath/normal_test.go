package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/go-raytracer/raymath"
)

var _ = Describe("Normal", func() {
	Context("GetNormal", func() {
		It("returns a Normal with non-nil co-ordinates", func() {
			result := GetNormal(0, 0, 0)

			Expect(result.XYZ).NotTo(BeNil())
		})

		It("returns zero-length Normal, initialized to (-12.2, 54, 0.001)", func() {
			result := GetNormal(-12.2, 54, 0.001)

			expected := Normal{}
			expected.XYZ = make([]float64, 4)
			expected.XYZ[0] = -12.2
			expected.XYZ[1] = 54
			expected.XYZ[2] = 0.001
			expected.XYZ[3] = 0

			Expect(result).To(Equal(expected))
		})
	})

	Context("MatMult", func() {
		It("postmultiplies a Normal by the identity matrix", func() {
			result := GetNormal(2, 4, 3)
			result.MatMult(GetIdentityMatrix())

			expected := GetNormal(2, 4, 3)

			Expect(result).To(Equal(expected))
		})

		It("postmultiplies the Normal (2, 4, 3) by the a non-identity matrix", func() {
			result := GetNormal(2, 4, 3)
			result.MatMult(GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12)))

			expected := GetNormal(
				2*1+4*2+3*3,
				2*4+4*5+3*6,
				2*7+4*8+3*9,
			)

			Expect(result).To(Equal(expected))
		})
	})
})
