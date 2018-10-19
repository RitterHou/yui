package compiler

import "regexp"

var (
	stopWords  = []string{"\n", "\r", "\t", " "}
	floatReg   = regexp.MustCompile("(^[\\d]+\\.[\\d]+)")
	intReg     = regexp.MustCompile("(^[\\d]+)")
	commentReg = regexp.MustCompile("(//.*\n)")
	defineReg  = regexp.MustCompile(
		"(define +([A-Za-z_][A-Za-z0-9_]*) +([\\d]+\\.[\\d]+|[\\d]+))")
)

const (
	exprStart = "("
	exprEnd   = ")"
	plus      = "+"
	minus     = "-"
	multiply  = "*"
	divide    = "/"
)
