package compiler

import (
	"fmt"
	"github.com/ritterhou/yui/common"
	"log"
)

func prettyPrint(ast *node) {
	// 首先定义一个类型为function的变量
	var prettyPrint func(node1 *node, indent string)
	// 实现function
	prettyPrint = func(node1 *node, indent string) {
		fmt.Println(indent + " " + node1.value)
		if node1.left != nil {
			prettyPrint(node1.left, indent+"──")
		}
		if node1.right != nil {
			prettyPrint(node1.right, indent+"──")
		}
	}
	// 调用function
	prettyPrint(ast, "├──")
}

type instruct struct {
	op    byte // 操作类型
	value string
}

// 指令
var instructions []instruct

func addIns(ins instruct) {
	instructions = append(instructions, ins)
}

// 遍历语法树生成指令数组
func traverse(node1 *node) {
	left := node1.left
	right := node1.right
	value := node1.value
	// 叶子节点保存数据
	if left == nil && right == nil {
		ins := instruct{op: common.PUSH, value: value}
		addIns(ins)
		return
	}
	// 继续遍历内部节点的子节点
	traverse(left)
	traverse(right)
	// 内部节点保存了操作符
	switch value {
	case plus:
		addIns(instruct{op: common.PLUS})
	case minus:
		addIns(instruct{op: common.MINUS})
	case multiply:
		addIns(instruct{op: common.MULTIPLY})
	case divide:
		addIns(instruct{op: common.DIVIDE})
	default:
		log.Panicf("unkonwn node %s", value)
	}
}

// 语义分析
func analysis(ast *node) []instruct {
	instructions = make([]instruct, 0, 10)
	traverse(ast)
	return instructions
}
