package calculator

type Engine struct{}

type Operation struct {
	Expression string
	Operator   string
	Operands   []float64
}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

func (e *Engine) Subtract(x, y float64) float64 {
	return x - y
}

func (e *Engine) Multiply(x, y float64) float64 {
	return x * y
}

func (e *Engine) Divide(x, y float64) float64 {
	return x / y
}
