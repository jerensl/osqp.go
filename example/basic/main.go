package main

import (
	"fmt"

	"github.com/jerensl/osqp.go"
	"github.com/jerensl/osqp.go/internal/pkg/binding"
)

func main() {
	newOSQP := osqp.NewOSQP()

	// p_mat := osqp.NewCSCMat(2, 2, []float64{4.0, 1.0, 1.0, 2.0})
	// a_mat := osqp.NewCSCMat(3, 2, []float64{1.0, 1.0, 1.0, 0.0, 0.0, 1.0})

	// data := osqp.Data{
	// 	P_mat: p_mat,
	// 	A_mat: a_mat,
	// 	Q: []float64{1.0, 1.0},
	// 	L: []float64{1.0, 0.0, 0.0},
	// 	U: []float64{1.0, 0.7, 0.7},
	// 	N: 2,
	// 	M: 3,
	// }

	p_mat := osqp.NewCSCMatrix(2, 2, [][]float64{{4.0, 1.0}, {0.0, 2.0}})
	a_mat := osqp.NewCSCMatrix(3, 2, [][]float64{{1.0, 1.0}, {1.0, 0.0}, {0.0, 1.0}})

	data := binding.Data{
		N: 2,
		M: 3,
		P_x: p_mat.Data,
		P_i: p_mat.Ind,
		P_p: p_mat.IdxPtr,
		P_nnz: int64(p_mat.NNZ),
		A_x: a_mat.Data,
		A_i: a_mat.Ind,
		A_p: a_mat.IdxPtr,
		A_nnz: int64(a_mat.NNZ),
		Q: []float64{1.0, 1.0},
		L: []float64{1.0, 0.0, 0.0},
		U: []float64{1.0, 0.7, 0.7},
	}

	newOSQP.SetData(data)
	
	newOSQP.Setup()

	newOSQP.Solve()

	x, y := newOSQP.Solution()

	fmt.Println(x)
	fmt.Println(y)

	newOSQP.CleanUp()
}

