package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// matPrint is a convinient function to pretty print the matrices
func matPrint(X mat.Matrix) {

	// Formatted method helps format the output
	// Prefix method helps us intend the rows when printing the matrix
	// Squeeze method sets the width for each column based on the entries
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func main() {

	// number of independent variables
	numOfvars := 2
	// number of records used
	observations := 4

	// slice of slice representing the independent variables
	data := [][]float64{
		[]float64{1, 2},
		[]float64{2, 3},
		[]float64{3, 3},
		[]float64{4, 8}}

	// dependent variable
	observed := []float64{5, 8, 8, 20}

	// creating empty X and y matrices
	// note: X has one extra column to accomodate the constant
	X := mat.NewDense(4, 3, nil)
	y := mat.NewDense(4, 1, nil)

	// loading data into X and y
	// note: the constant is initialized with 1
	for i := 0; i < observations; i++ {
		y.Set(i, 0, observed[i])
		for j := 0; j < numOfvars+1; j++ {
			if j == 0 {
				X.Set(i, 0, 1)
			} else {
				X.Set(i, j, data[i][j-1])
			}
		}
	}

	// looking at the matrices for correctness
	fmt.Println("X:")
	matPrint(X)
	fmt.Println("y:")
	matPrint(y)

	// n will have the 3, the number of independent variables
	// including the constant
	_, n := X.Dims()

	// initializing a QR matrix
	qr := new(mat.QR)

	// decomposing the X matrix
	qr.Factorize(X)

	// accessing the Q and R matrices
	q := qr.QTo(nil)
	r := qr.RTo(nil)

	// transposing Q matrix
	qt := q.T()

	// calculating Q.T*y and storing it in qty
	qty := new(mat.Dense)
	qty.Mul(qt, y)

	// initializing c to store constants
	c := make([]float64, n)

	// using back-substitution calculating all constants
	for i := n - 1; i >= 0; i-- {
		c[i] = qty.At(i, 0)
		for j := i + 1; j < n; j++ {
			c[i] -= c[j] * r.At(i, j)
		}
		c[i] /= r.At(i, i)
	}

	// printing the constants
	fmt.Printf("Printing the calculated constants:\nb0: %.2f, b1: %.2f, b2: %.2f\n", c[0], c[1], c[2])
}
