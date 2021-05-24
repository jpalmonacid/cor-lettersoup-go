package lettersoup

import (
	"fmt"
	"testing"
)

// func TestGivenInvalidSizeSoupReturnsMinus1AndError(t *testing.T) {
// 	want := -1
// 	count, err := CountOIEMatchesInSoup([]string{"OI", "XX", "OY"})
// 	if count != want || err == nil {
// 		t.Fatalf(`CountOIEMatchesInSoup("", {"OIE", "XXX", "OYE"}) = ( %d, "%v" ), want ( %d, error )`,
// 			count,
// 			err,
// 			want,
// 		)
// 	}
// }

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
		testDescription := fmt.Sprintf("Test case %v", testCase)
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
