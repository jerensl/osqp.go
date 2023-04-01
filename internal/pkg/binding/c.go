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
	n 		int64
	m 		int64
	p_x 	[]C.c_float
	p_i 	[]C.c_int
	p_p 	[]C.c_int
	p_nnz	C.c_int
	a_x 	[]C.c_float
	a_i 	[]C.c_int
	a_p 	[]C.c_int
	a_nnz	C.c_int
	q 		[]C.c_float
	l 		[]C.c_float
	u 		[]C.c_float
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
	C.osqp_solve(o.work)
}

func (o *OSQPWorkSpace) SetData(newData Data) {
	data := (*C.OSQPData)(C.c_malloc(C.sizeof_OSQPData))

	if data != nil {
		data.n = (C.c_int)(newData.n)
		data.m = (C.c_int)(newData.m)

		data.P = C.csc_matrix(data.n, data.n, newData.p_nnz, (*C.c_float)(unsafe.Pointer(&newData.p_x[0])), (*C.c_int)(unsafe.Pointer(&newData.p_i[0])), (*C.c_int)(unsafe.Pointer(&newData.p_p[0])))
		data.q = (*C.c_float)(unsafe.Pointer(&newData.q))
		data.A = C.csc_matrix(data.m, data.n, newData.a_nnz, (*C.c_float)(unsafe.Pointer(&newData.a_x[0])), (*C.c_int)(unsafe.Pointer(&newData.a_i[0])), (*C.c_int)(unsafe.Pointer(&newData.a_p[0])))

		data.l = (*C.c_float)(unsafe.Pointer(&newData.l[0]))
		data.u = (*C.c_float)(unsafe.Pointer(&newData.u[0]))
	}
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
