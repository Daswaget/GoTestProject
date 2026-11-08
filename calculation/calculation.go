package Calculation

import operations "github.com/Daswaget/GoTestProject/operations"

type Calculator interface {
	Calculate(x, y float64) float64
}

type AdditionCalculator struct{}

func (c AdditionCalculator) Calculate(x, y float64) float64 {
	return operations.Addition(x, y)
}

type SubtractionCalculator struct{}

func (c SubtractionCalculator) Calculate(x, y float64) float64 {
	return operations.Subtraction(x, y)
}

type MultiplicationCalculator struct{}

func (c MultiplicationCalculator) Calculate(x, y float64) float64 {
	return operations.Multiplication(x, y)
}

type DivisionCalculator struct{}

func (c DivisionCalculator) Calculate(x, y float64) float64 {
	return operations.Division(x, y)
}
