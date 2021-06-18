package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/go-raytracer/raymath"
)

var _ = Describe("Vector", func() {
	Context("GetVector", func() {
		It("returns a Vector with four non-nil co-ordinates", func() {
			result := GetVector(-12.2, 54, 0.001)

			Expect(result.XYZ).NotTo(BeNil())
			Expect(len(result.XYZ)).To(Equal(4))
		})

		It("returns a Vector initialized to (-12.2, 54, 0.001)", func() {
			result := GetVector(-12.2, 54, 0.001)

			expected := Vector{}
			expected.XYZ = make([]float64, 4)
			expected.XYZ[0] = -12.2
			expected.XYZ[1] = 54
			expected.XYZ[2] = 0.001
			expected.XYZ[3] = 0

			Expect(result).To(Equal(expected))
		})
	})

	Context("Add", func() {
		It("adds one Vector to another", func() {
			result := GetVector(2, 4, 3)
			result.Add(GetVector(5.321, -3333.221, 92))

			expected := GetVector(7.321, -3331.221, 95)

			Expect(result).To(Equal(expected))
		})
	})

	Context("Cross", func() {
		It("calculates the cross product of one Vector with another", func() {
			vec := GetVector(2, 4, 3)
			result := vec.Cross(GetVector(1, -3, -6))

			expected := GetNormal(
				(4*-6)-(-3*3),
				(3*1)-(-6*2),
				(2*-3)-(1*4),
			)

			Expect(result).To(Equal(expected))
		})
	})

	Context("Dot", func() {
		It("calculates the dot product of one Vector with another", func() {
			vec := GetVector(2, 4, 3)
			result := vec.Dot(GetVector(14, -3, 5))

			expected := 2*14 + 4*-3 + 3*5

			Expect(result).To(Equal(expected))
		})
	})

	Context("Len", func() {
		It("calculates the length of a Vector within EPSILON of the alebraic calculation", func() {
			vec := GetVector(2, 1, -2)
			result := vec.Len()

			expected := 3

			Expect(result).Should(BeNumerically("~", expected, EPSILON))
		})
	})

	Context("Mult", func() {
		It("multiplies a Vector by a scalar", func() {
			result := GetVector(-2, 4.5, 33.45)
			result.Mult(4.2)

			expected := GetVector(-8.4, 18.9, 140.49)

			Expect(result).To(Equal(expected))
		})
	})

	Context("MatMult", func() {
		It("premultiplies a Vector by the identity matrix", func() {
			result := GetVector(2, 4, 3)
			result.MatMult(GetIdentityMatrix())

			expected := GetVector(2, 4, 3)

			Expect(result).To(Equal(expected))
		})

		It("premultiplies the Vector (2, 4, 3) by the a non-identity matrix", func() {
			result := GetVector(2, 4, 3)
			result.MatMult(GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12)))

			expected := GetVector(
				2*1+4*2+3*3,
				2*4+4*5+3*6,
				2*7+4*8+3*9,
			)

			Expect(result).To(Equal(expected))
		})
	})
})
