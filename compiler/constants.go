package compiler

import "regexp"

var (
	stopWords = []string{"\n", "\r", "\t", " "}
	floatReg  = regexp.MustCompile("(^[\\d]+\\.[\\d]+)")
	intReg    = regexp.MustCompile("(^[\\d]+)")
)

const (
	exprStart = "("
	exprEnd   = ")"
	plus      = "+"
	minus     = "-"
	multiply  = "*"
	divide    = "/"
)
