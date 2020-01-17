package letter

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	capitalA = 65
	capitalZ = 90
)

func NextId(value string) (string, error) {
	bytes := []byte(strings.ToUpper(value))
	matched, _ := regexp.Match(`^[A-Z]+$`, bytes)

	if !matched {
		return value, fmt.Errorf("Invalid Id. Must be [a-z or A-Z]")
	}

	result := string(incrementIdStartingAt(0, bytes))
	return result, nil
}

func incrementIdStartingAt(index int, value []byte) []byte {
	sourceLength := len(value)
	letter := nextLetterAt(index, value)
	value[index] = letter

	for letter == capitalA {
		index = index + 1

		if sourceLength == index {
			value = append(value, capitalA)
			break
		} else {
			letter = nextLetterAt(index, value)
			value[index] = letter
		}
	}

	return value
}

func nextLetterAt(index int, value []byte) byte {
	candidateChar := value[index]
	nextChar := candidateChar + 1

	if nextChar > capitalZ {
		return capitalA
	}

	return nextChar
}
