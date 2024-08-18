package rgx

import (
	"fmt"
	"strings"
)

func Match(expr, str string) (bool, int, int) {
	if len(expr) == 0 || len(str) == 0 {
		return false, 0, 0
	}

	matched := false // TODO: do we need to explicitly initialize?
	matchPos := 0    // main cursor on str
	// maxMatchPos := 0 // TODO: shouldn't always be (len(string) - 1)?
	maxMatchPos := len(str) - 1
	matchLen := 0
	if isStart(expr[0]) {
		expr = expr[1:]
	}
	// else{
	//     maxMatchPos = len(string) - 1 }// TODO: why

	for matchPos <= maxMatchPos {
		matched, matchLen = matchExpr(expr, str[matchPos:], 0)
		if matched {
			// return matched, matchPos, matchLen
			break
		}
		matchPos += 1
	}

	return matched, matchPos, matchLen
}

func matchExpr(expr string, str string, matchLen int) (bool, int) {
	if len(expr) == 0 {
		return true, matchLen
	} else if isEnd(expr[0]) {
		// if len(str) == 0 {
		// 	return true, matchLen
		// } else {
		// 	return false, 0
		// }
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
		fmt.Printf("Unknown token in expr %s.", expr)
	}
	return false, 0
}

// matchers //
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
	// if !minMatchLen{
	//     minMatchLen = 0
	// }
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

	for submatchLen >= minMatchLen {
		matched, newMatchLen := matchExpr(
			strings.Repeat(head, submatchLen)+rest, str, matchLen,
		)
		if matched {
			return matched, newMatchLen
		}
		submatchLen -= 1
	}
	return false, 0
}

func matchAlternate(expr string, str string, matchLen int) (bool, int) {
	head, _, rest := splitExpr(expr)
	// options = splitAlternate(head)
	options := strings.Split(head[1:len(head)-1], "|")

	for _, option := range options {
		matched, minMatchLen := matchExpr(option+rest, str, matchLen)
		//TODO: refactor with break and match/minMLen declared outside of loop
		if matched {
			return matched, minMatchLen
		}
	}
	return false, 0
}
