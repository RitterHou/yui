package compiler

import (
	"fmt"
	"github.com/ritterhou/yui/common"
)

// 将二进制字节码转化为伪码
func Decompile(input []byte) []string {
	ins := make([]string, 0)

	index := 4
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
				ins = append(ins, "PUSH "+fmt.Sprintf("%.10f", num))
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
