package input

import "calculator-project/calculator"

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	return nil, nil
}
