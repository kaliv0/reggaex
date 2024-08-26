package rgx

import (
	"strings"
)

func splitExpr(expr string) (string, string, string) {
	var head string
	var operator string
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
		head = expr[:lastExprPos]
	} else {
		lastExprPos = 1
		head = string(expr[0])
	}

	if lastExprPos < len(expr) && isOperator(expr[lastExprPos]) {
		operator = string(expr[lastExprPos])
		lastExprPos += 1
	}

	if lastExprPos < len(expr) && isOpenQuantifier(expr[lastExprPos]) {
		closingQntPos := strings.IndexByte(expr, '}')
		operator = expr[lastExprPos+1 : closingQntPos]
		lastExprPos = closingQntPos + 1
		//validate quantifier for early failure
		validateQuantifier(operator)
	}

	rest = expr[lastExprPos:]
	return head, operator, rest
}

func splitSet(head string) string {
	if len(head) > 0 && head[0] == '[' && head[len(head)-1] == ']' {
		return head[1 : len(head)-1]
	}
	return head
}

func splitAlternate(head string) []string {
	return strings.Split(head[1:len(head)-1], "|")
}
