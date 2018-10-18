package vm

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := newStack()
	s.push(2.333)
	s.push(666)
	s.push(8888)
	s.push(000000)
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s)
}
