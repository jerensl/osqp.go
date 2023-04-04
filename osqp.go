package osqp

import (
	"github.com/james-bowman/sparse"
	"github.com/jerensl/osqp.go/internal/pkg/binding"
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

	return doxMatrix.ToCSC()
}

type Data struct {
	M 		int64
	N 		int64
	P_mat 	*sparse.CSC
	A_mat 	*sparse.CSC
	Q		[]float64
	L		[]float64
	U		[]float64
}

type OSQPConfig struct {
	bind *binding.OSQPWorkSpace
}

func NewOSQP() *OSQPConfig {
	initOSQP := binding.NewOSQP()

	newOSQP := &OSQPConfig{
		bind: initOSQP,
	}

	return newOSQP
}


func (o *OSQPConfig) SetData(newData Data)  {
	currData := binding.Data{
		M: newData.M,
		N: newData.N,
		P_x: newData.P_mat.RawMatrix().Data,
		P_i: newData.P_mat.RawMatrix().Ind,
		P_p: newData.P_mat.RawMatrix().Indptr,
		P_nnz: int64(newData.P_mat.NNZ()),
		A_x: newData.A_mat.RawMatrix().Data,
		A_i: newData.A_mat.RawMatrix().Ind,
		A_p: newData.A_mat.RawMatrix().Indptr,
		A_nnz: int64(newData.A_mat.NNZ()),
		Q: newData.Q,
		L: newData.L,
		U: newData.U,
	}

	// currData := binding.Data{
	// 	M: newData.M,
	// 	N: newData.N,
	// 	P_x: newData.P_mat.Data,
	// 	P_i: newData.P_mat.Ind,
	// 	P_p: newData.P_mat.IdxPtr,
	// 	P_nnz: int64(newData.P_mat.NNZ),
	// 	A_x: newData.A_mat.Data,
	// 	A_i: newData.A_mat.Ind,
	// 	A_p: newData.A_mat.IdxPtr,
	// 	A_nnz: int64(newData.A_mat.NNZ),
	// 	Q: newData.Q,
	// 	L: newData.L,
	// 	U: newData.U,
	// }
	o.bind.SetData(currData)
}

func (o OSQPConfig) Setup()  {
	o.bind.Setup()
}

func (o OSQPConfig) Solve()  {
	o.bind.Solve()
}

func (o OSQPConfig) Solution() (float32, float32) {
	return o.bind.Solution()
}

func (o OSQPConfig) CleanUp() {
	o.bind.CleanUp()
}