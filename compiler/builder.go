package compiler

import (
	"log"
)

func Build(input []byte) (target []byte) {
	source := string(input)

	tokens := listTokens(source)
	log.Println(tokens)
	ast := parse(tokens)
	target = analysis(ast)
	return
}
