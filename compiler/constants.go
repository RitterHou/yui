package compiler

import "regexp"

var (
	stopWords  = []string{"\n", "\r", "\t", " "}
	floatReg   = regexp.MustCompile("(^[\\d]+\\.[\\d]+)")
	intReg     = regexp.MustCompile("(^[\\d]+)")
	commentReg = regexp.MustCompile("(//.*\n)")
)

const (
	exprStart = "("
	exprEnd   = ")"
	plus      = "+"
	minus     = "-"
	multiply  = "*"
	divide    = "/"
)
