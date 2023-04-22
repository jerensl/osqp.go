package osqp

import (
	"github.com/jerensl/osqp.go/internal/pkg/binding"
) 



type Data struct {
	M 		int64
	N 		int64
	P_mat 	SparseMatrix
	A_mat 	SparseMatrix
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

func (o OSQPConfig) Setup(newData Data)  {
	currData := binding.Data{
		M: newData.M,
		N: newData.N,
		P_x: newData.P_mat.Data(),
		P_i: newData.P_mat.Ind(),
		P_p: newData.P_mat.IndPtr(),
		P_nnz: int64(newData.P_mat.NNZ()),
		A_x: newData.A_mat.Data(),
		A_i: newData.A_mat.Ind(),
		A_p: newData.A_mat.IndPtr(),
		A_nnz: int64(newData.A_mat.NNZ()),
		Q: newData.Q,
		L: newData.L,
		U: newData.U,
	}

	o.bind.Setup(currData)
}

func (o OSQPConfig) Solve()  {
	o.bind.Solve()
}

func (o OSQPConfig) UpdateLinCost(qNew []float64)  {
	o.bind.UpdateLinCost(qNew)
}

func (o OSQPConfig) UpdateBounds(lNew, uNew []float64)  {
	o.bind.UpdateBounds(lNew, uNew)
}

func (o OSQPConfig) UpdatePMat(pNew []float64)  {
	o.bind.UpdatePMat(pNew)
}

func (o OSQPConfig) UpdateAMat(aNew []float64)  {
	o.bind.UpdateAMat(aNew)
}

func (o OSQPConfig) Solution() (float32, float32) {
	return o.bind.Solution()
}

func (o OSQPConfig) CleanUp() {
	o.bind.CleanUp()
}