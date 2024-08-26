package rgx

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type MatchData struct {
	Matched  bool
	MatchStr string
}

func Match(expr, str string) (data MatchData, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("something went wrong")
			}
		}
	}()
	return match(expr, str), err
}

func match(expr, str string) MatchData {
	if len(expr) == 0 || len(str) == 0 {
		return MatchData{false, ""}
	}

	matched := false
	matchPos := -1 // main cursor on str
	matchLen := 0
	maxMatchPos := len(str) - 1
	if isStart(expr[0]) {
		expr = expr[1:]
		// in this case the cursor shouldn't go beyond the first char without matching
		maxMatchPos = 0
	}

	for !matched && matchPos <= maxMatchPos {
		matchPos += 1
		matched, matchLen = matchExpr(expr, str[matchPos:], 0)
	}
	if matched {
		return MatchData{matched, str[matchPos : matchPos+matchLen]}
	}
	return MatchData{matched, ""}
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
	}
	if isPlus(operator) {
		return matchPlus(expr, str, matchLen)
	}
	if isQuestion(operator) {
		return matchQuestion(expr, str, matchLen)
	}
	if ok, quantifier := isQuantifier(operator); ok {
		return matchQuantifier(expr, str, matchLen, quantifier)
	}
	if isAlternate(head) {
		return matchAlternate(expr, str, matchLen)
	}
	if isUnit(head) {
		if doesUnitMatch(expr, str) {
			return matchExpr(rest, str[1:], matchLen+1)
		}
		return false, 0
	}
	panic(fmt.Sprintf("unexpected token in expr %s\n", expr))
}

func matchStar(expr string, str string, matchLen int) (bool, int) {
	return matchMultiple(expr, str, matchLen, 0, math.MaxInt)
}

func matchPlus(expr string, str string, matchLen int) (bool, int) {
	return matchMultiple(expr, str, matchLen, 1, math.MaxInt)
}

func matchQuestion(expr string, str string, matchLen int) (bool, int) {
	return matchMultiple(expr, str, matchLen, 0, 1)
}

func matchQuantifier(expr string, str string, matchLen int, quantifier int) (bool, int) {
	return matchMultiple(expr, str, matchLen, quantifier, quantifier)
}

func matchMultiple(expr string, str string, matchLen int, minMatchLen int, maxMatchLen int) (bool, int) {
	head, _, rest := splitExpr(expr)
	submatchLen := -1
	for maxMatchLen == math.MaxInt || (submatchLen < maxMatchLen) {
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
