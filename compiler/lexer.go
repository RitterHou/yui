package compiler

import (
	"log"
	"strings"
)

// 移除注释行
func removeComments(source string) string {
	if !strings.HasSuffix(source, "\n") {
		source += "\n"
	}
	return commentReg.ReplaceAllString(source, "")
}

// 词法分析
func listTokens(source string) []string {
	source = removeComments(source)
	for _, word := range stopWords {
		source = strings.Replace(source, word, "", -1)
	}

	length := len(source)
	tokens := make([]string, 0, length+2)
	tokens = append(tokens, exprStart)
	index := 0
	for {
		if index == length {
			break
		}
		// 操作符
		word := string(rune(source[index])) // ascii码转化为字符
		if !(word == exprStart || word == exprEnd || word == plus || word == minus || word == multiply || word == divide) {
			// 浮点数
			floatParam := floatReg.FindStringSubmatch(source[index:])
			if len(floatParam) > 0 {
				word = floatParam[0]
			} else {
				// 整数
				intParam := intReg.FindStringSubmatch(source[index:])
				if len(intParam) > 0 {
					word = intParam[0]
				} else {
					log.Fatalf("invalid word \"%s\"", word)
				}
			}
		}
		index += len(word)
		tokens = append(tokens, word)
	}
	tokens = append(tokens, exprEnd)
	return tokens
}
