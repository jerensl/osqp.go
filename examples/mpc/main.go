package main

import (
	"fmt"
	"math"

	"github.com/jerensl/osqp.go"
	"gonum.org/v1/gonum/mat"
)

func main() {
	qp := osqp.NewOSQP()
	
	Ad, err := osqp.NewCSCMatrix([][]float64{
		{1.,      0.,     0., 0., 0., 0., 0.1,     0.,     0.,  0.,     0.,     0.    },
		{0.,      1.,     0., 0., 0., 0., 0.,      0.1,    0.,  0.,     0.,     0.    },
		{0.,      0.,     1., 0., 0., 0., 0.,      0.,     0.1, 0.,     0.,     0.    },
		{0.0488,  0.,     0., 1., 0., 0., 0.0016,  0.,     0.,  0.0992, 0.,     0.    },
		{0.,     -0.0488, 0., 0., 1., 0., 0.,     -0.0016, 0.,  0.,     0.0992, 0.    },
		{0.,      0.,     0., 0., 0., 1., 0.,      0.,     0.,  0.,     0.,     0.0992},
		{0.,      0.,     0., 0., 0., 0., 1.,      0.,     0.,  0.,     0.,     0.    },
		{0.,      0.,     0., 0., 0., 0., 0.,      1.,     0.,  0.,     0.,     0.    },
		{0.,      0.,     0., 0., 0., 0., 0.,      0.,     1.,  0.,     0.,     0.    },
		{0.9734,  0.,     0., 0., 0., 0., 0.0488,  0.,     0.,  0.9846, 0.,     0.    },
		{0.,     -0.9734, 0., 0., 0., 0., 0.,     -0.0488, 0.,  0.,     0.9846, 0.    },
		{0.,      0.,     0., 0., 0., 0., 0.,      0.,     0.,  0.,     0.,     0.9846},
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	Bd, err := osqp.NewCSCMatrix([][]float64{
		{0.,      -0.0726,  0.,     0.0726},
		{-0.0726,  0.,      0.0726, 0.    },
		{-0.0152,  0.0152, -0.0152, 0.0152},
		{-0.,     -0.0006, -0.,     0.0006},
		{0.0006,   0.,     -0.0006, 0.0000},
		{0.0106,   0.0106,  0.0106, 0.0106},
		{0,       -1.4512,  0.,     1.4512},
		{-1.4512,  0.,      1.4512, 0.    },
		{-0.3049,  0.3049, -0.3049, 0.3049},
		{-0.,     -0.0236,  0.,     0.0236},
		{0.0236,   0.,     -0.0236, 0.    },
		{0.2107,   0.2107,  0.2107, 0.2107},
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	nx, nu := Bd.Dimension()

	// Constraints
	u0 := 10.5916
	umin := mat.NewVecDense(4, []float64{9.6-u0, 9.6-u0, 9.6-u0, 9.6-u0})
	umax := mat.NewVecDense(4, []float64{13.0-u0, 13.0-u0, 13.0-u0, 13.0-u0})
	xmin := mat.NewVecDense(12, []float64{-math.Pi/6, -math.Pi/6, math.Inf(-1), math.Inf(-1), math.Inf(-1), -1,
		math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1), math.Inf(-1),
	})
	xmax := mat.NewVecDense(12, []float64{math.Pi/6, math.Pi/6, math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1),
		math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1),	
	})

	// Objective function
	Q := mat.NewDiagDense(12,[]float64{0., 0., 10., 10., 10., 10., 0., 0., 0., 5., 5., 5.})
	QN := Q
	R := mat.NewDiagDense(4, []float64{0.1, 0.1, 0.1, 0.1})
	
	// Initial and reference states
	x0 := mat.NewVecDense(12, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,})
	xr := mat.NewVecDense(12, []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,})

	// Prediction horizon
	N := 10

	var QK mat.Dense
	QK.Kronecker(osqp.DenseEye(N, 1.0), Q)
	
	var RK mat.Dense
	RK.Kronecker(osqp.DenseEye(N, 1.0), R)

	row, col := QN.Dims()

	QNDense := mat.NewDense(row, col, nil)
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            QNDense.Set(i, j, QN.At(i, j))
        }
    } 

	// Cast MPC problem to a QP: x = (x(0),x(1),...,x(N),u(0),...,u(N-1))
	// - quadratic objective
	P := osqp.BlockDiag([]*mat.Dense{&QK, QNDense, &RK})

	var QDot mat.VecDense
	QDot.MulVec(Q, xr)

	Ones := mat.NewVecDense(10, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	Nnu := mat.NewVecDense(40, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	var QKron mat.Dense
	QKron.Kronecker(Ones, osqp.ToNegativeVecDense(QDot))

	// - linear objective
	var hstack mat.Dense
	var q mat.Dense
	
	hstack.Stack(&QKron, osqp.ToNegativeVecDense(QDot))
	q.Stack(&hstack, Nnu)

	_ = P

	// - linear dynamics
	var AxSubA mat.Dense
	var AxSubB mat.Dense
	var Ax mat.Dense

	AxSubA.Kronecker(osqp.DenseEye(N+1, 1.0), osqp.ToNegativeDense(*osqp.DenseEye(nx, 1.0)))
	AxSubB.Kronecker(osqp.DenseEyeK(N+1, 1.0, -1), Ad.ToDense())
	
	Ax.Add(&AxSubA, &AxSubB)
	
	CSCMat, err := osqp.NewCSCMatrix([][]float64{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}})
	if err != nil {
		fmt.Println(err)
		return
	}

	var vstack mat.Dense
	vstack.Stack(CSCMat.ToDense(), osqp.DenseEye(N, 1.0))

	var Bu mat.Dense
	Bu.Kronecker(&vstack, Bd.ToDense())
	var Aeq mat.Dense
	Aeq.Stack(&Ax, (&Bu).T())
	var leq mat.Dense
	leq.Stack(osqp.ToNegative(x0), osqp.VecZeros(N*nx))
	ueq := leq

	// - input and state constraints
	Aineq := osqp.DenseEye((N+1)*nx+N*nu, 1.0)
	var lineq mat.Dense
	var lineqKron mat.Dense
	lineqKron.Kronecker(osqp.VecZeros(N+1), xmin)
	var lineqKronTwo mat.Dense
	lineqKronTwo.Kronecker(osqp.VecZeros(N), umin)
	lineq.Stack(&lineqKron, &lineqKronTwo)

	var uineqKron mat.Dense
	uineqKron.Kronecker(osqp.VecZeros(N+1), xmax)
	var uineqKronTwo mat.Dense
	uineqKronTwo.Kronecker(osqp.VecZeros(N), umax)
	var uineq mat.Dense
	uineq.Stack(&uineqKron, &uineqKronTwo)

	_ = Aineq
	// - OSQP constraints
	var A mat.Dense
	var l mat.Dense
	var u mat.Dense

	A.Stack((&Aeq).T(), Aineq) 
	l.Stack(&leq, &lineq)
	u.Stack(&ueq, &uineq)


	p, err := osqp.NewCSCDenseMatrix(P)
	if err != nil {
		fmt.Println(err)
		return
	}

	a, err := osqp.NewCSCDenseMatrix(&A)
	if err != nil {
		fmt.Println(err)
		return
	}

	qp.Setup(p, q.RawMatrix().Data, a, l.RawMatrix().Data, u.RawMatrix().Data)

		qp.Solve()
		x, y := qp.Solution()

		fmt.Println(x)
		fmt.Println(y)

	qp.CleanUp()
}

