package compiler

import (
	"strings"
)

// 保存定义的变量
var variables = make(map[string]string)

func removeComments(source string) string {
	if !strings.HasSuffix(source, "\n") {
		source += "\n"
	}
	return commentReg.ReplaceAllString(source, "")
}

func removeStopWords(source string) string {
	for _, word := range stopWords {
		source = strings.Replace(source, word, "", -1)
	}
	return source
}

func getDefines(source string) string {
	defines := defineReg.FindAllStringSubmatch(source, -1)
	if defines != nil {
		// 移除源代码中的定义代码
		source = defineReg.ReplaceAllString(source, "")
	}
	for _, define := range defines {
		name := define[2]
		value := define[3]
		variables[name] = value
	}
	return source
}

func setDefines(source string) string {
	for key, value := range variables {
		source = strings.Replace(source, key, value, -1)
	}
	return source
}

func splitExpressions(source string) []string {
	// 没有任何表达式
	if source == "" {
		return []string{}
	}

	results := exprReg.FindAllStringSubmatch(source, -1)
	if results == nil {
		// 没能解析出结果则认为只有一个表达式
		return []string{source}
	}

	expressions := make([]string, 0)
	for _, result := range results {
		expressions = append(expressions, result[1])
	}
	return expressions
}

// 预处理
func preProcess(source string) []string {
	source = removeComments(source)
	source = getDefines(source)
	source = setDefines(source)
	source = removeStopWords(source)
	expressions := splitExpressions(source)
	return expressions
}
