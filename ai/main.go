package main

import (
	"fmt"
)

// Linear regression model
type LinearRegression struct {
	slope     float64
	intercept float64
}

// Fit the linear regression model to the data
func (lr *LinearRegression) Fit(x []float64, y []float64) {
	n := float64(len(x))
	sumX, sumY, sumXY, sumXSquare := 0.0, 0.0, 0.0, 0.0
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}
	lr.slope = (n*sumXY - sumX*sumY) / (n*sumXSquare - sumX*sumX)
	lr.intercept = (sumY - lr.slope*sumX) / n
}

// Predict the output based on input x
func (lr *LinearRegression) Predict(x float64) float64 {
	return lr.slope*x + lr.intercept
}

func main() {
	// Sample data
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	// Create and fit the linear regression model
	var lr LinearRegression
	lr.Fit(x, y)

	// Predict output for a new input
	newInput := 7.0
	prediction := lr.Predict(newInput)
	fmt.Printf("Predicted output for input %.2f: %.2f\n", newInput, prediction)
}
