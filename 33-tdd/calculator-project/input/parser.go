package input

import (
	"calculator-project/calculator"
)

type OperationProcessor interface {
	ProcessOperation(operation *calculator.Operation) (*string, error)
}

type ValidationHelper interface {
	CheckInput(operator string, operands []float64) error
}

type Parser struct {
	engine    OperationProcessor
	validator ValidationHelper
}
