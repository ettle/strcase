package strcase

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestCaserAll(t *testing.T) {
	c := NewCaser(true, nil, nil)

	type data struct {
		input  string
		snake  string
		SNAKE  string
		kebab  string
		KEBAB  string
		pascal string
		camel  string
		title  string
	}
	for _, test := range []data{
		{
			input:  "Hello world!",
			snake:  "hello_world!",
			SNAKE:  "HELLO_WORLD!",
			kebab:  "hello-world!",
			KEBAB:  "HELLO-WORLD!",
			pascal: "HelloWorld!",
			camel:  "helloWorld!",
			title:  "Hello World!",
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			output := data{
				input:  test.input,
				snake:  c.ToSnake(test.input),
				SNAKE:  c.ToSNAKE(test.input),
				kebab:  c.ToKebab(test.input),
				KEBAB:  c.ToKEBAB(test.input),
				pascal: c.ToPascal(test.input),
				camel:  c.ToCamel(test.input),
				title:  c.ToCase(test.input, TitleCase, ' '),
			}
			assert.Equal(t, test, output)
		})
	}
}

func TestNewCaser(t *testing.T) {
	t.Run("Has defaults when unspecified", func(t *testing.T) {
		c := NewCaser(true, nil, nil)
		assert.Equal(t, golintInitialisms, c.initialisms)
		assert.NotNil(t, c.splitFn)
	})
	t.Run("Merges", func(t *testing.T) {
		c := NewCaser(true, map[string]bool{"SSL": true, "HTML": false}, nil)
		assert.NotEqual(t, golintInitialisms, c.initialisms)
		assert.True(t, c.initialisms["UUID"])
		assert.True(t, c.initialisms["SSL"])
		assert.False(t, c.initialisms["HTML"])
		assert.NotNil(t, c.splitFn)
	})

	t.Run("No Go initialisms", func(t *testing.T) {
		c := NewCaser(false, map[string]bool{"SSL": true, "HTML": false}, NewSplitFn([]rune{' '}))
		assert.NotEqual(t, golintInitialisms, c.initialisms)
		assert.False(t, c.initialisms["UUID"])
		assert.True(t, c.initialisms["SSL"])
		assert.False(t, c.initialisms["HTML"])
		assert.Equal(t, "hTml with SSL", c.ToCase("hTml with SsL", Original, ' '))
		assert.NotNil(t, c.splitFn)
	})

	t.Run("Preserve number formatting", func(t *testing.T) {
		c := NewCaser(
			false,
			map[string]bool{"SSL": true, "HTML": false},
			NewSplitFn(
				[]rune{'*', '.', ','},
				SplitCase,
				SplitAcronym,
				PreserveNumberFormatting,
			))
		assert.NotEqual(t, golintInitialisms, c.initialisms)
		assert.Equal(t, "http200", c.ToSnake("http200"))
		assert.Equal(t, "VERSION2.3_R3_8A_HTTP_ERROR_CODE", c.ToSNAKE("version2.3R3*8a,HTTPErrorCode"))
	})

	t.Run("Preserve number formatting and split before and after number", func(t *testing.T) {
		c := NewCaser(
			false,
			map[string]bool{"SSL": true, "HTML": false},
			NewSplitFn(
				[]rune{'*', '.', ','},
				SplitCase,
				SplitAcronym,
				PreserveNumberFormatting,
				SplitBeforeNumber,
				SplitAfterNumber,
			))
		assert.Equal(t, "http_200", c.ToSnake("http200"))
		assert.Equal(t, "VERSION_2.3_R_3_8_A_HTTP_ERROR_CODE", c.ToSNAKE("version2.3R3*8a,HTTPErrorCode"))
	})

	t.Run("Skip non letters", func(t *testing.T) {
		c := NewCaser(
			false,
			nil,
			func(prec, curr, next rune) SplitAction {
				if unicode.IsNumber(curr) {
					return Noop
				} else if unicode.IsSpace(curr) {
					return SkipSplit
				}
				return Skip
			})
		assert.Equal(t, "", c.ToSnake(""))
		assert.Equal(t, "1130_23_2009", c.ToCase("DateTime: 11:30 AM May 23rd, 2009", Original, '_'))
	})
}
