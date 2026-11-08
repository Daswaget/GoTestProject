package calculation

type Calculator interface {
	Calculate(x, y float64) float64
}

type AdditionCalculator struct{}

func (c AdditionCalculator) Calculate(x, y float64) float64 {
	return x + y
}

type SubtractionCalculator struct{}

func (c SubtractionCalculator) Calculate(x, y float64) float64 {
	return x - y
}

type MultiplicationCalculator struct{}

func (c MultiplicationCalculator) Calculate(x, y float64) float64 {
	return x * y
}

type DivisionCalculator struct{}

func (c DivisionCalculator) Calculate(x, y float64) float64 {
	return x / y
}
