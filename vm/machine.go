package vm

import (
	"errors"
	"github.com/ritterhou/yui/common"
)

type instruct struct {
	op    byte
	value float32
}

func getInstructs(input []byte) []instruct {
	ins := make([]instruct, 0)

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

func Run(input []byte) (float32, error) {
	instructs := getInstructs(input)
	if len(instructs) == 0 {
		return 0, errors.New("no instructions run")
	}
	s := newStack()
	for _, ins := range instructs {
		switch ins.op {
		case common.PUSH:
			s.push(ins.value)
		case common.PLUS:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v1 + v2)
		case common.MINUS:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 - v1)
		case common.MULTIPLY:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v1 * v2)
		case common.DIVIDE:
			v1 := s.pop()
			v2 := s.pop()
			s.push(v2 / v1)
		}
	}
	return s.pop(), nil
}
