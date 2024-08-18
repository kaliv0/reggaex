package tests

import (
	"reggaex/rgx"
	"testing"
)

func Test1(t *testing.T) {
	expr := `^http://(\a|\d)+.(com|net|org)[@/#]+.*$`
	str := `http://qwerty123.com@hey/there`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	if !matched {
		matchRange := str[matchPos : matchPos+matchLen]
		t.Errorf("matchExpr(`%s`, '%s') = %s", expr, str, matchRange)
	}
}
