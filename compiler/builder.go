package compiler

import (
	"github.com/ritterhou/yui/common"
	"log"
	"strconv"
)

// 将指令序列化为二进制的字节码
func serialize(instructions []instruct) []byte {
	byteCode := common.Magic
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
	// 预处理
	source := preProcess(string(input))
	if source == "" {
		// 没有任何指令则只返回一个头部
		return common.Magic
	}

	// 词法分析
	tokens := listTokens(source)
	// 语法分析
	ast := parse(tokens)
	// 语义分析
	ir := analysis(ast)

	return serialize(ir)
}
