package tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"reggaex/rgx"
)

// limits
func TestStrLimits(t *testing.T) {
	expr := `^https://\w+.(com|net|org)[@/#]+.*$`
	str := `https://qwerty123.com@hey/there`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestStrLimitsFailsAtStart(t *testing.T) {
	expr := `^https://w+.(com|net|org)[@/#]+.*$`
	str := `xxxhttps://qwerty123.com@hey/there`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestStrLimitsFailsAtEnd(t *testing.T) {
	expr := `^https://\w+.(com|net|org)[@/#]+[ab]$`
	str := `https://qwerty123.com@xxx`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestNoStartLimit(t *testing.T) {
	expr := `https://\w+.(com|net|org)[@/#]+.*$`
	str := `xxxhttps://qwerty123.com@hey/there`
	res, _ := rgx.Match(expr, str)
	expected := `https://qwerty123.com@hey/there`
	assert.True(t, res.Matched)
	assert.Equal(t, expected, res.MatchStr)
}

func TestNoEndLimit(t *testing.T) {
	expr := `^https://\w+.(com|net|org)[@/#]+`
	str := `https://qwerty123.com@xxx`
	res, _ := rgx.Match(expr, str)
	expected := `https://qwerty123.com@`
	assert.True(t, res.Matched)
	assert.Equal(t, expected, res.MatchStr)
}

func TestInvalidToken(t *testing.T) {
	expr := `^]ab$`
	str := `aa`
	_, err := rgx.Match(expr, str)
	expected := errors.New("unexpected token in expr ]ab$\n")
	assert.Equal(t, expected, err)
}

// empty
func TestEmptyExpr(t *testing.T) {
	expr := ``
	str := `https://qwerty123.com@xxx`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestEmptyStr(t *testing.T) {
	expr := `[xyz]+`
	str := ``
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// dot
func TestDot(t *testing.T) {
	expr := `^.{2}y$`
	str := `x2y`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestDotFails(t *testing.T) {
	expr := `^.{2}y$`
	str := `1xx`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// escape sequences
func TestWhitespace(t *testing.T) {
	expr := `^\s[abc]+\s[xyz]+$`
	str := ` abc	yzx`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestWhitespaceFails(t *testing.T) {
	expr := `^\s[abc]+\s[xyz]+$`
	str := `abc	yzx`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestDigit(t *testing.T) {
	expr := `^\d+$`
	str := `123`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestDigitFails(t *testing.T) {
	expr := `^\d+$`
	str := `1xx`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestWordChar(t *testing.T) {
	expr := `^\w+$`
	str := `xyz_a23`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestWordCharFails(t *testing.T) {
	expr := `^\w+$`
	str := `x z`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// negative escape sequences
func TestNegWhitespace(t *testing.T) {
	expr := `^\S+$`
	str := `xzy123./?`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestNegWhitespaceFails(t *testing.T) {
	expr := `^\S+$`
	str := `xzy 123./?`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestNegDigit(t *testing.T) {
	expr := `^\D+$`
	str := `xzy ./?`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestNegDigitFails(t *testing.T) {
	expr := `^\D+$`
	str := `xzy 22./?`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestNegWordChar(t *testing.T) {
	expr := `^\W+$`
	str := `: ./?`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestNegWordCharFails(t *testing.T) {
	expr := `^\W+$`
	str := `x_1./?`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// sets
func TestSet(t *testing.T) {
	expr := `^[abc]+$`
	str := `aabcbba`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestSetFails(t *testing.T) {
	expr := `^[abc]+$`
	str := `ab123a`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestNegSet(t *testing.T) {
	expr := `^[^abc]+$`
	str := `xyz`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestNegSetFails(t *testing.T) {
	expr := `^[^abc]+$`
	str := `xaz`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// ranges
func TestRange(t *testing.T) {
	expr := `^[a-g]+$`
	str := `bda`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestRangeFails(t *testing.T) {
	expr := `^[a-g]+$`
	str := `x`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestNegRange(t *testing.T) {
	expr := `^[^a-c]+$`
	str := `xyz`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestNegRangeFails(t *testing.T) {
	expr := `^[^a-g]+$`
	str := `ab`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestDoubleRange(t *testing.T) {
	expr := `^[a-g0-9]+$`
	str := `ab23`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestDoubleNegRange(t *testing.T) {
	expr := `^[^a-g0-9]+$`
	str := `ab23`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestComplexRange(t *testing.T) {
	expr := `^[a-g0-9#$%@]+$`
	str := `ab23@`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestComplexNegRange(t *testing.T) {
	expr := `^[^a-g0-9#$%@]+$`
	str := `ab23@`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// options
func TestOptions(t *testing.T) {
	expr := `^(foo|bar)$`
	str := `foo`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestOptionsFails(t *testing.T) {
	expr := `^(foo|bar)$`
	str := `fizzbuzz`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

// operators
func TestPlus(t *testing.T) {
	expr := `^[ab]+$`
	for _, str := range []string{"a", "ab", "ababbababbab"} {
		res, _ := rgx.Match(expr, str)
		expected := res.MatchStr
		assert.True(t, res.Matched)
		assert.Equal(t, expected, str)
	}
}

func TestPlusFails(t *testing.T) {
	expr := `^[ab]+$`
	for _, str := range []string{"", "xyz"} {
		res, _ := rgx.Match(expr, str)
		assert.False(t, res.Matched)
	}
}

func TestStar(t *testing.T) {
	expr := `^x[yz]*$`
	for _, str := range []string{"x", "xy", "xyz", "xyzzzzyzy"} {
		res, _ := rgx.Match(expr, str)
		expected := res.MatchStr
		assert.True(t, res.Matched)
		assert.Equal(t, expected, str)
	}
}

func TestStarFails(t *testing.T) {
	expr := `^[ab]*$`
	str := `xyz`
	res, _ := rgx.Match(expr, str)
	assert.False(t, res.Matched)
}

func TestQuestion(t *testing.T) {
	expr := `^x[yz]?$`
	for _, str := range []string{"x", "xy"} {
		res, _ := rgx.Match(expr, str)
		expected := res.MatchStr
		assert.True(t, res.Matched)
		assert.Equal(t, expected, str)
	}
}

func TestQuestionFails(t *testing.T) {
	expr := `^x[yz]?$`
	for _, str := range []string{"xyy", "xab"} {
		res, _ := rgx.Match(expr, str)
		assert.False(t, res.Matched)
	}
}

// quantifier
func TestQuantifier(t *testing.T) {
	expr := `^[ab]{10}$`
	str := `aabbaaabba`
	res, _ := rgx.Match(expr, str)
	expected := res.MatchStr
	assert.True(t, res.Matched)
	assert.Equal(t, expected, str)
}

func TestInvalidQuantifier(t *testing.T) {
	expr := `^[ab]{x}$`
	str := `aa`
	_, err := rgx.Match(expr, str)
	expected := errors.New("supplied value x is not a number\n")
	assert.Equal(t, expected, err)
}
