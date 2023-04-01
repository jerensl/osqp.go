package binding

/*
#cgo CFLAGS: -I../../../build/include
#cgo LDFLAGS: -L../../../build/out -losqp -Wl,-rpath=./build/out
#include "osqp.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"unsafe"
) 

type Data struct {
	N 		int64
	M 		int64
	P_x 	[]float64
	P_i 	[]int64
	P_p 	[]int64
	P_nnz	int64
	A_x 	[]float64
	A_i 	[]int64
	A_p 	[]int64
	A_nnz	int64
	Q 		[]float64
	L 		[]float64
	U 		[]float64
}

type OSQPWorkSpace struct {
	work 		*C.OSQPWorkspace
	settings 	*C.OSQPSettings
	data 		*C.OSQPData
}

func NewOSQP() *OSQPWorkSpace {
	settings := (*C.OSQPSettings)(C.c_malloc(C.sizeof_OSQPSettings))

	if settings != nil {
		C.osqp_set_default_settings(settings)
	}

	return &OSQPWorkSpace{
		settings: settings,
	}
}

func (o *OSQPWorkSpace) Solve() {
	// Setup workspace
	C.osqp_setup(&o.work, o.data, o.settings)
	
	C.osqp_solve(o.work)
}

func (o *OSQPWorkSpace) SetData(newData Data) {
	data := (*C.OSQPData)(C.c_malloc(C.sizeof_OSQPData))

	if data != nil {
		data.n = (C.c_int)(newData.N)
		data.m = (C.c_int)(newData.M)

		data.P = C.csc_matrix(data.n, data.n, (C.c_int)(newData.P_nnz), (*C.c_float)(unsafe.Pointer(&newData.P_x[0])), (*C.c_int)(unsafe.Pointer(&newData.P_i[0])), (*C.c_int)(unsafe.Pointer(&newData.P_p[0])))
		data.q = (*C.c_float)(unsafe.Pointer(&newData.Q))
		data.A = C.csc_matrix(data.m, data.n, (C.c_int)(newData.A_nnz), (*C.c_float)(unsafe.Pointer(&newData.A_x[0])), (*C.c_int)(unsafe.Pointer(&newData.A_i[0])), (*C.c_int)(unsafe.Pointer(&newData.A_p[0])))

		data.l = (*C.c_float)(unsafe.Pointer(&newData.L[0]))
		data.u = (*C.c_float)(unsafe.Pointer(&newData.U[0]))
	}

	o.data = data
}

func (o *OSQPWorkSpace) CleanUp() {
	C.osqp_cleanup(o.work)

	if o.data != nil {
		if o.data.A != nil {
			C.c_free(unsafe.Pointer(o.data.A))
		}
		if o.data.P != nil {
			C.c_free(unsafe.Pointer(o.data.P))
		}
		C.c_free(unsafe.Pointer(o.data))
	}
	if o.settings != nil {
		C.c_free(unsafe.Pointer(o.settings))
	}
}
