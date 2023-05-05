package osqp

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

type SparseMatrix struct {
	r, c 		int
	nnz			int64
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

func (s SparseMatrix) NNZ() int64 {
	return s.nnz
}

func (s SparseMatrix) Dimension() (int, int) {
	return s.r, s.c
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
			if matrix[rowIdx][colIdx] != 0 {
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

func NewCSCDenseMatrix(matrix mat.Matrix) (SparseMatrix, error) {
	row, col := matrix.Dims()
	
	sparse := SparseMatrix{
		r: row,
		c: col,
		nnz: 0,
		indPtr: []int{0},
		ind: []int{},
		data: []float64{},
	}

	totalItem := 0
	for colIdx := 0; colIdx < sparse.c; colIdx++ {
		for rowIdx := 0; rowIdx < sparse.r; rowIdx++ {
			if matrix.At(rowIdx, colIdx) != 0 {
				sparse.data = append(sparse.data, matrix.At(rowIdx,colIdx))
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
	   
	   indPos := 0

	   col := 0

	   for i := 0; i < len(s.indPtr)-1; i++ {
			totalNumInRow := s.indPtr[i+1] - s.indPtr[i]
			if totalNumInRow > 0 {
				for totalNumInRow > 0 {
					matrix[s.ind[indPos]][col] = s.data[indPos]
					indPos++
					totalNumInRow--
				}
				col++
			} else {
				col++
			}
	   }

	return matrix
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

func DenseEye(size int, val float64) *mat.Dense {
	matrix := make([]float64, size*size)
	
	for i := range matrix {
		r := i/size
		c := i%size
		if r == c {
			matrix[i] = val
		}
	}

	matDense := mat.NewDense(size, size, matrix)
	
	return matDense
}

func DenseEyeK(size int, val float64, k int) *mat.Dense {
	matrix := make([]float64, size*size)
	count := 1
	
	for i := range matrix {
		r := i/size
		c := i%size
		if r == count && c == r+k {
			matrix[i] = val
			count++
		}
	}

	matDense := mat.NewDense(size, size, matrix)
	
	return matDense
}

func ToNegativeDense(vecDense mat.Dense) *mat.Dense {
	row, col := vecDense.Dims()

	newMat := []float64{}

	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			newMat = append(newMat, -vecDense.At(r, c))
		}
	}

	newDense := mat.NewDense(row, col, newMat)

	return newDense
}

func VStack(matrix []mat.Matrix) *mat.Dense {
	rows := 0
	cols := 0

	var newMatrix *mat.Dense
	
	for i := range matrix {
		r, c := matrix[i].Dims()
		rows += r
		cols += c

		newMatrix.Copy(matrix[i])
	}

	return newMatrix
}

func ToNegativeVecDense(vecDense mat.VecDense) *mat.Dense {
	row, col := vecDense.Dims()

	newMat := []float64{}

	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			newMat = append(newMat, -vecDense.At(r, c))
		}
	}

	newDense := mat.NewDense(row, col, newMat)

	return newDense
}

func VecZeros(num int) mat.Matrix {
	items := []float64{}

	for i := 0; i < num; i++ {
		items = append(items, 0)
	}

	newVec := mat.NewVecDense(num, items)
	
	return newVec
}

func ToNegative(matDense mat.Matrix) mat.Matrix {
	return toNegative(matDense)
}

func toNegative(matDense mat.Matrix) mat.Matrix {
	row, col := matDense.Dims()

	newMat := []float64{}

	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			newMat = append(newMat, -matDense.At(r, c))
		}
	}

	newDense := mat.NewDense(row, col, newMat)

	return newDense
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