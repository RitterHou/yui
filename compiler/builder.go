package compiler

import (
	"github.com/ritterhou/yui/common"
	"log"
	"strconv"
)

// 将指令序列化为二进制的字节码
func serialize(instructions []instruct) []byte {
	byteCode := make([]byte, 0)
	for _, ins := range instructions {
		byteCode = append(byteCode, ins.op)
		value := ins.value
		if value != "" {
			// 整数
			if num, err := strconv.Atoi(value); err == nil {
				byteCode = append(byteCode, common.INT)
				buf := common.Int2ByteArray(int32(num))
				byteCode = append(byteCode, buf...)
				continue
			}
			// 浮点数
			if num, err := strconv.ParseFloat(value, 32); err == nil {
				byteCode = append(byteCode, common.FLOAT)
				buf := common.Float2ByteArray(float32(num))
				byteCode = append(byteCode, buf...)
				continue
			}
			log.Fatalf("value %s is not num", value)
		}
	}
	return byteCode
}

func Build(input []byte) []byte {
	byteCodes := common.Magic
	appendByteCodes := func(byteCode []byte) {
		byteCodes = append(byteCodes, common.Int2ByteArray(int32(len(byteCode)))...)
		byteCodes = append(byteCodes, byteCode...)
	}

	// 预处理
	sources := preProcess(string(input))
	// 每一个source对应着一个计算表达式
	for _, source := range sources {
		// 词法分析
		tokens := listTokens(source)
		// 语法分析
		ast := parse(tokens)
		// 语义分析
		ir := analysis(ast)

		appendByteCodes(serialize(ir))
	}
	return byteCodes
}
