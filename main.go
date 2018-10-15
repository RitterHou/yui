package main

import (
	"bufio"
	"fmt"
	"github.com/ritterhou/yui/common"
	"github.com/ritterhou/yui/compiler"
	"github.com/ritterhou/yui/ylog"
	"log"
	"os"
	"strings"
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
	log.Println("run " + filename)
}

// 进入交互式的shell
func shell() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Yui shell " + common.Version)
	for {
		fmt.Print("> ")
		expr, _ := reader.ReadString('\n')
		expr = strings.Trim(expr, "\n")
		if expr == "quit" || expr == "exit" {
			break
		}
		log.Print(expr)
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
	default:
		fmt.Printf("unknown command \"%s\"\n", command)
	}

}
