package main

import (
	"bufio"
	"fmt"
	"github.com/ritterhou/yui/common"
	"github.com/ritterhou/yui/compiler"
	"github.com/ritterhou/yui/vm"
	"os"
	"strings"
	"time"
)

var desc = "Welcome to Yui shell " + common.Version + ".\nType in expressions for evaluation."
var stopWords = []string{"\n", "\r"}
var histories = make([]string, 0)
var reader = bufio.NewReader(os.Stdin)
var order uint32

func getOrderAndTime() string {
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	order++
	return fmt.Sprintf(" %2d  %s  ", order, now)
}

func command() {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Printf("\033[0;31m%s\033[0m\n", err)
		}
	}()
	fmt.Print(">>> ")
	expr, _ := reader.ReadString('\n')
	for _, word := range stopWords {
		expr = strings.Replace(expr, word, "", -1)
	}
	if expr == "quit" || expr == "exit" {
		os.Exit(0)
	}
	if expr == "history" {
		for _, history := range histories {
			fmt.Println(history)
		}
		return
	}
	if expr == "" {
		return
	}
	histories = append(histories, getOrderAndTime()+expr)
	byteCode := compiler.Build([]byte(expr))
	results := vm.Run(byteCode)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 进入交互式的shell
func shell() {
	fmt.Println(desc)
	for {
		command()
	}
}
