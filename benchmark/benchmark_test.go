package strcase

import (
	"regexp"
	"strings"
	"testing"

	"github.com/ettle/strcase"
	fa "github.com/fatih/camelcase"
	ia "github.com/iancoleman/strcase"
	se "github.com/segmentio/go-camelcase"
	st "github.com/stoewer/go-strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Test cases:
// Minor variations in length , but doesn't seem to vastly change the relative
// results
// var testLower = "id"
// var testLower = "get user id"
// var testLower = "filter on user permission group id request date"

var testLower = "get user id"

var testSnake = strcase.ToSnake(testLower)
var testSNAKE = strcase.ToSNAKE(testLower)
var testKebab = strcase.ToKebab(testLower)
var testGoSnake = strcase.ToGoSnake(testLower)
var testCamel = strcase.ToCamel(testLower)
var testPascal = strcase.ToPascal(testLower)
var testTitle = strcase.ToCase(testLower, strcase.TitleCase, ' ')
var testSplit = strcase.ToCase(testLower, strcase.TitleCase, '_')

var testCustom = "ssl*user*id"
var testCustomExpected = "SSL_user_id"

func BenchmarkToTitle(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = strcase.ToCase(testLower, strcase.TitleCase, ' ')
	}
	expected := testTitle
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}
func BenchmarkToSnake(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = strcase.ToSnake(testCamel)
	}
	expected := testSnake
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}
func BenchmarkToSNAKE(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = strcase.ToSNAKE(testCamel)
	}
	expected := testSNAKE
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}
func BenchmarkToGoSnake(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = strcase.ToGoSnake(testCamel)
	}
	expected := testGoSnake
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}
func BenchmarkToCustomCaser(b *testing.B) {
	c := strcase.NewCaser(false, map[string]bool{"SSL": true}, strcase.NewSplitFn([]rune{'*'}))
	b.ResetTimer()
	var s string
	for n := 0; n < b.N; n++ {
		s = c.ToSnake(testCustom)
	}
	expected := testCustomExpected
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// ********************************************************
// Stdlib
//

// golang.org/x/text/cases is the recommended replacement
// for stdlib's now deprecated strings.ToTitle
func BenchmarkXTextCases(b *testing.B) {
	caser := cases.Title(language.AmericanEnglish)
	var s string
	for n := 0; n < b.N; n++ {
		s = caser.String(testLower)
	}
	expected := testTitle
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// ********************************************************
// Other packages
//

// From github.com/segmentio/go-camelcase
// MIT License
// A very fast package - no unicode, intialism, or customizations
// If you need speed, use or fork this library or segmentio/go-snakecase
func BenchmarkSegment(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = se.Camelcase(testSnake)
	}
	expected := testCamel
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// From github.com/iancoleman/strcase
// The most popular go strcase packages I found
// About an order of magnitude slower
func BenchmarkToSnakeIan(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = ia.ToSnake(testCamel)
	}
	expected := testSnake
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// From github.com/stoewer/go-strcase
// MIT License
// In most tests, it's just a smidge slower than the comparable function
func BenchmarkToSnakeStoewer(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = st.SnakeCase(testCamel)
	}
	expected := testSnake
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// From github.com/fatih/camelcase
// MIT License
// This isn't quite an even comparison since it only splits and doesn't
// specify how to join the words back together. Figured strings.Join was
// reasonable
func BenchmarkToSnakeFatih(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		sa := fa.Split(testPascal)
		s = strings.Join(sa, "_")
	}
	expected := testSplit
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// ********************************************************
// Code snippets
//

// From https://www.golangprograms.com/golang-convert-string-into-snake-case.html
// Terms of use: https://www.golangprograms.com/terms-of-use
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCaseGolangPrograms(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func BenchmarkToSnakeGolangPrograms(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = ToSnakeCaseGolangPrograms(testCamel)
	}
	expected := testSnake
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}

// From https://siongui.github.io/2017/02/18/go-kebab-case-to-camel-case/
// License: UNLICENSE
// https://github.com/siongui/userpages/blob/master/UNLICENSE
func kebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}

func BenchmarkToSnakeSiongui(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = kebabToCamelCase(testKebab)
	}
	expected := testCamel
	if expected != s {
		b.Fatalf("Expected %s, got %s", expected, s)
	}
}
