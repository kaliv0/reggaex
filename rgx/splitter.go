package rgx

import "strings"

func splitExpr(expr string) (string, byte, string) {
	var head string
	var operator byte
	var rest string
	var lastExprPos int

	if isOpenSet(expr[0]) {
		lastExprPos = strings.IndexByte(expr, ']') + 1
		head = expr[:lastExprPos]
	} else if isOpenAlternate(expr[0]) {
		lastExprPos = strings.IndexByte(expr, ')') + 1
		head = expr[:lastExprPos]
	} else if isEscape(expr[0]) {
		lastExprPos += 2
		// head = expr[:2]
		head = expr[:lastExprPos] // TODO: check
	} else {
		lastExprPos = 1
		head = string(expr[0])
	}

	if lastExprPos < len(expr) && isOperator(expr[lastExprPos]) {
		operator = expr[lastExprPos]
		lastExprPos += 1
	}

	rest = expr[lastExprPos:]
	return head, operator, rest
}

func splitSet(setHead string) string {
	return setHead[1 : len(setHead)-1]
}
