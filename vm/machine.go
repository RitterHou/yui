package vm

import (
	"github.com/ritterhou/yui/common"
)

type instruct struct {
	op    byte
	value float32
}

// 将表达式的字节码转换为指令列表
func getInstructs(input []byte) []instruct {
	ins := make([]instruct, 0)

	index := 0
	length := len(input)
	for index < length {
		op := input[index]
		index++
		switch op {
		case common.PUSH:
			numType := input[index] // 数字的类型（整数或浮点数）
			index++
			buf := input[index : index+4] // 数字的值，长度为4个字节
			index += 4

			var num float32
			if numType == common.INT {
				num = float32(common.ByteArray2Int(buf))
			} else {
				num = common.ByteArray2Float(buf)
			}
			ins = append(ins, instruct{op: op, value: num})
		default:
			ins = append(ins, instruct{op: op})
		}
	}
	return ins
}

// 从原始的字节码中获取表达式字节数组列表
func getExpressions(input []byte) [][]byte {
	expressions := make([][]byte, 0)

	index := 4 // 去掉魔数
	length := len(input)
	for index < length {
		numByte := input[index : index+4]
		index += 4
		// 表达式本身的字节长度
		num := int(common.ByteArray2Int(numByte))

		expr := input[index : index+num]
		index += num
		expressions = append(expressions, expr)
	}
	return expressions
}

// 执行表达式的指令
func executeExpr(expr []byte) float32 {
	instructs := getInstructs(expr)

	s := newStack()
	for _, ins := range instructs {
		switch ins.op {
		case common.PUSH:
			s.push(ins.value)
		case common.PLUS:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 + v1)
		case common.MINUS:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 - v1)
		case common.MULTIPLY:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 * v1)
		case common.DIVIDE:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 / v1)
		}
	}
	return s.pop()
}

// 虚拟机执行字节码
func Run(input []byte) []float32 {
	expressions := getExpressions(input)

	results := make([]float32, 0)
	for _, expr := range expressions {
		result := executeExpr(expr)
		results = append(results, result)
	}
	return results
}
