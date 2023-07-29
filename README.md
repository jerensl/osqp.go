<h3 align="center">osqp.go</h3>

<div align="center">

  [![Status](https://img.shields.io/badge/status-active-success.svg)]() 
  [![GitHub Issues](https://img.shields.io/github/issues/jerensl/osqp.go.svg)](https://github.com/jerensl/osqp.go/issues)
  [![GitHub Pull Requests](https://img.shields.io/github/issues-pr/jerensl/osqp.go.svg)](https://github.com/jerensl/osqp.go/pulls)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)

</div>

---

<p align="center"> The goal of this project is to create a Go interface for OSQP implementation in C. The interface should be easy to use and should allow users to solve optimization problems using OSQP in golang..
    <br> 
</p>

## 📝 Table of Contents
- [📝 Table of Contents](#-table-of-contents)
- [🧐 Problem Statement ](#-problem-statement-)
- [💡 Idea / Solution ](#-idea--solution-)
- [⛓️ Dependencies / Limitations ](#️-dependencies--limitations-)
- [🚀 Future Scope ](#-future-scope-)
- [🏁 Getting Started ](#-getting-started-)
	- [Prerequisites](#prerequisites)
	- [Installing](#installing)
- [🎈 Usage ](#-usage-)
- [⛏️ Built With ](#️-built-with-)
- [✍️ Authors ](#️-authors-)
- [🎉 Acknowledgments ](#-acknowledgments-)

## 🧐 Problem Statement <a name = "problem_statement"></a>
 Currently, there is no Quadratic Programming implementation on golang. Users who want to use Quadratic Programming in golang have to write their own interface or create the new one for their specific use case. Without a Go interface build from OSQP implementation on C, users who want to use Quadratic Programming in Go will have to spend time and resources writing their own interface or using an create the new one for their specific use case. This can lead to suboptimal performance and increased development time.

- IDEAL: The goal of this project is to create a Go interface for OSQP implementation in C. The interface should be easy to use and should allow users to solve optimization problems using OSQP in golang.
- REALITY: Currently, there is no Go interface for OSQP implementation on C. Users who want to use OSQP in Go have to write their own interface or use an create the new one for their specific use case. 
- CONSEQUENCES: Without a Go interface for OSQP implementation on C, users who want to use OSQP in Go will have to spend time and resources writing their own interface or using an alternative such as linear programming for their specific use case. This can lead to suboptimal application.

## 💡 Idea / Solution <a name = "idea"></a>
To achieve this goal, we will create a Go package that provides a simple and intuitive interface for OSQP implementation on C. The package will be designed to be easy to use and will provide users with all the functionality they need to solve optimization problems using OSQP in Go.

## ⛓️ Dependencies / Limitations <a name = "limitations"></a>
- OSQP implementation on C

## 🚀 Future Scope <a name = "future_scope"></a>
The long-term plan is to cover a lot of features available from osqp and build an easy way to use csc matrix and do some calculation utility.

## 🏁 Getting Started <a name = "getting_started"></a>
These instructions will get you a copy of the project up and running on your local machine.
### Prerequisites

- OSQP Library  
You can find detailed instructions on building an osqp library and many tips about quadratic programming using osqp in its own documentation [here](https://osqp.org/docs/get_started/sources.html).


### Installing

To get started you need to compile the OSQP build in C or you can copy from libs directory in the example which not recommended because of the different between machine it's compile, so assuming you are using the same environment like me which is arch linux using wsl 2, you can try to copy them instead.

To install the library you can run the command below.

```
go get -u github.com/jerensl/osqp.go
```

## 🎈 Usage <a name="usage"></a>
See [examples/](https://github.com/jerensl/osqp.go/examples/) for a variety of examples.

**Simple:**

```go
package main
import (
	"fmt"
	"github.com/jerensl/osqp.go"
)

func main() {
	newOSQP := osqp.NewOSQP()
	p_mat, err := osqp.NewCSCMatrix([][]float64{{4, 1}, {0, 2}})
	if err != nil {
		fmt.Println(err)
		return
	}
	a_mat, err := osqp.NewCSCMatrix([][]float64{{1, 1}, {1, 0}, {0, 1}})
	if err != nil {
		fmt.Println(err)
		return
	}
	q := []float64{1.0, 1.0}
	l := []float64{1.0, 0.0, 0.0}
	u := []float64{1.0, 0.7, 0.7}
	newOSQP.Setup(p_mat, q, a_mat, l, u)
	newOSQP.Solve()
	fmt.Println(newOSQP.Solution())
	newOSQP.CleanUp()
}
```

## ⛏️ Built With <a name = "tech_stack"></a>
- [OSQP](https://osqp.org/) - OSQP Library in C

## ✍️ Authors <a name = "authors"></a>
- [@jerensl](https://github.com/jerensl) - Idea & Initial work

See also the list of [contributors](https://github.com/jerensl/osqp.go/contributors) 
who participated in this project.

## 🎉 Acknowledgments <a name = "acknowledgments"></a>
- [OSQP](https://github.com/osqp/osqp)