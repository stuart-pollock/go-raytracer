package raymath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/go-raytracer/raymath"
)

var _ = Describe("Matrix", func() {
	Context("GetIdentityMatrix", func() {
		It("s rows should not be nil", func() {
			result := GetIdentityMatrix()

			Expect(result.ROW).NotTo(BeNil())
			Expect(len(result.ROW)).To(Equal(4))
			Expect(result.ROW[0]).NotTo(BeNil())
			Expect(len(result.ROW[0])).To(Equal(4))
			Expect(result.ROW[1]).NotTo(BeNil())
			Expect(len(result.ROW[1])).To(Equal(4))
			Expect(result.ROW[2]).NotTo(BeNil())
			Expect(len(result.ROW[2])).To(Equal(4))
			Expect(result.ROW[3]).NotTo(BeNil())
			Expect(len(result.ROW[3])).To(Equal(4))
		})

		It("returns an identity Matrix", func() {
			result := GetIdentityMatrix()

			expected := Matrix{}
			expected.ROW = make([][]float64, 4)
			expected.ROW[0] = make([]float64, 4)
			expected.ROW[0][0] = 1
			expected.ROW[0][1] = 0
			expected.ROW[0][2] = 0
			expected.ROW[0][3] = 0
			expected.ROW[1] = make([]float64, 4)
			expected.ROW[1][0] = 0
			expected.ROW[1][1] = 1
			expected.ROW[1][2] = 0
			expected.ROW[1][3] = 0
			expected.ROW[2] = make([]float64, 4)
			expected.ROW[2][0] = 0
			expected.ROW[2][1] = 0
			expected.ROW[2][2] = 1
			expected.ROW[2][3] = 0
			expected.ROW[3] = make([]float64, 4)
			expected.ROW[3][0] = 0
			expected.ROW[3][1] = 0
			expected.ROW[3][2] = 0
			expected.ROW[3][3] = 1

			Expect(result).To(Equal(expected))
		})
	})

	Context("GetMatrix", func() {
		It("s rows should not be nil", func() {
			result := GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12))

			Expect(result.ROW).NotTo(BeNil())
			Expect(len(result.ROW)).To(Equal(4))
			Expect(result.ROW[0]).NotTo(BeNil())
			Expect(len(result.ROW[0])).To(Equal(4))
			Expect(result.ROW[1]).NotTo(BeNil())
			Expect(len(result.ROW[1])).To(Equal(4))
			Expect(result.ROW[2]).NotTo(BeNil())
			Expect(len(result.ROW[2])).To(Equal(4))
			Expect(result.ROW[3]).NotTo(BeNil())
			Expect(len(result.ROW[3])).To(Equal(4))
		})

		It("returns an initialized Matrix", func() {
			result := GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12))

			expected := Matrix{}
			expected.ROW = make([][]float64, 4)
			expected.ROW[0] = make([]float64, 4)
			expected.ROW[0][0] = 1
			expected.ROW[0][1] = 2
			expected.ROW[0][2] = 3
			expected.ROW[0][3] = 10
			expected.ROW[1] = make([]float64, 4)
			expected.ROW[1][0] = 4
			expected.ROW[1][1] = 5
			expected.ROW[1][2] = 6
			expected.ROW[1][3] = 11
			expected.ROW[2] = make([]float64, 4)
			expected.ROW[2][0] = 7
			expected.ROW[2][1] = 8
			expected.ROW[2][2] = 9
			expected.ROW[2][3] = 12
			expected.ROW[3] = make([]float64, 4)
			expected.ROW[3][0] = 0
			expected.ROW[3][1] = 0
			expected.ROW[3][2] = 0
			expected.ROW[3][3] = 1

			Expect(result).To(Equal(expected))
		})
	})

	Context("MatMult", func() {
		It("successfully premultiplies by the given Matrix by the identity Matrix, and the content stays the same", func() {
			result := GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12))
			result.MatMult(GetIdentityMatrix())

			expected := GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12))
			Expect(result).To(Equal(expected))
		})

		It("successfully premultiplies an identity Matrix by another", func() {
			result := GetIdentityMatrix()
			matrix := GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12))
			result.MatMult(matrix)

			expected := GetMatrix(1, 2, 3, 4, 5, 6, 7, 8, 9, GetPoint(10, 11, 12))
			Expect(result).To(Equal(expected))
		})

		It("successfully premultiplies a Matrix by another", func() {
			result := GetMatrix(0, 2, 0, 4, 0, 6, 0, 8, 0, GetPoint(5, 7, 2))
			matrix := GetMatrix(1, 0, 3, 0, 5, 0, 7, 0, 9, GetPoint(-2, 3, -4))
			result.MatMult(matrix)

			expected := GetMatrix(0, 26, 0, 20, 0, 30, 0, 86, 0, GetPoint(9, 38, 49))
			Expect(result).To(Equal(expected))
		})
	})
})
