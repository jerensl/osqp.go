package main

import (
	"math"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// newOSQP := osqp.NewOSQP()

	
	// Ad := osqp.NewCSCMat(12, 12, [][]float64{
	// 	{1.,      0.,     0., 0., 0., 0., 0.1,     0.,     0.,  0.,     0.,     0.    },
	// 	{0.,      1.,     0., 0., 0., 0., 0.,      0.1,    0.,  0.,     0.,     0.    },
	// 	{0.,      0.,     1., 0., 0., 0., 0.,      0.,     0.1, 0.,     0.,     0.    },
	// 	{0.0488,  0.,     0., 1., 0., 0., 0.0016,  0.,     0.,  0.0992, 0.,     0.    },
	// 	{0.,     -0.0488, 0., 0., 1., 0., 0.,     -0.0016, 0.,  0.,     0.0992, 0.    },
	// 	{0.,      0.,     0., 0., 0., 1., 0.,      0.,     0.,  0.,     0.,     0.0992},
	// 	{0.,      0.,     0., 0., 0., 0., 1.,      0.,     0.,  0.,     0.,     0.    },
	// 	{0.,      0.,     0., 0., 0., 0., 0.,      1.,     0.,  0.,     0.,     0.    },
	// 	{0.,      0.,     0., 0., 0., 0., 0.,      0.,     1.,  0.,     0.,     0.    },
	// 	{0.9734,  0.,     0., 0., 0., 0., 0.0488,  0.,     0.,  0.9846, 0.,     0.    },
	// 	{0.,     -0.9734, 0., 0., 0., 0., 0.,     -0.0488, 0.,  0.,     0.9846, 0.    },
	// 	{0.,      0.,     0., 0., 0., 0., 0.,      0.,     0.,  0.,     0.,     0.9846},
	// 	})
	// Bd := osqp.NewCSCMat(4, 12, [][]float64{
	// 	{0.,      -0.0726,  0.,     0.0726},
	// 	{-0.0726,  0.,      0.0726, 0.    },
	// 	{-0.0152,  0.0152, -0.0152, 0.0152},
	// 	{-0.,     -0.0006, -0.,     0.0006},
	// 	{0.0006,   0.,     -0.0006, 0.0000},
	// 	{0.0106,   0.0106,  0.0106, 0.0106},
	// 	{0,       -1.4512,  0.,     1.4512},
	// 	{-1.4512,  0.,      1.4512, 0.    },
	// 	{-0.3049,  0.3049, -0.3049, 0.3049},
	// 	{-0.,     -0.0236,  0.,     0.0236},
	// 	{0.0236,   0.,     -0.0236, 0.    },
	// 	})
	// nx, nu := 4, 12

	// Constraints
	u0 := 10.5916
	umin := mat.NewDense(1, 4, []float64{9.6-u0, 9.6-u0, 9.6-u0, 9.6-u0})
	umax := mat.NewDense(1, 4, []float64{13.0-u0, 13.0-u0, 13.0-u0, 13.0-u0})
	xmin := mat.NewDense(1, 12, []float64{-math.Pi/6, -math.Pi/6, math.Inf(-1), math.Inf(-1), math.Inf(-1), -1,
		math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1),
	})
	xmax := mat.NewDense(1, 12, []float64{math.Pi/6, math.Pi/6, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1),
		math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1),	
	})

	_ = umin
	_ = umax
	_ = xmin
	_ = xmax

	// Objective function
	Q := sparse.NewDIA(12,12, []float64{0., 0., 10., 10., 10., 10., 0., 0., 0., 5., 5., 5.})
	QN := Q
	R := sparse.NewDIA(4,4, []float64{0.1, 0.1, 0.1, 0.1})


	_ = QN
	_ = R

	
	// Initial and reference states
	x0 := mat.NewVecDense(12, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,})
	xr := mat.NewVecDense(12, []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,})

	_ = x0 
	_ = xr 
	// fmt.Println(x0)
	// fmt.Println(xr)

	// Prediction horizon
	N := sparse.NewDIA(10,10, []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0})

	var QK mat.Dense
	var RK mat.Dense
	QK.Kronecker(N, Q)

	RK.Kronecker(N, R)


	row, col := QN.Dims()

	QNDense := mat.NewDense(row, col, nil)
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            QNDense.Set(i, j, QN.At(i, j))
        }
    }

	blocks := []*mat.Dense{
       &QK,
	   QNDense,
	   &RK,
    }

    var rows int
    var cols int
    for _, m := range blocks {
        r, c := m.Dims()
        rows += r
        cols += c
    }

	// Cast MPC problem to a QP: x = (x(0),x(1),...,x(N),u(0),...,u(N-1))
	// - quadratic objective
    P := mat.NewDense(rows, cols, nil)
    rowOffset := 0
    colOffset := 0
    for _, m := range blocks {
        r, c := m.Dims()
        P.Slice(rowOffset,rowOffset+r,colOffset,colOffset+c).(*mat.Dense).Copy(m)
        rowOffset += r
        colOffset += c
    }


	// fmt.Println(RK.RawMatrix())

	// newOSQP.Solve()

	// newOSQP.CleanUp()
}

