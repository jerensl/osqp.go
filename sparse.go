package osqp

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type SparseMatrix struct {
	r, c 		int
	nnz			int
	indPtr		[]int
	ind			[]int
	data		[]float64
}

func (s SparseMatrix) Data() []float64 {
	return s.data
}

func (s SparseMatrix) IndPtr() []int {
	return s.indPtr
}

func (s SparseMatrix) Ind() []int {
	return s.ind
}

func (s SparseMatrix) NNZ() int {
	return s.nnz
}

func NewCSCMatrix(matrix [][]float64) (SparseMatrix, error) {
	sparse := SparseMatrix{
		r: len(matrix),
		c: len(matrix[0]),
		nnz: 0,
		indPtr: []int{0},
		ind: []int{},
		data: []float64{},
	}

	totalItem := 0
	for colIdx := 0; colIdx < sparse.c; colIdx++ {
		for rowIdx := 0; rowIdx < sparse.r; rowIdx++ {
			if len(matrix[rowIdx]) != sparse.c {
				return sparse, errors.New("size of the row is not same")
			}
			if matrix[rowIdx][colIdx] != 0.0 {
				sparse.data = append(sparse.data, matrix[rowIdx][colIdx])
				sparse.ind = append(sparse.ind, rowIdx)
				totalItem++
				sparse.nnz++
			}
		}
		sparse.indPtr = append(sparse.indPtr, totalItem)
	}
	

	return sparse, nil
}

func NewDiagCSCMatrix(size int, value float64) (*SparseMatrix, error) {
	sparse := SparseMatrix{
		r: size,
		c: size,
		nnz: 0,
		indPtr: []int{},
		ind: []int{},
		data: []float64{},
	}

	for s := 0; s <= size; s++ {
		sparse.data = append(sparse.data, value)
		sparse.ind = append(sparse.ind, s)
		sparse.indPtr = append(sparse.indPtr, s)
		sparse.nnz++
	}

	return &sparse, nil
}

func (s SparseMatrix) Transpose(matrix [][]float64) [][]float64 {
    rows := len(matrix)
    cols := len(matrix[0])

    transpose := make([][]float64, cols)
    for i := range transpose {
        transpose[i] = make([]float64, rows)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            transpose[j][i] = matrix[i][j]
        }
    }

    return transpose
}


func (s SparseMatrix) unmarshalFromCSC() [][]float64 {
	   matrix := make([][]float64, s.r)
	   for i := range matrix {
		   matrix[i] = make([]float64, s.c)
	   }
	   
	   pos := 0

	   row := 0

	   for i := 0; i < len(s.indPtr)-1; i++ {
			totalItemRow := s.indPtr[i+1] - s.indPtr[i]
			if totalItemRow > 0 {
				for totalItemRow > 0 {
					fmt.Println(s.data[pos])
					matrix[row][s.ind[pos]] = s.data[pos]
					pos++
					totalItemRow--
				}
				row++
			} else {
				row++
			}
	   }


	return s.Transpose(matrix)
}

func (s SparseMatrix) ToDense() *mat.Dense {
	var newMat []float64
	matrix := s.unmarshalFromCSC()
	for _, row := range matrix {
		newMat = append(newMat, row...)
    }

	matrixDense := mat.NewDense(s.r, s.c, newMat)
	return matrixDense
}

func BlockDiag(blocks []*mat.Dense) *mat.Dense { 
	 var rows int
	 var cols int
	 for _, m := range blocks {
		 r, c := m.Dims()
		 rows += r
		 cols += c
	 }
 
	 result := mat.NewDense(rows, cols, nil)
	 rowOffset := 0
	 colOffset := 0
	 for _, m := range blocks {
		 r, c := m.Dims()
		 result.Slice(rowOffset,rowOffset+r,colOffset,colOffset+c).(*mat.Dense).Copy(m)
		 rowOffset += r
		 colOffset += c
	 }
	return result
}