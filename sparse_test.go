package osqp_test

import (
	"testing"

	"github.com/jerensl/osqp.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type expectedData struct {
	data	[]float64
	ind		[]int
	indPtr	[]int
	nnz		int
}

func TestValidCSCMatrix(t *testing.T) {
	testCases := []struct {
		name			string
		matrix			[][]float64
		expected		expectedData
	}{
		{
			name: "Test new Matrix",
			matrix: [][]float64{{4, 1}, {0, 2}},
			expected: expectedData{
				data: []float64{4, 1, 2},
				ind: []int{0, 0, 1},
				indPtr: []int{0, 1, 3},
				nnz: 3,
			},
		},
		{
			name: "Test new Matrix with bigger size",
			matrix: [][]float64{
				{8, 0, 2, 0, 0},
				{0, 0, 5, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 7, 1, 2},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 9, 0}},
			expected:expectedData {
				data: []float64{8, 2, 5, 7, 1, 9, 2},
				ind: []int{0, 0, 1, 4, 4, 6, 4},
				indPtr: []int{0, 1, 1, 4, 6, 7},
				nnz: 7,
			},
		},
		{
			name: "Test new Matrix with float number",
			matrix: [][]float64{{4.0, 1.0}, {0.0, 2.0}},
			expected:expectedData {
				data: []float64{4, 1, 2},
				ind: []int{0, 0, 1},
				indPtr: []int{0, 1, 3},
				nnz: 3,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			mtrx, err := osqp.NewCSCMatrix(tC.matrix)
			require.NoError(t, err)

			assert.Equal(t, tC.expected.data, mtrx.Data())
			assert.Equal(t, tC.expected.ind, mtrx.Ind())
			assert.Equal(t, tC.expected.indPtr, mtrx.IndPtr())
			assert.Equal(t, tC.expected.nnz, mtrx.NNZ())
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
