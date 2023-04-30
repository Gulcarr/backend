package unification

import (
	"awesomeProject/structs"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareStrings(t *testing.T) {
	testTable := []struct {
		a, b            string
		flagi, expected bool
	}{
		{
			a:        "abcd",
			b:        "abcd",
			flagi:    false,
			expected: true,
		},
		{
			a:        "abce",
			b:        "abcd",
			flagi:    false,
			expected: false,
		},
		{
			a:        "aBcd",
			b:        "abcd",
			flagi:    false,
			expected: false,
		},
		{
			a:        "aBcd",
			b:        "abcd",
			flagi:    true,
			expected: true,
		},
		{
			a:        "aBce",
			b:        "abcd",
			flagi:    true,
			expected: false,
		},
		{
			a:        "",
			b:        "abcd",
			flagi:    false,
			expected: false,
		},
		{
			a:        "abcgfjgfd",
			b:        "abcd",
			flagi:    true,
			expected: false,
		},
	}
	for _, testCase := range testTable {
		result := CompareStrings(testCase.a, testCase.b, testCase.flagi)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected: %t, got: %t",
			testCase.expected, result))
	}
}
func TestParseWithIgnr(t *testing.T) {
	testTable := []struct {
		toIgnr     []string
		fcnt, ccnt int
		flagi      bool
		expected   []string
	}{
		{
			toIgnr: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			fcnt:  1,
			ccnt:  0,
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"We love music.", 3},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love music of Kartik.", 2},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			toIgnr: []string{
				"A love music.",
				"B love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			fcnt:  0,
			ccnt:  1,
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"A love music.", 3},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love music of Kartik.", 1},
				stringWithRepCnt{"We love music of Kartik.", 1},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			toIgnr: []string{
				"Of course Adi loves music.",
				"It's true Bob loves music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			fcnt:  2,
			ccnt:  3,
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 2},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 2},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			toIgnr: []string{
				"Of course Adi loves music.",
				"Of course Adi lovEs muSic.",
				"It's true Bob LOVES music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			fcnt:  2,
			ccnt:  3,
			flagi: true,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 3},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 2},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			toIgnr: []string{
				"Of course Adi loves music.",
				"Of course Adi lovEs muSic.",
				"It's true Bob LOVES music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			fcnt:  2,
			ccnt:  3,
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 1},
				stringWithRepCnt{"Of course Adi lovEs muSic.", 1},
				stringWithRepCnt{"It's true Bob LOVES music.", 1},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 2},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			toIgnr: []string{
				"Of course Adi loves music.",
				"It's true Bob loves music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			fcnt:  2,
			ccnt:  3,
			flagi: true,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 2},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 2},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
	}
	for _, testCase := range testTable {
		result := parseWithIgnr(testCase.toIgnr, structs.FlagSet{"c", testCase.fcnt, testCase.ccnt, testCase.flagi})
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected: %s ..., got: %s ...",
			testCase.expected[0], result[0]))
	}
}
func TestParseNoIgnr(t *testing.T) {
	testTable := []struct {
		args     []string
		flagi    bool
		expected []string
	}{
		/*{
			args: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				{"We love music.", 1},
				{"I love music.", 1},
				stringWithRepCnt{"They love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love music of Kartik.", 1},
				stringWithRepCnt{"We love music of Kartik.", 1},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			args: []string{
				"A love music.",
				"B love music.",
				"C love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"A love music.", 1},
				stringWithRepCnt{"B love music.", 1},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love music of Kartik.", 1},
				stringWithRepCnt{"We love music of Kartik.", 1},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},*/
		{
			args: []string{
				"Of course Adi loves music.",
				"Of course Adi lovEs musIc.",
				"It's true Bob LOVES music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			flagi: true,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 2},
				stringWithRepCnt{"It's true Bob LOVES music.", 1},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 1},
				stringWithRepCnt{"We love per music of Kartik.", 1},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		/*{
			args: []string{
				"Of course Adi loves music.",
				"Of course Adi lovEs muSic.",
				"It's true Bob LOVES music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			flagi: false,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 1},
				stringWithRepCnt{"Of course Adi lovEs muSic.", 1},
				stringWithRepCnt{"It's true Bob LOVES music.", 1},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 1},
				stringWithRepCnt{"We love per music of Kartik.", 1},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},
		{
			args: []string{
				"Of course Adi loves music.",
				"It's true Bob loves music.",
				"C love music.",
				"",
				"I love and music of Kartik.",
				"We love per music of Kartik.",
				"Thanks.",
			},
			flagi: true,
			expected: chooseToPrint([]stringWithRepCnt{
				stringWithRepCnt{"Of course Adi loves music.", 1},
				stringWithRepCnt{"It's true Bob loves music.", 1},
				stringWithRepCnt{"C love music.", 1},
				stringWithRepCnt{"", 1},
				stringWithRepCnt{"I love and music of Kartik.", 1},
				stringWithRepCnt{"We love per music of Kartik.", 1},
				stringWithRepCnt{"Thanks.", 1},
			}, "c"),
		},*/
	}
	for _, testCase := range testTable {
		result := parseNoIgnr(testCase.args, testCase.flagi, "c")
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected: %s..., got: %s...",
			testCase.expected[0], result[0]))
	}
}
