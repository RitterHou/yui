package vm

import (
	"fmt"
	"log"
	"strings"
)

func newStack() *stack {
	return &stack{value: make([]float32, 0)}
}

type stack struct {
	value []float32
}

func (s *stack) push(v float32) {
	s.value = append(s.value, v)
}

func (s *stack) pop() float32 {
	value := s.value
	length := len(value)
	if length == 0 {
		log.Fatal("stack value is nil, can't pop anything")
	}
	top := value[length-1]
	s.value = value[:length-1]
	return top
}

func (s stack) String() string {
	value := s.value
	strVal := make([]string, len(value))
	for i := 0; i < len(value); i++ {
		strVal[i] = fmt.Sprintf("%.6f", value[i])
	}
	result := "["
	result += strings.Join(strVal, ", ") + "]"
	return result
}
