package main

import (
	"fmt"

	"github.com/jerensl/osqp.go"
)

func main() {
	newOSQP := osqp.NewOSQP()

	p_mat := osqp.NewCSCMat(2, 2, [][]float64{{4.0, 1.0}, {0.0, 2.0}})
	a_mat := osqp.NewCSCMat(3, 2, [][]float64{{1.0, 1.0}, {1.0, 0.0}, {0.0, 1.0}})

	fmt.Println(p_mat)
	fmt.Println(a_mat)
	data := osqp.Data{
		P_mat: p_mat,
		A_mat: a_mat,
		Q: []float64{1.0, 1.0},
		L: []float64{1.0, 0.0, 0.0},
		U: []float64{1.0, 0.7, 0.7},
		N: 2,
		M: 3,
	}

	newOSQP.SetData(data)
	
	newOSQP.Setup()

	newOSQP.Solve()

	newOSQP.CleanUp()
}

