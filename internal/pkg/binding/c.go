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

	"gonum.org/v1/gonum/mat"
) 

type Data struct {
	n 		int64
	m 		int64
	mat_p 	mat.Dense
	mat_a 	mat.Dense
	q 		float32
	l 		float32
	u 		float32
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
