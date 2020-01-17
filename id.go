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
		return value, fmt.Errorf("Invalid value. Must be [a-z or A-Z]")
	}

	result := string(incrementIdStartingAt(0, bytes))
	return result, nil
}

func incrementIdStartingAt(index int, value []byte) []byte {
	sourceLength := len(value)
	letter := nextLetter(value[index])
	value[index] = letter

	for letter == capitalA {
		index = index + 1

		if sourceLength == index {
			value = append(value, capitalA)
			break
		} else {
			letter = nextLetter(value[index])
			value[index] = letter
		}
	}

	return value
}

func nextLetter(value byte) byte {
	nextChar := value + 1

	if nextChar > capitalZ {
		return capitalA
	}

	return nextChar
}
