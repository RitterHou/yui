package compiler

import (
	"fmt"
	"testing"
)

func TestRemoveComments(t *testing.T) {
	source := ` // this is a comment
				// this is also a comment
				(1 + 1) * 20 // this is another comment
				// and last comment`
	target := removeComments(source)
	fmt.Println(target)
}
