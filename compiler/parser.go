package compiler

type node struct {
	left  *node
	right *node
	value string
}

var index int
var tokens []string
var length int

// 计算表达式块，包含括号内的表达式或数值
func factor() *node {
	token := tokens[index]
	index++
	switch token {
	case exprStart:
		expression := expr()
		index++ // 移除最后的右括号
		return expression
	case plus:
		value := tokens[index]
		index++
		return &node{value: value}
	case minus:
		value := minus + tokens[index]
		index++
		return &node{value: value}
	default:
		return &node{value: token}
	}
}

// 计算乘除法
func term() *node {
	factor1 := factor()
	if index == length {
		return factor1
	}
	op := tokens[index]
	for op == multiply || op == divide {
		index++
		factor2 := factor()
		factor1 = &node{value: op, left: factor1, right: factor2}
		op = tokens[index]
	}
	return factor1
}

// 计算加减法
func expr() *node {
	term1 := term()
	if index == length {
		return term1
	}
	op := tokens[index]
	for op == plus || op == minus {
		index++
		term2 := term()
		term1 = &node{value: op, left: term1, right: term2}
		op = tokens[index]
	}
	return term1
}

// 语法分析
func parse(input []string) *node {
	index = 0
	tokens = input
	length = len(tokens)
	return expr()
}
