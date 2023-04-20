package osqp_test

import (
	"testing"

	"github.com/jerensl/osqp.go"
	"github.com/stretchr/testify/assert"
)

func TestNewCSCMatrix(t *testing.T) {
	mtrx := osqp.NewCSCMatrix([][]float64{{4, 1}, {0, 2}})
	
	assert.Equal(t, mtrx.Data(), []float64{4, 1, 2})
	assert.Equal(t, mtrx.Ind(), []int{0, 0, 1})
	assert.Equal(t, mtrx.Indptr(), []int{0, 1, 3})
}