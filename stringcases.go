package stringcases

import (
	"regexp"
	"strings"
	"unicode"
)

// ConvertToSnake converts a string to snake_case
// Returns a string
func ConvertToSnake(s string) string {
	return CommonSnakeKebabCase(s, "_", "-")
}

// ConvertToKebab converts a string to kebab-case
// Returns a string
func ConvertToKebab(s string) string {
	return CommonSnakeKebabCase(s, "-", "_")
}

// CommonSnakeKebabCase common method to convert the given string to snake/kebab case
// Returns a string
func CommonSnakeKebabCase(s string, seperator string, identifier string) string {
	var output strings.Builder
	runeContains := identifier + " "

	for _, character := range s {
		if unicode.IsUpper(character) {
			output.WriteString(seperator + string(character))
		} else if strings.ContainsRune(runeContains, character) {
			output.WriteString(seperator)
		} else {
			output.WriteString(string(character))
		}
	}

	return strings.ToLower(strings.TrimPrefix(output.String(), seperator))
}

// ConvertToPascal converts a string to PascalCase
// Returns a string
func ConvertToPascal(s string) string {
	r := regexp.MustCompile("[_]|[-]|[\\s]")
	split := r.Split(s, -1)
	var output strings.Builder

	for _, content := range split {
		output.WriteString(strings.Title(content))
	}

	return output.String()
}

// ConvertToCamel converts a string to camelCase
// Returns a string
func ConvertToCamel(s string) string {
	return FirstLetterLower(ConvertToPascal(s))
}

// FirstLetterLower Change the first letter to caps of a string
// Returns a string
func FirstLetterLower(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
