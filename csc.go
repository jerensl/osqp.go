package osqp

type CSCMatrix struct {
	Data 	[]float64
	Ind		[]int
	IdxPtr	[]int
	NNZ		int
}

func NewCSCMatrix(m, n int, matrix [][]float64) CSCMatrix {
	mat := CSCMatrix{
		Data: []float64{},
		Ind: []int{},
		IdxPtr: []int{0},
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[j][i] != 0.0 {
				count++
				mat.Data = append(mat.Data, matrix[j][i])
				mat.Ind = append(mat.Ind, j)
			}
		}
		mat.IdxPtr = append(mat.IdxPtr, count)
	}

	mat.NNZ = len(mat.Data)

	return mat
}