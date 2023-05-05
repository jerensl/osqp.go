package osqp

import (
	"github.com/jerensl/osqp.go/internal/pkg/binding"
) 

type Data struct {
	M 		int64
	N 		int64
	P_mat 	SparseMatrix
	Q		[]float64
	A_mat 	SparseMatrix
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

func (o OSQPConfig) Setup(p SparseMatrix, q []float64, a SparseMatrix, l []float64, u []float64)  {
	currData := binding.Data{
		M: int64(a.r),
		N: int64(p.r),
		P_x: p.Data(),
		P_i: p.Ind(),
		P_p: p.IndPtr(),
		P_nnz: p.NNZ(),
		Q: q,
		A_x: a.Data(),
		A_i: a.Ind(),
		A_p: a.IndPtr(),
		A_nnz: a.NNZ(),
		L: l,
		U: u,
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