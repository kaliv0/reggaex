package main

import (
	"fmt"

	// "github.com/kaliv0/reggaex/rgx/matcher"
	"reggaex/rgx"
)

// type Matcher struct {
// 	pattern    string
// 	expression string

// 	subPattern    string
// 	subExpression string
// }

// func (m *Matcher) setSubPattern(start, end int){
// 	m.subPattern = m.pattern[start:end]
// }

func main() {
	expr := `^http://(\a|\d)+.(com|net|org)[@/#]+.*$`
	str := `http://qwerty123.com@hey/there`
	matcher := rgx.Matcher{Expr: expr}
	//str := `http://clumsy_123_computer.com@hey/there`
	//str := ""

	matched, matchPos, matchLen := matcher.Match(expr, str)
	if matched {
		matchRange := str[matchPos : matchPos+matchLen]
		fmt.Printf("matchExpr(`%s`, '%s') = %s", expr, str, matchRange)
	} else {
		fmt.Printf("matchExpr(`%s`, '%s') = False", expr, str)
	}
}
