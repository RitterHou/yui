package compiler

import (
	"strings"
)

type node struct {
	parent   *node
	children []*node
	value    string
}

func Build(source []byte) (target []byte) {
	tokens := lexer(string(source))
	parser(tokens)
	return source
}

// 词法分析
func lexer(source string) []string {
	for _, word := range StopWords {
		source = strings.Replace(source, word, "", -1)
	}
	tokens := strings.Split(source, " ")

	tokens1 := make([]string, 0, len(tokens)+2)
	tokens1 = append(tokens1, ExprStart)
	for _, token := range tokens {
		if token != "" {
			tokens1 = append(tokens1, token)
		}
	}
	tokens1 = append(tokens1, ExprEnd)
	return tokens1
}

// 语法分析
func parser(tokens []string) node {
	addChild := func(parent *node, child node) {

	}
	var root node
	var parent node

	for _, token := range tokens {
		newNode := node{value: token}
		addChild(&parent, newNode)

		switch token {
		case ExprStart:
		case ExprEnd:
		case Plus:
		case Minus:
		case Multiply:
		case Divide:
		default:

		}
	}
	return root
}
