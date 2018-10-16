package compiler

import (
	"log"
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
	for _, word := range stopWords {
		source = strings.Replace(source, word, "", -1)
	}

	length := len(source)
	tokens := make([]string, 0, length+2)
	tokens = append(tokens, exprStart)
	index := 0
	for {
		if index == length {
			break
		}
		// 操作符
		word := string(rune(source[index])) // ascii码转化为字符
		if !(word == exprStart || word == exprEnd || word == plus || word == minus || word == multiply || word == divide) {
			// 浮点数
			floatParam := floatReg.FindStringSubmatch(source[index:])
			if len(floatParam) > 0 {
				word = floatParam[0]
			} else {
				// 整数
				intParam := intReg.FindStringSubmatch(source[index:])
				if len(intParam) > 0 {
					word = intParam[0]
				} else {
					log.Fatalf("invalid word \"%s\"", word)
				}
			}
		}
		index += len(word)
		tokens = append(tokens, word)
	}
	tokens = append(tokens, exprEnd)
	return tokens
}

// 语法分析
func parser(tokens []string) *node {
	// 给父节点增加一个子节点
	addChild := func(parent *node, child *node) {
		children := parent.children
		children = append(children, child)
		parent.children = children
	}
	root := new(node)
	parent := new(node)

	for _, token := range tokens {
		log.Println(token)
		newNode := &node{value: token}
		if root.value == "" {
			root = newNode
		} else {
			addChild(parent, newNode)
		}

		switch token {
		case exprStart:
			parent = newNode
		case exprEnd:
			parent = parent.parent
		case plus:
		case minus:
		case multiply:
		case divide:
		default:

		}
	}
	return root
}
