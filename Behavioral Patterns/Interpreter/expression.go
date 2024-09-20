package main

type Expression interface {
	Interpret() int
}

type NumberExpression struct {
	value int
}

func (n *NumberExpression) Interpret() int {
	return n.value
}

type AddExpression struct {
	left  Expression
	right Expression
}

func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

type MultiplyExpression struct {
	left  Expression
	right Expression
}

func (m *MultiplyExpression) Interpret() int {
	return m.left.Interpret() * m.right.Interpret()
}
