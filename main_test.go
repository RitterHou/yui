package main

import (
	"github.com/ritterhou/yui/compiler"
	"github.com/ritterhou/yui/vm"
	"testing"
)

func TestBuildAndRun(t *testing.T) {
	sources := []string{
		`// 计算圆的面积
		 define PI 3.1415926
		 define radius 10
		 (radius * radius) * PI // 314.15926`,

		`define name 2333
		 define age 18
		 define salary 3024.85
		 name + age + salary // 5375.85`,

		`define _ 666`,

		`_ + 334`,
		`{_ + 334} {age} {name + age + salary}`}
	for _, source := range sources {
		byteCode := compiler.Build([]byte(source))
		results := vm.Run(byteCode)
		for _, result := range results {
			t.Log(result)
		}
	}
}
