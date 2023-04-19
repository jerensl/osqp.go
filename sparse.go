package osqp

import (
	"github.com/james-bowman/sparse"
)


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

