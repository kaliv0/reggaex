package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"reggaex/rgx"
)

func TestStrLimits(t *testing.T) {
	expr := `^http://\w+.(com|net|org)[@/#]+.*$`
	str := `http://qwerty123.com@hey/there`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestStrLimitsFailsAtStart(t *testing.T) {
	expr := `^http://w+.(com|net|org)[@/#]+.*$`
	str := `xxxhttp://qwerty123.com@hey/there`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

func TestStrLimitsFailsAtEnd(t *testing.T) {
	expr := `^http://\w+.(com|net|org)[@/#]+[ab]$`
	str := `http://qwerty123.com@xxx`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

func TestNoStartLimit(t *testing.T) {
	expr := `http://\w+.(com|net|org)[@/#]+.*$`
	str := `xxxhttp://qwerty123.com@hey/there`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	actual := str[matchPos : matchPos+matchLen]
	expected := `http://qwerty123.com@hey/there`
	assert.True(t, matched)
	assert.Equal(t, actual, expected)
}

func TestNoEndLimit(t *testing.T) {
	expr := `^http://\w+.(com|net|org)[@/#]+`
	str := `http://qwerty123.com@xxx`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	actual := str[matchPos : matchPos+matchLen]
	expected := `http://qwerty123.com@`
	assert.True(t, matched)
	assert.Equal(t, actual, expected)
}

// escape sequences
func TestWhitespace(t *testing.T) {
	expr := `^\s[abc]+\s[xyz]+$`
	str := ` abc	yzx`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestWhitespaceFails(t *testing.T) {
	expr := `^\s[abc]+\s[xyz]+$`
	str := `abc	yzx`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

func TestDigit(t *testing.T) {
	expr := `^\d+$`
	str := `123`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestDigitFails(t *testing.T) {
	expr := `^\d+$`
	str := `1xx`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

func TestWordChar(t *testing.T) {
	expr := `^\w+$`
	str := `xyz_a23`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestWordCharFails(t *testing.T) {
	expr := `^\w+$`
	str := `x z`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

// negative escape sequences
func TestNegWhitespace(t *testing.T) {
	expr := `^\S+$`
	str := `xzy123./?`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestNegWhitespaceFails(t *testing.T) {
	expr := `^\S+$`
	str := `xzy 123./?`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

func TestNegDigit(t *testing.T) {
	expr := `^\D+$`
	str := `xzy ./?`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestNegDigitFails(t *testing.T) {
	expr := `^\D+$`
	str := `xzy 22./?`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

func TestNegWordChar(t *testing.T) {
	expr := `^\W+$`
	str := `: ./?`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestNegWordCharFails(t *testing.T) {
	expr := `^\W+$`
	str := `x_1./?`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}

// sets
func TestSet(t *testing.T) {
	expr := `^[abc]+$`
	str := `aabcbba`
	matched, matchPos, matchLen := rgx.Match(expr, str)
	expected := str[matchPos : matchPos+matchLen]
	assert.True(t, matched)
	assert.Equal(t, expected, str)
}

func TestSetFails(t *testing.T) {
	expr := `^[abc]+$`
	str := `ab123a`
	matched, _, _ := rgx.Match(expr, str)
	assert.False(t, matched)
}
