package main

import (
	"fmt"

	"github.com/jerensl/osqp.go"
)

func main() {
	newOSQP := osqp.NewOSQP()

	p_mat, err := osqp.NewCSCMatrix([][]float64{{4, 1}, {0, 2}})
	if err != nil {
		fmt.Println(err)
		return
	}

	a_mat, err := osqp.NewCSCMatrix([][]float64{{1, 1}, {1, 0}, {0, 1}})
	if err != nil {
		fmt.Println(err)
		return
	}

	t_mat, err := osqp.NewDiagCSCMatrix(4, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p_mat)
	fmt.Println(p_mat.ToDense())
	fmt.Println(t_mat.ToDense())

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

 