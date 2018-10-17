package compiler

import "fmt"

func prettyPrint(node1 *node, indent string) {
	fmt.Println(indent + " " + node1.value)
	if node1.left != nil {
		prettyPrint(node1.left, indent+"──")
	}
	if node1.right != nil {
		prettyPrint(node1.right, indent+"──")
	}
}

// 语义分析
func analysis(ast *node) []byte {
	prettyPrint(ast, "├──")
	return make([]byte, 0, 0)
}
