package strcase

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// Obviously 100% test coverage isn't everything but...
func TestEdges(t *testing.T) {
	t.Run("Original WordCase", func(t *testing.T) {
		assertEqual(t, "FreeBSD", convertWithoutInitialisms("FreeBSD", 0, Original))
		assertEqual(t, "FreeBSD", convertWithGoInitialisms("FreeBSD", 0, Original))
	})
	t.Run("Don't call convertWithInitialisms for UpperCase", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		convertWithGoInitialisms("foo", 0, UpperCase)
	})
}

func TestAll(t *testing.T) {
	// Instead of testing, we can generate the outputs to make it easier to
	// add more test cases or functions
	generate := false

	type data struct {
		input    string
		snake    string
		goSnake  string
		SNAKE    string
		kebab    string
		goKebab  string
		KEBAB    string
		pascal   string
		goPascal string
		camel    string
		goCamel  string
		// Test ToCase function
		title   string
		goTitle string
	}
	for _, test := range []data{
		{
			input:    "Hello world!",
			snake:    "hello_world!",
			goSnake:  "hello_world!",
			SNAKE:    "HELLO_WORLD!",
			kebab:    "hello-world!",
			goKebab:  "hello-world!",
			KEBAB:    "HELLO-WORLD!",
			pascal:   "HelloWorld!",
			goPascal: "HelloWorld!",
			camel:    "helloWorld!",
			goCamel:  "helloWorld!",
			title:    "Hello World!",
			goTitle:  "Hello World!",
		},
		{
			input:    "",
			snake:    "",
			goSnake:  "",
			SNAKE:    "",
			kebab:    "",
			goKebab:  "",
			KEBAB:    "",
			pascal:   "",
			goPascal: "",
			camel:    "",
			goCamel:  "",
			title:    "",
			goTitle:  "",
		},
		{
			input:    ".",
			snake:    "",
			goSnake:  "",
			SNAKE:    "",
			kebab:    "",
			goKebab:  "",
			KEBAB:    "",
			pascal:   "",
			goPascal: "",
			camel:    "",
			goCamel:  "",
			title:    "",
			goTitle:  "",
		},
		{
			input:    "A",
			snake:    "a",
			goSnake:  "a",
			SNAKE:    "A",
			kebab:    "a",
			goKebab:  "a",
			KEBAB:    "A",
			pascal:   "A",
			goPascal: "A",
			camel:    "a",
			goCamel:  "a",
			title:    "A",
			goTitle:  "A",
		},
		{
			input:    "a",
			snake:    "a",
			goSnake:  "a",
			SNAKE:    "A",
			kebab:    "a",
			goKebab:  "a",
			KEBAB:    "A",
			pascal:   "A",
			goPascal: "A",
			camel:    "a",
			goCamel:  "a",
			title:    "A",
			goTitle:  "A",
		},
		{
			input:    "foo",
			snake:    "foo",
			goSnake:  "foo",
			SNAKE:    "FOO",
			kebab:    "foo",
			goKebab:  "foo",
			KEBAB:    "FOO",
			pascal:   "Foo",
			goPascal: "Foo",
			camel:    "foo",
			goCamel:  "foo",
			title:    "Foo",
			goTitle:  "Foo",
		},
		{
			input:    "snake_case",
			snake:    "snake_case",
			goSnake:  "snake_case",
			SNAKE:    "SNAKE_CASE",
			kebab:    "snake-case",
			goKebab:  "snake-case",
			KEBAB:    "SNAKE-CASE",
			pascal:   "SnakeCase",
			goPascal: "SnakeCase",
			camel:    "snakeCase",
			goCamel:  "snakeCase",
			title:    "Snake Case",
			goTitle:  "Snake Case",
		},
		{
			input:    "SNAKE_CASE",
			snake:    "snake_case",
			goSnake:  "snake_case",
			SNAKE:    "SNAKE_CASE",
			kebab:    "snake-case",
			goKebab:  "snake-case",
			KEBAB:    "SNAKE-CASE",
			pascal:   "SnakeCase",
			goPascal: "SnakeCase",
			camel:    "snakeCase",
			goCamel:  "snakeCase",
			title:    "Snake Case",
			goTitle:  "Snake Case",
		},
		{
			input:    "kebab-case",
			snake:    "kebab_case",
			goSnake:  "kebab_case",
			SNAKE:    "KEBAB_CASE",
			kebab:    "kebab-case",
			goKebab:  "kebab-case",
			KEBAB:    "KEBAB-CASE",
			pascal:   "KebabCase",
			goPascal: "KebabCase",
			camel:    "kebabCase",
			goCamel:  "kebabCase",
			title:    "Kebab Case",
			goTitle:  "Kebab Case",
		},
		{
			input:    "PascalCase",
			snake:    "pascal_case",
			goSnake:  "pascal_case",
			SNAKE:    "PASCAL_CASE",
			kebab:    "pascal-case",
			goKebab:  "pascal-case",
			KEBAB:    "PASCAL-CASE",
			pascal:   "PascalCase",
			goPascal: "PascalCase",
			camel:    "pascalCase",
			goCamel:  "pascalCase",
			title:    "Pascal Case",
			goTitle:  "Pascal Case",
		},
		{
			input:    "camelCase",
			snake:    "camel_case",
			goSnake:  "camel_case",
			SNAKE:    "CAMEL_CASE",
			kebab:    "camel-case",
			goKebab:  "camel-case",
			KEBAB:    "CAMEL-CASE",
			pascal:   "CamelCase",
			goPascal: "CamelCase",
			camel:    "camelCase",
			goCamel:  "camelCase",
			title:    "Camel Case",
			goTitle:  "Camel Case",
		},
		{
			input:    "Title Case",
			snake:    "title_case",
			goSnake:  "title_case",
			SNAKE:    "TITLE_CASE",
			kebab:    "title-case",
			goKebab:  "title-case",
			KEBAB:    "TITLE-CASE",
			pascal:   "TitleCase",
			goPascal: "TitleCase",
			camel:    "titleCase",
			goCamel:  "titleCase",
			title:    "Title Case",
			goTitle:  "Title Case",
		},
		{
			input:    "point.case",
			snake:    "point_case",
			goSnake:  "point_case",
			SNAKE:    "POINT_CASE",
			kebab:    "point-case",
			goKebab:  "point-case",
			KEBAB:    "POINT-CASE",
			pascal:   "PointCase",
			goPascal: "PointCase",
			camel:    "pointCase",
			goCamel:  "pointCase",
			title:    "Point Case",
			goTitle:  "Point Case",
		},
		{
			input:    "snake_case_with_more_words",
			snake:    "snake_case_with_more_words",
			goSnake:  "snake_case_with_more_words",
			SNAKE:    "SNAKE_CASE_WITH_MORE_WORDS",
			kebab:    "snake-case-with-more-words",
			goKebab:  "snake-case-with-more-words",
			KEBAB:    "SNAKE-CASE-WITH-MORE-WORDS",
			pascal:   "SnakeCaseWithMoreWords",
			goPascal: "SnakeCaseWithMoreWords",
			camel:    "snakeCaseWithMoreWords",
			goCamel:  "snakeCaseWithMoreWords",
			title:    "Snake Case With More Words",
			goTitle:  "Snake Case With More Words",
		},
		{
			input:    "SNAKE_CASE_WITH_MORE_WORDS",
			snake:    "snake_case_with_more_words",
			goSnake:  "snake_case_with_more_words",
			SNAKE:    "SNAKE_CASE_WITH_MORE_WORDS",
			kebab:    "snake-case-with-more-words",
			goKebab:  "snake-case-with-more-words",
			KEBAB:    "SNAKE-CASE-WITH-MORE-WORDS",
			pascal:   "SnakeCaseWithMoreWords",
			goPascal: "SnakeCaseWithMoreWords",
			camel:    "snakeCaseWithMoreWords",
			goCamel:  "snakeCaseWithMoreWords",
			title:    "Snake Case With More Words",
			goTitle:  "Snake Case With More Words",
		},
		{
			input:    "kebab-case-with-more-words",
			snake:    "kebab_case_with_more_words",
			goSnake:  "kebab_case_with_more_words",
			SNAKE:    "KEBAB_CASE_WITH_MORE_WORDS",
			kebab:    "kebab-case-with-more-words",
			goKebab:  "kebab-case-with-more-words",
			KEBAB:    "KEBAB-CASE-WITH-MORE-WORDS",
			pascal:   "KebabCaseWithMoreWords",
			goPascal: "KebabCaseWithMoreWords",
			camel:    "kebabCaseWithMoreWords",
			goCamel:  "kebabCaseWithMoreWords",
			title:    "Kebab Case With More Words",
			goTitle:  "Kebab Case With More Words",
		},
		{
			input:    "PascalCaseWithMoreWords",
			snake:    "pascal_case_with_more_words",
			goSnake:  "pascal_case_with_more_words",
			SNAKE:    "PASCAL_CASE_WITH_MORE_WORDS",
			kebab:    "pascal-case-with-more-words",
			goKebab:  "pascal-case-with-more-words",
			KEBAB:    "PASCAL-CASE-WITH-MORE-WORDS",
			pascal:   "PascalCaseWithMoreWords",
			goPascal: "PascalCaseWithMoreWords",
			camel:    "pascalCaseWithMoreWords",
			goCamel:  "pascalCaseWithMoreWords",
			title:    "Pascal Case With More Words",
			goTitle:  "Pascal Case With More Words",
		},
		{
			input:    "camelCaseWithMoreWords",
			snake:    "camel_case_with_more_words",
			goSnake:  "camel_case_with_more_words",
			SNAKE:    "CAMEL_CASE_WITH_MORE_WORDS",
			kebab:    "camel-case-with-more-words",
			goKebab:  "camel-case-with-more-words",
			KEBAB:    "CAMEL-CASE-WITH-MORE-WORDS",
			pascal:   "CamelCaseWithMoreWords",
			goPascal: "CamelCaseWithMoreWords",
			camel:    "camelCaseWithMoreWords",
			goCamel:  "camelCaseWithMoreWords",
			title:    "Camel Case With More Words",
			goTitle:  "Camel Case With More Words",
		},
		{
			input:    "Title Case With More Words",
			snake:    "title_case_with_more_words",
			goSnake:  "title_case_with_more_words",
			SNAKE:    "TITLE_CASE_WITH_MORE_WORDS",
			kebab:    "title-case-with-more-words",
			goKebab:  "title-case-with-more-words",
			KEBAB:    "TITLE-CASE-WITH-MORE-WORDS",
			pascal:   "TitleCaseWithMoreWords",
			goPascal: "TitleCaseWithMoreWords",
			camel:    "titleCaseWithMoreWords",
			goCamel:  "titleCaseWithMoreWords",
			title:    "Title Case With More Words",
			goTitle:  "Title Case With More Words",
		},
		{
			input:    "point.case.with.more.words",
			snake:    "point_case_with_more_words",
			goSnake:  "point_case_with_more_words",
			SNAKE:    "POINT_CASE_WITH_MORE_WORDS",
			kebab:    "point-case-with-more-words",
			goKebab:  "point-case-with-more-words",
			KEBAB:    "POINT-CASE-WITH-MORE-WORDS",
			pascal:   "PointCaseWithMoreWords",
			goPascal: "PointCaseWithMoreWords",
			camel:    "pointCaseWithMoreWords",
			goCamel:  "pointCaseWithMoreWords",
			title:    "Point Case With More Words",
			goTitle:  "Point Case With More Words",
		},
		{
			input:    "snake_case__with___multiple____delimiters",
			snake:    "snake_case_with_multiple_delimiters",
			goSnake:  "snake_case_with_multiple_delimiters",
			SNAKE:    "SNAKE_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab:    "snake-case-with-multiple-delimiters",
			goKebab:  "snake-case-with-multiple-delimiters",
			KEBAB:    "SNAKE-CASE-WITH-MULTIPLE-DELIMITERS",
			pascal:   "SnakeCaseWithMultipleDelimiters",
			goPascal: "SnakeCaseWithMultipleDelimiters",
			camel:    "snakeCaseWithMultipleDelimiters",
			goCamel:  "snakeCaseWithMultipleDelimiters",
			title:    "Snake Case With Multiple Delimiters",
			goTitle:  "Snake Case With Multiple Delimiters",
		},
		{
			input:    "SNAKE_CASE__WITH___multiple____DELIMITERS",
			snake:    "snake_case_with_multiple_delimiters",
			goSnake:  "snake_case_with_multiple_delimiters",
			SNAKE:    "SNAKE_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab:    "snake-case-with-multiple-delimiters",
			goKebab:  "snake-case-with-multiple-delimiters",
			KEBAB:    "SNAKE-CASE-WITH-MULTIPLE-DELIMITERS",
			pascal:   "SnakeCaseWithMultipleDelimiters",
			goPascal: "SnakeCaseWithMultipleDelimiters",
			camel:    "snakeCaseWithMultipleDelimiters",
			goCamel:  "snakeCaseWithMultipleDelimiters",
			title:    "Snake Case With Multiple Delimiters",
			goTitle:  "Snake Case With Multiple Delimiters",
		},
		{
			input:    "kebab-case--with---multiple----delimiters",
			snake:    "kebab_case_with_multiple_delimiters",
			goSnake:  "kebab_case_with_multiple_delimiters",
			SNAKE:    "KEBAB_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab:    "kebab-case-with-multiple-delimiters",
			goKebab:  "kebab-case-with-multiple-delimiters",
			KEBAB:    "KEBAB-CASE-WITH-MULTIPLE-DELIMITERS",
			pascal:   "KebabCaseWithMultipleDelimiters",
			goPascal: "KebabCaseWithMultipleDelimiters",
			camel:    "kebabCaseWithMultipleDelimiters",
			goCamel:  "kebabCaseWithMultipleDelimiters",
			title:    "Kebab Case With Multiple Delimiters",
			goTitle:  "Kebab Case With Multiple Delimiters",
		},
		{
			input:    "Title Case  With   Multiple    Delimiters",
			snake:    "title_case_with_multiple_delimiters",
			goSnake:  "title_case_with_multiple_delimiters",
			SNAKE:    "TITLE_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab:    "title-case-with-multiple-delimiters",
			goKebab:  "title-case-with-multiple-delimiters",
			KEBAB:    "TITLE-CASE-WITH-MULTIPLE-DELIMITERS",
			pascal:   "TitleCaseWithMultipleDelimiters",
			goPascal: "TitleCaseWithMultipleDelimiters",
			camel:    "titleCaseWithMultipleDelimiters",
			goCamel:  "titleCaseWithMultipleDelimiters",
			title:    "Title Case With Multiple Delimiters",
			goTitle:  "Title Case With Multiple Delimiters",
		},
		{
			input:    "point.case..with...multiple....delimiters",
			snake:    "point_case_with_multiple_delimiters",
			goSnake:  "point_case_with_multiple_delimiters",
			SNAKE:    "POINT_CASE_WITH_MULTIPLE_DELIMITERS",
			kebab:    "point-case-with-multiple-delimiters",
			goKebab:  "point-case-with-multiple-delimiters",
			KEBAB:    "POINT-CASE-WITH-MULTIPLE-DELIMITERS",
			pascal:   "PointCaseWithMultipleDelimiters",
			goPascal: "PointCaseWithMultipleDelimiters",
			camel:    "pointCaseWithMultipleDelimiters",
			goCamel:  "pointCaseWithMultipleDelimiters",
			title:    "Point Case With Multiple Delimiters",
			goTitle:  "Point Case With Multiple Delimiters",
		},
		{
			input:    " leading space",
			snake:    "leading_space",
			goSnake:  "leading_space",
			SNAKE:    "LEADING_SPACE",
			kebab:    "leading-space",
			goKebab:  "leading-space",
			KEBAB:    "LEADING-SPACE",
			pascal:   "LeadingSpace",
			goPascal: "LeadingSpace",
			camel:    "leadingSpace",
			goCamel:  "leadingSpace",
			title:    "Leading Space",
			goTitle:  "Leading Space",
		},
		{
			input:    "   leading spaces",
			snake:    "leading_spaces",
			goSnake:  "leading_spaces",
			SNAKE:    "LEADING_SPACES",
			kebab:    "leading-spaces",
			goKebab:  "leading-spaces",
			KEBAB:    "LEADING-SPACES",
			pascal:   "LeadingSpaces",
			goPascal: "LeadingSpaces",
			camel:    "leadingSpaces",
			goCamel:  "leadingSpaces",
			title:    "Leading Spaces",
			goTitle:  "Leading Spaces",
		},
		{
			input:    "\t\t\r\n leading whitespaces",
			snake:    "leading_whitespaces",
			goSnake:  "leading_whitespaces",
			SNAKE:    "LEADING_WHITESPACES",
			kebab:    "leading-whitespaces",
			goKebab:  "leading-whitespaces",
			KEBAB:    "LEADING-WHITESPACES",
			pascal:   "LeadingWhitespaces",
			goPascal: "LeadingWhitespaces",
			camel:    "leadingWhitespaces",
			goCamel:  "leadingWhitespaces",
			title:    "Leading Whitespaces",
			goTitle:  "Leading Whitespaces",
		},
		{
			input:    "trailing space ",
			snake:    "trailing_space",
			goSnake:  "trailing_space",
			SNAKE:    "TRAILING_SPACE",
			kebab:    "trailing-space",
			goKebab:  "trailing-space",
			KEBAB:    "TRAILING-SPACE",
			pascal:   "TrailingSpace",
			goPascal: "TrailingSpace",
			camel:    "trailingSpace",
			goCamel:  "trailingSpace",
			title:    "Trailing Space",
			goTitle:  "Trailing Space",
		},
		{
			input:    "trailing spaces   ",
			snake:    "trailing_spaces",
			goSnake:  "trailing_spaces",
			SNAKE:    "TRAILING_SPACES",
			kebab:    "trailing-spaces",
			goKebab:  "trailing-spaces",
			KEBAB:    "TRAILING-SPACES",
			pascal:   "TrailingSpaces",
			goPascal: "TrailingSpaces",
			camel:    "trailingSpaces",
			goCamel:  "trailingSpaces",
			title:    "Trailing Spaces",
			goTitle:  "Trailing Spaces",
		},
		{
			input:    "trailing whitespaces\t\t\r\n",
			snake:    "trailing_whitespaces",
			goSnake:  "trailing_whitespaces",
			SNAKE:    "TRAILING_WHITESPACES",
			kebab:    "trailing-whitespaces",
			goKebab:  "trailing-whitespaces",
			KEBAB:    "TRAILING-WHITESPACES",
			pascal:   "TrailingWhitespaces",
			goPascal: "TrailingWhitespaces",
			camel:    "trailingWhitespaces",
			goCamel:  "trailingWhitespaces",
			title:    "Trailing Whitespaces",
			goTitle:  "Trailing Whitespaces",
		},
		{
			input:    " on both sides ",
			snake:    "on_both_sides",
			goSnake:  "on_both_sides",
			SNAKE:    "ON_BOTH_SIDES",
			kebab:    "on-both-sides",
			goKebab:  "on-both-sides",
			KEBAB:    "ON-BOTH-SIDES",
			pascal:   "OnBothSides",
			goPascal: "OnBothSides",
			camel:    "onBothSides",
			goCamel:  "onBothSides",
			title:    "On Both Sides",
			goTitle:  "On Both Sides",
		},
		{
			input:    "    many on both sides  ",
			snake:    "many_on_both_sides",
			goSnake:  "many_on_both_sides",
			SNAKE:    "MANY_ON_BOTH_SIDES",
			kebab:    "many-on-both-sides",
			goKebab:  "many-on-both-sides",
			KEBAB:    "MANY-ON-BOTH-SIDES",
			pascal:   "ManyOnBothSides",
			goPascal: "ManyOnBothSides",
			camel:    "manyOnBothSides",
			goCamel:  "manyOnBothSides",
			title:    "Many On Both Sides",
			goTitle:  "Many On Both Sides",
		},
		{
			input:    "\rwhitespaces on both sides\t\t\r\n",
			snake:    "whitespaces_on_both_sides",
			goSnake:  "whitespaces_on_both_sides",
			SNAKE:    "WHITESPACES_ON_BOTH_SIDES",
			kebab:    "whitespaces-on-both-sides",
			goKebab:  "whitespaces-on-both-sides",
			KEBAB:    "WHITESPACES-ON-BOTH-SIDES",
			pascal:   "WhitespacesOnBothSides",
			goPascal: "WhitespacesOnBothSides",
			camel:    "whitespacesOnBothSides",
			goCamel:  "whitespacesOnBothSides",
			title:    "Whitespaces On Both Sides",
			goTitle:  "Whitespaces On Both Sides",
		},
		{
			input:    "  extraSpaces in_This TestCase Of MIXED_CASES\t",
			snake:    "extra_spaces_in_this_test_case_of_mixed_cases",
			goSnake:  "extra_spaces_in_this_test_case_of_mixed_cases",
			SNAKE:    "EXTRA_SPACES_IN_THIS_TEST_CASE_OF_MIXED_CASES",
			kebab:    "extra-spaces-in-this-test-case-of-mixed-cases",
			goKebab:  "extra-spaces-in-this-test-case-of-mixed-cases",
			KEBAB:    "EXTRA-SPACES-IN-THIS-TEST-CASE-OF-MIXED-CASES",
			pascal:   "ExtraSpacesInThisTestCaseOfMixedCases",
			goPascal: "ExtraSpacesInThisTestCaseOfMixedCases",
			camel:    "extraSpacesInThisTestCaseOfMixedCases",
			goCamel:  "extraSpacesInThisTestCaseOfMixedCases",
			title:    "Extra Spaces In This Test Case Of Mixed Cases",
			goTitle:  "Extra Spaces In This Test Case Of Mixed Cases",
		},
		{
			input:    "CASEBreak",
			snake:    "case_break",
			goSnake:  "case_break",
			SNAKE:    "CASE_BREAK",
			kebab:    "case-break",
			goKebab:  "case-break",
			KEBAB:    "CASE-BREAK",
			pascal:   "CaseBreak",
			goPascal: "CaseBreak",
			camel:    "caseBreak",
			goCamel:  "caseBreak",
			title:    "Case Break",
			goTitle:  "Case Break",
		},
		{
			input:    "ID",
			snake:    "id",
			goSnake:  "ID",
			SNAKE:    "ID",
			kebab:    "id",
			goKebab:  "ID",
			KEBAB:    "ID",
			pascal:   "Id",
			goPascal: "ID",
			camel:    "id",
			goCamel:  "id",
			title:    "Id",
			goTitle:  "ID",
		},
		{
			input:    "userID",
			snake:    "user_id",
			goSnake:  "user_ID",
			SNAKE:    "USER_ID",
			kebab:    "user-id",
			goKebab:  "user-ID",
			KEBAB:    "USER-ID",
			pascal:   "UserId",
			goPascal: "UserID",
			camel:    "userId",
			goCamel:  "userID",
			title:    "User Id",
			goTitle:  "User ID",
		},
		{
			input:    "JSON_blob",
			snake:    "json_blob",
			goSnake:  "JSON_blob",
			SNAKE:    "JSON_BLOB",
			kebab:    "json-blob",
			goKebab:  "JSON-blob",
			KEBAB:    "JSON-BLOB",
			pascal:   "JsonBlob",
			goPascal: "JSONBlob",
			camel:    "jsonBlob",
			goCamel:  "jsonBlob",
			title:    "Json Blob",
			goTitle:  "JSON Blob",
		},
		{
			input:    "HTTPStatusCode",
			snake:    "http_status_code",
			goSnake:  "HTTP_status_code",
			SNAKE:    "HTTP_STATUS_CODE",
			kebab:    "http-status-code",
			goKebab:  "HTTP-status-code",
			KEBAB:    "HTTP-STATUS-CODE",
			pascal:   "HttpStatusCode",
			goPascal: "HTTPStatusCode",
			camel:    "httpStatusCode",
			goCamel:  "httpStatusCode",
			title:    "Http Status Code",
			goTitle:  "HTTP Status Code",
		},
		{
			input:    "FreeBSD and SSLError are not golang initialisms",
			snake:    "free_bsd_and_ssl_error_are_not_golang_initialisms",
			goSnake:  "free_bsd_and_ssl_error_are_not_golang_initialisms",
			SNAKE:    "FREE_BSD_AND_SSL_ERROR_ARE_NOT_GOLANG_INITIALISMS",
			kebab:    "free-bsd-and-ssl-error-are-not-golang-initialisms",
			goKebab:  "free-bsd-and-ssl-error-are-not-golang-initialisms",
			KEBAB:    "FREE-BSD-AND-SSL-ERROR-ARE-NOT-GOLANG-INITIALISMS",
			pascal:   "FreeBsdAndSslErrorAreNotGolangInitialisms",
			goPascal: "FreeBsdAndSslErrorAreNotGolangInitialisms",
			camel:    "freeBsdAndSslErrorAreNotGolangInitialisms",
			goCamel:  "freeBsdAndSslErrorAreNotGolangInitialisms",
			title:    "Free Bsd And Ssl Error Are Not Golang Initialisms",
			goTitle:  "Free Bsd And Ssl Error Are Not Golang Initialisms",
		},
		{
			input:    "David's Computer",
			snake:    "david's_computer",
			goSnake:  "david's_computer",
			SNAKE:    "DAVID'S_COMPUTER",
			kebab:    "david's-computer",
			goKebab:  "david's-computer",
			KEBAB:    "DAVID'S-COMPUTER",
			pascal:   "David'sComputer",
			goPascal: "David'sComputer",
			camel:    "david'sComputer",
			goCamel:  "david'sComputer",
			title:    "David's Computer",
			goTitle:  "David's Computer",
		},
		{
			input:    "Ünicode support for Æthelred and Øyvind",
			snake:    "ünicode_support_for_æthelred_and_øyvind",
			goSnake:  "ünicode_support_for_æthelred_and_øyvind",
			SNAKE:    "ÜNICODE_SUPPORT_FOR_ÆTHELRED_AND_ØYVIND",
			kebab:    "ünicode-support-for-æthelred-and-øyvind",
			goKebab:  "ünicode-support-for-æthelred-and-øyvind",
			KEBAB:    "ÜNICODE-SUPPORT-FOR-ÆTHELRED-AND-ØYVIND",
			pascal:   "ÜnicodeSupportForÆthelredAndØyvind",
			goPascal: "ÜnicodeSupportForÆthelredAndØyvind",
			camel:    "ünicodeSupportForÆthelredAndØyvind",
			goCamel:  "ünicodeSupportForÆthelredAndØyvind",
			title:    "Ünicode Support For Æthelred And Øyvind",
			goTitle:  "Ünicode Support For Æthelred And Øyvind",
		},
		{
			input:    "http200",
			snake:    "http200",
			goSnake:  "http200",
			SNAKE:    "HTTP200",
			kebab:    "http200",
			goKebab:  "http200",
			KEBAB:    "HTTP200",
			pascal:   "Http200",
			goPascal: "Http200",
			camel:    "http200",
			goCamel:  "http200",
			title:    "Http200",
			goTitle:  "Http200",
		},
		{
			input:    "NumberSplittingVersion1.0r3",
			snake:    "number_splitting_version1.0r3",
			goSnake:  "number_splitting_version1.0r3",
			SNAKE:    "NUMBER_SPLITTING_VERSION1.0R3",
			kebab:    "number-splitting-version1.0r3",
			goKebab:  "number-splitting-version1.0r3",
			KEBAB:    "NUMBER-SPLITTING-VERSION1.0R3",
			pascal:   "NumberSplittingVersion1.0r3",
			goPascal: "NumberSplittingVersion1.0r3",
			camel:    "numberSplittingVersion1.0r3",
			goCamel:  "numberSplittingVersion1.0r3",
			title:    "Number Splitting Version1.0r3",
			goTitle:  "Number Splitting Version1.0r3",
		},
		{
			input:    "When you have a comma, odd results",
			snake:    "when_you_have_a_comma,_odd_results",
			goSnake:  "when_you_have_a_comma,_odd_results",
			SNAKE:    "WHEN_YOU_HAVE_A_COMMA,_ODD_RESULTS",
			kebab:    "when-you-have-a-comma,-odd-results",
			goKebab:  "when-you-have-a-comma,-odd-results",
			KEBAB:    "WHEN-YOU-HAVE-A-COMMA,-ODD-RESULTS",
			pascal:   "WhenYouHaveAComma,OddResults",
			goPascal: "WhenYouHaveAComma,OddResults",
			camel:    "whenYouHaveAComma,OddResults",
			goCamel:  "whenYouHaveAComma,OddResults",
			title:    "When You Have A Comma, Odd Results",
			goTitle:  "When You Have A Comma, Odd Results",
		},
		{
			input:    "Ordinal numbers work: 1st 2nd and 3rd place",
			snake:    "ordinal_numbers_work:_1st_2nd_and_3rd_place",
			goSnake:  "ordinal_numbers_work:_1st_2nd_and_3rd_place",
			SNAKE:    "ORDINAL_NUMBERS_WORK:_1ST_2ND_AND_3RD_PLACE",
			kebab:    "ordinal-numbers-work:-1st-2nd-and-3rd-place",
			goKebab:  "ordinal-numbers-work:-1st-2nd-and-3rd-place",
			KEBAB:    "ORDINAL-NUMBERS-WORK:-1ST-2ND-AND-3RD-PLACE",
			pascal:   "OrdinalNumbersWork:1st2ndAnd3rdPlace",
			goPascal: "OrdinalNumbersWork:1st2ndAnd3rdPlace",
			camel:    "ordinalNumbersWork:1st2ndAnd3rdPlace",
			goCamel:  "ordinalNumbersWork:1st2ndAnd3rdPlace",
			title:    "Ordinal Numbers Work: 1st 2nd And 3rd Place",
			goTitle:  "Ordinal Numbers Work: 1st 2nd And 3rd Place",
		},
		{
			input:    "BadUTF8\xe2\xe2\xa1",
			snake:    "bad_utf8_���",
			goSnake:  "bad_UTF8_���",
			SNAKE:    "BAD_UTF8_���",
			kebab:    "bad-utf8-���",
			goKebab:  "bad-UTF8-���",
			KEBAB:    "BAD-UTF8-���",
			pascal:   "BadUtf8���",
			goPascal: "BadUTF8���",
			camel:    "badUtf8���",
			goCamel:  "badUTF8���",
			title:    "Bad Utf8 ���",
			goTitle:  "Bad UTF8 ���",
		},
		{
			input:    "ID3_v2_3",
			snake:    "id3_v2_3",
			goSnake:  "ID3_v2_3",
			SNAKE:    "ID3_V2_3",
			kebab:    "id3-v2-3",
			goKebab:  "ID3-v2-3",
			KEBAB:    "ID3-V2-3",
			pascal:   "Id3V23",
			goPascal: "ID3V23",
			camel:    "id3V23",
			goCamel:  "id3V23",
			title:    "Id3 V2 3",
			goTitle:  "ID3 V2 3",
		},
		{
			input:    "ID3",
			snake:    "id3",
			goSnake:  "ID3",
			SNAKE:    "ID3",
			kebab:    "id3",
			goKebab:  "ID3",
			KEBAB:    "ID3",
			pascal:   "Id3",
			goPascal: "ID3",
			camel:    "id3",
			goCamel:  "id3",
			title:    "Id3",
			goTitle:  "ID3",
		},
		{
			input:    "IDENT3",
			snake:    "ident3",
			goSnake:  "ident3",
			SNAKE:    "IDENT3",
			kebab:    "ident3",
			goKebab:  "ident3",
			KEBAB:    "IDENT3",
			pascal:   "Ident3",
			goPascal: "Ident3",
			camel:    "ident3",
			goCamel:  "ident3",
			title:    "Ident3",
			goTitle:  "Ident3",
		},
		{
			input:    "LogRouterS3BucketName",
			snake:    "log_router_s3_bucket_name",
			goSnake:  "log_router_s3_bucket_name",
			SNAKE:    "LOG_ROUTER_S3_BUCKET_NAME",
			kebab:    "log-router-s3-bucket-name",
			goKebab:  "log-router-s3-bucket-name",
			KEBAB:    "LOG-ROUTER-S3-BUCKET-NAME",
			pascal:   "LogRouterS3BucketName",
			goPascal: "LogRouterS3BucketName",
			camel:    "logRouterS3BucketName",
			goCamel:  "logRouterS3BucketName",
			title:    "Log Router S3 Bucket Name",
			goTitle:  "Log Router S3 Bucket Name",
		},
		{
			input:    "PINEAPPLE",
			snake:    "pineapple",
			goSnake:  "pineapple",
			SNAKE:    "PINEAPPLE",
			kebab:    "pineapple",
			goKebab:  "pineapple",
			KEBAB:    "PINEAPPLE",
			pascal:   "Pineapple",
			goPascal: "Pineapple",
			camel:    "pineapple",
			goCamel:  "pineapple",
			title:    "Pineapple",
			goTitle:  "Pineapple",
		},
		{
			input:    "Int8Value",
			snake:    "int8_value",
			goSnake:  "int8_value",
			SNAKE:    "INT8_VALUE",
			kebab:    "int8-value",
			goKebab:  "int8-value",
			KEBAB:    "INT8-VALUE",
			pascal:   "Int8Value",
			goPascal: "Int8Value",
			camel:    "int8Value",
			goCamel:  "int8Value",
			title:    "Int8 Value",
			goTitle:  "Int8 Value",
		},
		{
			input:    "first.last",
			snake:    "first_last",
			goSnake:  "first_last",
			SNAKE:    "FIRST_LAST",
			kebab:    "first-last",
			goKebab:  "first-last",
			KEBAB:    "FIRST-LAST",
			pascal:   "FirstLast",
			goPascal: "FirstLast",
			camel:    "firstLast",
			goCamel:  "firstLast",
			title:    "First Last",
			goTitle:  "First Last",
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			output := data{
				input:    test.input,
				snake:    ToSnake(test.input),
				goSnake:  ToGoSnake(test.input),
				SNAKE:    ToSNAKE(test.input),
				kebab:    ToKebab(test.input),
				goKebab:  ToGoKebab(test.input),
				KEBAB:    ToKEBAB(test.input),
				pascal:   ToPascal(test.input),
				goPascal: ToGoPascal(test.input),
				camel:    ToCamel(test.input),
				goCamel:  ToGoCamel(test.input),
				title:    ToCase(test.input, TitleCase, ' '),
				goTitle:  ToGoCase(test.input, TitleCase, ' '),
			}
			if generate || test != output {
				line := fmt.Sprintf("%#v", output)
				line = strings.TrimPrefix(line, "strcase.data")
				line = strings.Replace(line, "\", ", "\",\n", -1)
				line = strings.Replace(line, "{", "{\n", -1)
				line = strings.Replace(line, "}", "\n},", -1)
				line = regexp.MustCompile("\"\n").ReplaceAllString(line, "\",\n")
				fmt.Println(line)
			}
			assertTrue(t, test == output)
		})
	}
}
