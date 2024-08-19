package rgx

import (
	"fmt"
	"strings"
)

func Match(expr, str string) (bool, int, int) {
	if len(expr) == 0 || len(str) == 0 {
		return false, 0, 0
	}

	matched := false
	matchPos := -1 // main cursor on str
	matchLen := 0
	maxMatchPos := len(str) - 1
	if isStart(expr[0]) {
		expr = expr[1:]
		maxMatchPos = 0 // in this case the cursor shouldn't go beyond the first char without matching
	}

	for !matched && matchPos <= maxMatchPos {
		matchPos += 1
		matched, matchLen = matchExpr(expr, str[matchPos:], 0)
	}
	return matched, matchPos, matchLen
}

func matchExpr(expr string, str string, matchLen int) (bool, int) {
	if len(expr) == 0 {
		return true, matchLen
	} else if isEnd(expr[0]) {
		return len(str) == 0, matchLen
	}

	head, operator, rest := splitExpr(expr)

	if isStar(operator) {
		return matchStar(expr, str, matchLen)
	} else if isPlus(operator) {
		return matchPlus(expr, str, matchLen)
	} else if isQuestion(operator) {
		return matchQuestion(expr, str, matchLen)
	} else if isAlternate(head) {
		return matchAlternate(expr, str, matchLen)
	} else if isUnit(head) {
		if doesUnitMatch(expr, str) {
			return matchExpr(rest, str[1:], matchLen+1)
		}
	} else {
		// TODO: or return err?
		panic(fmt.Sprintf("unknown token in expr %s", expr))
	}
	return false, 0
}

func matchStar(expr string, str string, matchLen int) (bool, int) {
	return matchMultiple(expr, str, matchLen, 0, 0)
}

func matchPlus(expr string, str string, matchLen int) (bool, int) {
	return matchMultiple(expr, str, matchLen, 1, 0)
}

func matchQuestion(expr string, str string, matchLen int) (bool, int) {
	return matchMultiple(expr, str, matchLen, 0, 1)
}

func matchMultiple(expr string, str string, matchLen int, minMatchLen int, maxMatchLen int) (bool, int) {
	head, _, rest := splitExpr(expr)
	submatchLen := -1
	// TODO: check in which case (maxMatchLen == 0) is used
	for maxMatchLen == 0 || (submatchLen < maxMatchLen) {
		subexprMatched, _ := matchExpr(
			strings.Repeat(head, submatchLen+1), str, matchLen,
		)
		if subexprMatched {
			submatchLen += 1
		} else {
			break
		}
	}

	matched := false
	newMatchLen := 0
	for submatchLen >= minMatchLen {
		matched, newMatchLen = matchExpr(
			strings.Repeat(head, submatchLen)+rest, str, matchLen,
		)
		if matched {
			break
		}
		submatchLen -= 1
	}
	return matched, newMatchLen
}

func matchAlternate(expr string, str string, matchLen int) (bool, int) {
	head, _, rest := splitExpr(expr)
	options := splitAlternate(head)

	matched := false
	minMatchLen := 0
	for _, option := range options {
		matched, minMatchLen = matchExpr(option+rest, str, matchLen)
		if matched {
			break
		}
	}
	return matched, minMatchLen
}
