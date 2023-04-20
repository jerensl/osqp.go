package osqp

import (
	"github.com/james-bowman/sparse"
)

type SparseMatrix struct {
	r, c 		int
	indPtr		[]int
	ind			[]int
	data		[]float64
}

func (s SparseMatrix) Data() []float64 {
	return s.data
}

func (s SparseMatrix) Indptr() []int {
	return s.indPtr
}

func (s SparseMatrix) Ind() []int {
	return s.ind
}

func NewCSCMatrix(matrix [][]float64) SparseMatrix {
	sparse := SparseMatrix{
		r: 0,
		c: 0,
		indPtr: []int{0},
		ind: []int{},
		data: []float64{},
	}

	totalItem := 0
	for colIdx := 0; colIdx < len(matrix[0]); colIdx++ {
		for rowIdx := 0; rowIdx < len(matrix); rowIdx++ {
			if matrix[rowIdx][colIdx] != 0.0 {
				sparse.data = append(sparse.data, matrix[rowIdx][colIdx])
				sparse.ind = append(sparse.ind, rowIdx)
				totalItem++
			}
		}
		sparse.indPtr = append(sparse.indPtr, totalItem)
	}

	sparse.r = len(matrix)
	sparse.c = len(matrix[0])

	return sparse 
}

func NewCSCMat(m, n int, matrix [][]float64) *sparse.CSC {
	doxMatrix := sparse.NewDOK(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0.0 {
				val := matrix[i][j]
				doxMatrix.Set(i, j, val)
			}
		}
	}


	csc := doxMatrix.ToCSC()

	return csc
}

func Block_diag() {

}