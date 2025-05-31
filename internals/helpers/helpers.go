package helpers

import "strings"

func CleanInput(text string) (output []string) {
	inputStrings := strings.Fields(strings.ToLower(text))
	for _, str := range inputStrings {
		output = append(output, strings.Trim(str, " "))
	}
	return output
}
