package lettersoup

import (
	"errors"
	"fmt"
)

const (
	word                = "OIE"
	soupSizeMaxBoundary = 100
)

func countOIEMatchesInSoup(soup []string) (count int, err error) {
	wordLength, rowsCount, colsCount := len(word), len(soup), 0
	if rowsCount != 0 {
		colsCount = len(soup[0])
	}
	if rowsCount == 0 || rowsCount > soupSizeMaxBoundary || colsCount > soupSizeMaxBoundary {
		message := fmt.Sprintf(
			"Soup size {W: %d, H: %d} is out of bounds. It should not be empty"+
				"nor have any dimension greater than %d",
			rowsCount,
			colsCount,
			soupSizeMaxBoundary,
		)
		return -1, errors.New(message)
	}
	count = 0
	for i := 0; i < rowsCount; i += 1 {
		for j := 0; j < colsCount; j += 1 {
			if soup[i][j] == word[0] {
				// Horizontal Left-Right
				if colsCount-j >= wordLength && soup[i][j:j+wordLength] == word {
					count += 1
				}
				// Horizontal Right-Left
				if j-(wordLength-1) >= 0 && string(soup[i][j])+string(soup[i][j-1])+string(soup[i][j-2]) == word {
					count += 1
				}
				// Vertical Top-Bottom
				if rowsCount-i >= wordLength && string(soup[i][j])+string(soup[i+1][j])+string(soup[i+2][j]) == word {
					count += 1
				}
				// Vertical Bottom-Top
				if i-(wordLength-1) >= 0 && string(soup[i][j])+string(soup[i-1][j])+string(soup[i-2][j]) == word {
					count += 1
				}
				// Diagonal LeftRight-TopBottom
				if j+(wordLength-1) < colsCount && i+(wordLength-1) < rowsCount && string(soup[i][j])+string(soup[i+1][j+1])+string(soup[i+2][j+2]) == word {
					count += 1
				}
				// Diagonal RightLeft-TopBottom
				if j-(wordLength-1) >= 0 && i+(wordLength-1) < rowsCount && string(soup[i][j])+string(soup[i+1][j-1])+string(soup[i+2][j-2]) == word {
					count += 1
				}
				// Diagonal LeftRight-BottomTop
				if j+(wordLength-1) < colsCount && i-(wordLength-1) >= 0 && string(soup[i][j])+string(soup[i-1][j+1])+string(soup[i-2][j+2]) == word {
					count += 1
				}
				// Diagonal RightLeft-BottomTop
				if j-(wordLength-1) >= 0 && i-(wordLength-1) >= 0 && string(soup[i][j])+string(soup[i-1][j-1])+string(soup[i-2][j-2]) == word {
					count += 1
				}
			}
		}
	}

	return count, nil
}
