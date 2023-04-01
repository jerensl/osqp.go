package main

import (
	"github.com/jerensl/osqp.go"
	"github.com/jerensl/osqp.go/internal/pkg/binding"
)

func main() {
	newOSQP := osqp.Setup()

	data := binding.Data{
		P_x: []float64{4.0, 1.0, 2.0},
		P_nnz: 3,
		P_i: []int64{0, 0, 1},
		P_p: []int64{0, 1, 3},
		Q: []float64{1.0, 1.0},
		A_x: []float64{1.0, 1.0, 1.0, 1.0},
		A_nnz: 4,
		A_i: []int64{0, 1, 0, 2},
		A_p: []int64{0, 2, 4},
		L: []float64{1.0, 0.0, 0.0},
		U: []float64{1.0, 0.7, 0.7},
		N: 2,
		M: 3,
	}
	newOSQP.SetData(data)

	newOSQP.Solve()
	newOSQP.CleanUp()
}

