package osqp_test

import (
	"testing"

	"github.com/jerensl/osqp.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidCSCMatrix(t *testing.T) {
	testCases := []struct {
		name			string
		matrix			[][]float64
		expectData		[]float64
	}{
		{
			name: "Test new Matrix",
			matrix: [][]float64{{4, 1}, {0, 2}},
			expectData: []float64{4, 1, 2},
		},
		{
			name: "Test new Matrix with float number",
			matrix: [][]float64{{4.0, 1.0}, {0.0, 2.0}},
			expectData: []float64{4, 1, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mtrx, err := osqp.NewCSCMatrix(tC.matrix)
			require.NoError(t, err)

			assert.Equal(t, tC.expectData, mtrx.Data())
		})
	}
}

func TestInvalidCSCMatrix(t *testing.T) {
	testCases := []struct {
		name			string
		matrix			[][]float64
		expectError		string
	}{
		{
			name: "Test new Matrix",
			matrix: [][]float64{{4}, {0, 2}},
			expectError: "size of the row is not same",
		},
		{
			name: "Test new Matrix",
			matrix: [][]float64{{4, 1}, {2}},
			expectError: "size of the row is not same",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := osqp.NewCSCMatrix(tC.matrix)
			assert.ErrorContains(t, err, tC.expectError)
		})
	}
}
