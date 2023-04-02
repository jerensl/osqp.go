package osqp

import (
	"github.com/james-bowman/sparse"
	"github.com/jerensl/osqp.go/internal/pkg/binding"
) 

func NewCSC(m, n int, matrix []float64) *sparse.CSC {
	doxMatrix := sparse.NewDOK(m, n)
	
	for i, val := range matrix {
		r := i / m
		c := i % m

		doxMatrix.Set(r, c, val)
	}

	return doxMatrix.ToCSC()
}

func Setup() *binding.OSQPWorkSpace {
	newOSQP := binding.NewOSQP()

	return newOSQP
}

