package osqp

import (
	"github.com/jerensl/osqp.go/internal/pkg/binding"
) 

type Data struct {
	M 		int64
	N 		int64
	P_mat 	CSCMatrix
	A_mat 	CSCMatrix
	Q		[]float64
	L		[]float64
	U		[]float64
}

type OSQPConfig struct {
	bind *binding.OSQPWorkSpace
}

func NewOSQP() *binding.OSQPWorkSpace {
	newOSQP := binding.NewOSQP()

	return newOSQP
}


func (o *OSQPConfig) SetData(newData Data)  {
	currData := binding.Data{
		M: newData.M,
		N: newData.N,
		P_x: newData.P_mat.Data,
		P_i: newData.P_mat.Ind,
		P_p: newData.P_mat.IdxPtr,
		P_nnz: int64(newData.P_mat.NNZ),
		A_x: newData.A_mat.Data,
		A_i: newData.A_mat.Ind,
		A_p: newData.A_mat.IdxPtr,
		A_nnz: int64(newData.A_mat.NNZ),
		Q: newData.Q,
		L: newData.L,
		U: newData.U,
	}
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