package lettersoup

import (
	"fmt"
	"testing"
)

func TestGivenInvalidSizeSoupReturnsMinus1AndError(t *testing.T) {
	testCases := []struct {
		desc string
		soup []string
	}{
		{"soup cols size = 101", []string{makeString(101)}},
		{"soup rows size = 101", make([]string, 101)},
		{"empty soup", []string{}},
	}

	for _, testCase := range testCases {
		t.Run("Case "+testCase.desc, func(t *testing.T) {
			want := -1
			gottenCount, err := countOIEMatchesInSoup(testCase.soup)
			if gottenCount != want || err == nil {
				t.Fatalf(`countOIEMatchesInSoup(invalidSoup) = ( %d, "%v" ), want ( %d, error )`,
					gottenCount,
					err,
					want,
				)
			}
		})
	}
}

func makeString(size int) (word string) {
	for i := 0; i <= size; i += 1 {
		word += "A"
	}
	return
}

func TestGivenValidSoupItReturnsCount(t *testing.T) {
	testCases := []struct {
		soup        []string
		wantedCount int
	}{
		{[]string{"OIE", "IIX", "EXE"}, 3},
		{[]string{"EIOIEIOEIO"}, 4},
		{[]string{"EAEAE", "AIIIA", "EIOIE", "AIIIA", "EAEAE"}, 8},
		{[]string{"OX", "IO", "EX", "II", "OX", "IE", "EX"}, 3},
		{[]string{"E"}, 0},
	}
	for _, testCase := range testCases {
		testDescription := fmt.Sprintf("Case soup = %v, wantedCount = %d ", testCase.soup, testCase.wantedCount)
		t.Run(testDescription, func(t *testing.T) {
			gottenCount, err := countOIEMatchesInSoup(testCase.soup)
			if gottenCount != testCase.wantedCount || err != nil {
				t.Fatalf(`countOIEMatchesInSoup() = ( %d, "%v" ), want ( %d, nil )`,
					gottenCount,
					err,
					testCase.wantedCount,
				)
			}
		})
	}
}
