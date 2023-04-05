package main

import (
	"github.com/jerensl/osqp.go"
)

func main() {
	newOSQP := osqp.NewOSQP()

	p_mat := osqp.NewCSCMat(2, 2, [][]float64{{4, 1}, {0, 2}})
	a_mat := osqp.NewCSCMat(3, 2, [][]float64{{1, 1}, {1, 0}, {0, 1}})

	data := osqp.Data{
		P_mat: p_mat,
		Q: []float64{1.0, 1.0},
		A_mat: a_mat,
		L: []float64{1.0, 0.0, 0.0},
		U: []float64{1.0, 0.7, 0.7},
		N: 2,
		M: 3,
	}

	newOSQP.Setup(data)

	newOSQP.Solve()

	newOSQP.CleanUp()
}

