package compiler

import "testing"

func TestPreProcess(t *testing.T) {
	source := `{1 + 1} {2 + 2}`
	expressions := preProcess(source)
	for _, expr := range expressions {
		t.Log(expr)
	}
}
