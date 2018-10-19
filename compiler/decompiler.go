package compiler

import (
	"fmt"
	"github.com/ritterhou/yui/common"
)

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

// 将二进制字节码转化为伪码
func unSerialize(input []byte) []string {
	ins := make([]string, 0)

	index := 0
	length := len(input)
	for index < length {
		op := input[index]
		index++
		switch op {
		case common.PUSH:
			numType := input[index]
			index++
			buf := input[index : index+4]
			index += 4

			if numType == common.INT {
				num := common.ByteArray2Int(buf)
				ins = append(ins, "PUSH "+fmt.Sprintf("%d", num))
			} else {
				num := common.ByteArray2Float(buf)
				ins = append(ins, "PUSH "+fmt.Sprintf("%v", num))
			}
		case common.PLUS:
			ins = append(ins, "PLUS")
		case common.MINUS:
			ins = append(ins, "MINUS")
		case common.MULTIPLY:
			ins = append(ins, "MULTIPLY")
		case common.DIVIDE:
			ins = append(ins, "DIVIDE")
		}
	}
	return ins
}

func Decompile(input []byte) []string {
	ins := make([]string, 0)
	expressions := getExpressions(input)

	for _, expr := range expressions {
		ins = append(ins, common.Begin)
		ins = append(ins, unSerialize(expr)...)
		ins = append(ins, common.End)
	}
	return ins
}
