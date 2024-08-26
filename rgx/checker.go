package rgx

import (
	"fmt"
	"strconv"
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
	return isStar(string(b)) || isPlus(string(b)) || isQuestion(string(b))
}

func isStar(b string) bool {
	return b == "*"
}

func isPlus(b string) bool {
	return b == "+"
}

func isQuestion(b string) bool {
	return b == "?"
}

func isOpenQuantifier(b byte) bool {
	return b == '{'
}

func isCloseQuantifier(b byte) bool {
	return b == '}'
}

func isQuantifier(term string) (bool, int) {
	val, err := strconv.Atoi(term)
	if err != nil {
		return false, 0
	}
	return true, val
}

func isLiteral(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b)) || isNonWordSymbol(b)
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
	if len(str) == 0 {
		return false
	}

	head, _, _ := splitExpr(expr)

	if isLiteral(head[0]) {
		return expr[0] == str[0]
	} else if isDot(head[0]) {
		return true
	} else if isEscapeSequence(head) {
		return validateEscapeSequence(head, str)
	} else if isSet(head) {
		return doesSetMatch(head, str, true)
	}
	return false
}

func doesSetMatch(head string, str string, flag bool) bool {
	setTerms := splitSet(head)
	curr := rune(str[0])
	flag = true
	if setTerms[0] == '^' {
		flag = false
	}
	if idx := strings.IndexByte(setTerms, '-'); idx >= 0 {
		return doesRangeMatch(setTerms, str, idx, curr, flag)
	}
	return strings.ContainsRune(setTerms, curr) == flag
}

func doesRangeMatch(setTerms string, str string, idx int, curr rune, flag bool) bool {
	rangeStart := rune(setTerms[idx-1])
	rangeEnd := rune(setTerms[idx+1])
	rest := ""
	if len(setTerms) > idx+1 {
		rest = setTerms[idx+2:]
	}

	result := (curr >= rangeStart && curr <= rangeEnd) == flag
	if !result && len(rest) > 0 {
		return doesSetMatch(rest, str, flag)
	} else {
		return result
	}
}

func validateEscapeSequence(head string, str string) bool {
	if head == "\\w" {
		return unicode.IsLetter(rune(str[0])) || unicode.IsDigit(rune(str[0])) || str[0] == '_'
	} else if head == "\\d" {
		return unicode.IsDigit(rune(str[0]))
	} else if head == "\\s" {
		return unicode.IsSpace(rune(str[0]))
	} else if head == "\\W" {
		return !unicode.IsLetter(rune(str[0])) && !unicode.IsDigit(rune(str[0])) && str[0] != '_'
	} else if head == "\\D" {
		return !unicode.IsDigit(rune(str[0]))
	} else if head == "\\S" {
		return !unicode.IsSpace(rune(str[0]))
	} else {
		return false
	}
}

func validateQuantifier(operator string) {
	for _, c := range operator {
		if !unicode.IsDigit(c) {
			panic(fmt.Sprintf("supplied value %s is not a number\n", operator))
		}
	}
}
