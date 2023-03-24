package osqp

/*

#cgo CFLAGS: -I./osqp/include
#cgo LDFLAGS: -L./osqp/out -losqp -Wl,-rpath=./osqp/out
#include "osqp.h"
#include <stdlib.h>
#include <stdio.h>
 
*/
import "C"
import (
	"fmt"
	"unsafe"
) 

func Hello() string {
	return "Hello"
}

func Init() {

	// Load problem data
	P_x := []C.c_float{4.0, 1.0, 2.0}
	P_nnz := C.c_int(3)
	P_i := []C.c_int{0, 0, 1}
	P_p := []C.c_int{0, 1, 3}
	q := []C.c_float{1.0, 1.0}
	A_x := []C.c_float{1.0, 1.0, 1.0, 1.0}
	A_nnz := C.c_int(4)
	A_i := []C.c_int{0, 1, 0, 2}
	A_p := []C.c_int{0, 2, 4}
	l := []C.c_float{1.0, 0.0, 0.0}
	u := []C.c_float{1.0, 0.7, 0.7}
	n := C.c_int(2)
	m := C.c_int(3)

	// Workspace structures
	var work *C.OSQPWorkspace
	settings := (*C.OSQPSettings)(C.c_malloc(C.sizeof_OSQPSettings))
	data := (*C.OSQPData)(C.c_malloc(C.sizeof_OSQPData))

	// Populate data
	if data != nil {
		data.n = n
		data.m = m
		data.P = C.csc_matrix(data.n, data.n, P_nnz, (*C.c_float)(unsafe.Pointer(&P_x[0])), (*C.c_int)(unsafe.Pointer(&P_i[0])), (*C.c_int)(unsafe.Pointer(&P_p[0])))
		data.q = (*C.c_float)(unsafe.Pointer(&q[0]))
		data.A = C.csc_matrix(data.m, data.n, A_nnz, (*C.c_float)(unsafe.Pointer(&A_x[0])), (*C.c_int)(unsafe.Pointer(&A_i[0])), (*C.c_int)(unsafe.Pointer(&A_p[0])))
		data.l = (*C.c_float)(unsafe.Pointer(&l[0]))
		data.u = (*C.c_float)(unsafe.Pointer(&u[0]))
	}

	// Define solver settings as default
	if settings != nil {
		C.osqp_set_default_settings(settings)
	}

	// Setup workspace
	C.osqp_setup(&work, data, settings)

	// Solve Problem
	C.osqp_solve(work)

	fmt.Println(*work.solution.x)
	fmt.Println(*work.solution.y)

	// Clean workspace
	C.osqp_cleanup(work)
	if data != nil {
		if data.A != nil {
			C.c_free(unsafe.Pointer(data.A))
		}
		if data.P != nil {
			C.c_free(unsafe.Pointer(data.P))
		}
		C.c_free(unsafe.Pointer(data))
	}
	if settings != nil {
		C.c_free(unsafe.Pointer(settings))
	}

}