package osqp

import "github.com/jerensl/osqp.go/internal/pkg/binding"

func Setup() *binding.OSQPWorkSpace {
	newOSQP := binding.NewOSQP()

	return newOSQP
}

