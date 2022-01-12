package tutorial_testify_1

import (
	"fmt"
	"unicode"
)

func IsAllLowerCase(input string) (bool, error) {

	for _, ch := range input {
		fmt.Println(ch, "----->")
		if unicode.IsUpper(ch) {
			return false, nil
		}
	}
	return true, nil
}
