package main

import (
	"bufio"
	"fmt"
	"github.com/ritterhou/yui/common"
	"github.com/ritterhou/yui/compiler"
	"github.com/ritterhou/yui/vm"
	"github.com/ritterhou/yui/ylog"
	"log"
	"os"
	"strings"
	"time"
)

// 编译源代码文件
func build(filename string) {
	sourceFile := common.GetAbsPath(filename)
	data := common.ReadFile(sourceFile)
	target := compiler.Build(data)

	filename = strings.Split(filename, ".")[0] + common.FileExtension
	targetFile := common.GetAbsPath(filename)
	common.WriteFile(targetFile, target)
}

// 运行编译文件或源代码文件
func run(filename string) {
	filename = common.GetAbsPath(filename)
	data := common.ReadFile(filename)
	// 如果输入的是源文件，则需要先对源文件进行编译
	if !common.IsCompiled(data) {
		data = compiler.Build(data)
	}
	results := vm.Run(data)
	for _, result := range results {
		fmt.Println(result)
	}
}

var order uint32

func getOrderAndTime() string {
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	order++
	return fmt.Sprintf(" %2d  %s  ", order, now)
}

// 进入交互式的shell
func shell() {
	stopWords := []string{"\n", "\r"}
	histories := make([]string, 0)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Yui shell " + common.Version)
	for {
		fmt.Print(">>> ")
		expr, _ := reader.ReadString('\n')
		for _, word := range stopWords {
			expr = strings.Replace(expr, word, "", -1)
		}
		if expr == "quit" || expr == "exit" {
			break
		}
		if expr == "history" {
			for _, history := range histories {
				fmt.Println(history)
			}
			continue
		}
		histories = append(histories, getOrderAndTime()+expr)

		byteCode := compiler.Build([]byte(expr))
		results := vm.Run(byteCode)
		for _, result := range results {
			fmt.Println(result)
		}
	}
}

// 反编译字节码并生成指令
func decompile(filename string) {
	filename = common.GetAbsPath(filename)
	data := common.ReadFile(filename)
	if !common.IsCompiled(data) {
		log.Fatal("error data format: not yuicode file")
	}
	instructions := compiler.Decompile(data)
	for _, ins := range instructions {
		if !(ins == common.Begin || ins == common.End) {
			fmt.Print("  ")
		}
		fmt.Println(ins)
	}
}

func main() {
	// 从参数中获取文件名
	getFilename := func(args []string) (string, bool) {
		if len(args) <= 2 {
			return "", false
		}
		return args[2], true
	}

	ylog.Init()

	args := os.Args
	if len(args) == 1 { // 没有参数
		fmt.Print(common.Help)
		return
	}

	command := args[1]
	switch command {
	case "build":
		if filename, exist := getFilename(args); exist {
			build(filename)
		} else {
			fmt.Println("build: no input file")
		}
	case "run":
		if filename, exist := getFilename(args); exist {
			run(filename)
		} else {
			fmt.Println("run: no input file")
		}
	case "shell":
		shell()
	case "dec":
		if filename, exist := getFilename(args); exist {
			decompile(filename)
		} else {
			fmt.Println("decompile: no input file")
		}
	default:
		fmt.Printf("unknown command \"%s\"\n", command)
	}

}
