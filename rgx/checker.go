package rgx

import (
	"strings"
	"unicode"
)

func isStart(b byte) bool {
	return b == '^'
}

func isEnd(b byte) bool {
	return b == '$'
}

func isOpenSet(b byte) bool {
	return b == '['
}

func isCloseSet(b byte) bool {
	return b == ']'
}

func isSet(term string) bool {
	return isOpenSet(term[0]) && isCloseSet(term[len(term)-1])
}

func isOpenAlternate(b byte) bool {
	return b == '('
}

func isCloseAlternate(b byte) bool {
	return b == ')'
}

func isAlternate(term string) bool {
	return isOpenAlternate(term[0]) && isCloseAlternate(term[len(term)-1])
}

func isEscape(b byte) bool {
	return b == '\\'
}

func isEscapeSequence(term string) bool {
	return isEscape(term[0])
}

func isOperator(b byte) bool {
	return isStar(b) || isPlus(b) || isQuestion(b)
}

func isStar(b byte) bool {
	return b == '*'
}

func isPlus(b byte) bool {
	return b == '+'
}

func isQuestion(b byte) bool {
	return b == '?'
}

func isLiteral(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b)) || isNonWordSymbol(b)
	// return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	// || (b == ' ' || b == ':' || b == '/')
}

func isNonWordSymbol(b byte) bool {
	return b == ' ' || b == ':' || b == '/'
}

func isDot(b byte) bool {
	return b == '.'
}

func isUnit(term string) bool {
	return isLiteral(term[0]) || isDot(term[0]) || isSet(term) || isEscapeSequence(term)
}

func doesUnitMatch(expr string, str string) bool {
	head, _, _ := splitExpr(expr)

	if len(str) == 0 {
		return false
	}

	// TODO: head is only one char?
	if isLiteral(head[0]) {
		return expr[0] == str[0]
	} else if isDot(head[0]) {
		return true
	} else if isEscapeSequence(head) {
		// TODO: fix \\a
		if head == "\\a" {
			return unicode.IsLetter(rune(str[0]))
		} else if head == "\\d" {
			return unicode.IsDigit(rune(str[0]))
		} else {
			return false
		}
	} else if isSet(head) {
		setTerms := splitSet(head)
		// setTerms := head[1 : len(head)-1]
		return strings.ContainsRune(setTerms, rune(str[0]))
	}
	return false
}
