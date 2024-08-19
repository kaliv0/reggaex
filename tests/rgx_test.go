package tests

import (
	"github.com/stretchr/testify/assert"
	"reggaex/rgx"
	"testing"
)

// with string limits
func Test1(t *testing.T) {
	expr := `^http://(\a|\d)+.(com|net|org)[@/#]+.*$`
	str := `http://qwerty123.com@hey/there`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

// with string limits fails at start
func Test2(t *testing.T) {
	expr := `^http://(\a|\d)+.(com|net|org)[@/#]+.*$`
	str := `xxxhttp://qwerty123.com@hey/there`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

// with string limits fails at end
func Test3(t *testing.T) {
	expr := `^http://(\a|\d)+.(com|net|org)[@/#]+[ab]$`
	str := `http://qwerty123.com@xxx`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

// no start limit
func Test4(t *testing.T) {
	expr := `http://(\a|\d)+.(com|net|org)[@/#]+.*$`
	str := `xxxhttp://qwerty123.com@hey/there`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	actual := str[matchPos : matchPos+matchLen]
	expected := `http://qwerty123.com@hey/there`
	assert.True(t, matched)
	assert.Equal(t, actual, expected)
}

// no end limit
func Test5(t *testing.T) {
	expr := `^http://(\a|\d)+.(com|net|org)[@/#]+`
	str := `http://qwerty123.com@xxx`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	actual := str[matchPos : matchPos+matchLen]
	expected := `http://qwerty123.com@`
	assert.True(t, matched)
	assert.Equal(t, actual, expected)
}
