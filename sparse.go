package osqp

import (
	"errors"

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

func NewCSCMatrix(matrix [][]float64) (SparseMatrix, error) {
	sparse := SparseMatrix{
		r: len(matrix),
		c: len(matrix[0]),
		indPtr: []int{0},
		ind: []int{},
		data: []float64{},
	}

	totalItem := 0
	for colIdx := 0; colIdx < sparse.c; colIdx++ {
		for rowIdx := 0; rowIdx < sparse.r; rowIdx++ {
			if len(matrix[rowIdx]) != sparse.r {
				return sparse, errors.New("size of the row is not same")
			}
			if matrix[rowIdx][colIdx] != 0.0 {
				sparse.data = append(sparse.data, matrix[rowIdx][colIdx])
				sparse.ind = append(sparse.ind, rowIdx)
				totalItem++
			}
		}
		sparse.indPtr = append(sparse.indPtr, totalItem)
	}



	return sparse, nil
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