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

	q := []float64{1.0, 1.0}
	l := []float64{1.0, 0.0, 0.0}
	u := []float64{1.0, 0.7, 0.7}

	newOSQP.Setup(p_mat, q, a_mat, l, u)

	newOSQP.Solve()

	fmt.Println(newOSQP.Solution())

	newOSQP.CleanUp()
}

 