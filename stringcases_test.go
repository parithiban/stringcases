package stringcases

import "testing"

func Test_ConvertToCamel(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "convert to camel case with spaces", input: "convert to camel case", expected: "convertToCamelCase"},
		{name: "convert to camel case with dashes", input: "convert-to-camel-case", expected: "convertToCamelCase"},
		{name: "convert to camel case with underscores", input: "convert_to_camel_case", expected: "convertToCamelCase"},
		{name: "convert to camel case with mixed", input: "convert-to_camel case", expected: "convertToCamelCase"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := ConvertToCamel(tc.input)

			if output != tc.expected {
				t.Fatalf("Output should be %s, but got %s", tc.expected, output)
			}
		})
	}
}

func Test_ConvertToPascal(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "convert to pascal case with spaces", input: "convert to pascal case", expected: "ConvertToPascalCase"},
		{name: "convert to pascal case with dashes", input: "convert-to-pascal-case", expected: "ConvertToPascalCase"},
		{name: "convert to pascal case with underscores", input: "convert_to_pascal_case", expected: "ConvertToPascalCase"},
		{name: "convert to pascal case with mixed", input: "convert-to_pascal case", expected: "ConvertToPascalCase"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := ConvertToPascal(tc.input)

			if output != tc.expected {
				t.Fatalf("Output should be %s, but got %s", tc.expected, output)
			}
		})
	}
}

func Test_FirstLetterLower(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "convert first letter lower case  with spaces", input: "Convert To Lower Case", expected: "convert To Lower Case"},
		{name: "convert first letter lower case with dashes", input: "Convert-to-Lower-case", expected: "convert-to-Lower-case"},
		{name: "convert first letter lower case  with underscores", input: "Convert_to_Lower_case", expected: "convert_to_Lower_case"},
		{name: "convert first letter lower case  with mixed", input: "Convert-to_Lower case", expected: "convert-to_Lower case"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := FirstLetterLower(tc.input)

			if output != tc.expected {
				t.Fatalf("Output should be %s, but got %s", tc.expected, output)
			}
		})
	}
}

func Test_ConvertToSnake(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "convert to snake case with input as pascal case", input: "ConvertToSnakeCase", expected: "convert_to_snake_case"},
		{name: "convert to snake case with dashes", input: "convert-to-snake-case", expected: "convert_to_snake_case"},
		{name: "convert to snake case with underscores", input: "convert_to_snake_case", expected: "convert_to_snake_case"},
		{name: "convert to snake case with mixed", input: "convert-to_snake case", expected: "convert_to_snake_case"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := ConvertToSnake(tc.input)

			if output != tc.expected {
				t.Fatalf("Output should be %s, but got %s", tc.expected, output)
			}
		})
	}
}

func Test_ConvertToKebab(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "convert to kebab case with input as pascal case", input: "ConvertToKebabCase", expected: "convert-to-kebab-case"},
		{name: "convert to snake case with dashes", input: "convert-to-kebab-case", expected: "convert-to-kebab-case"},
		{name: "convert to snake case with underscores", input: "convert_to_kebab_case", expected: "convert-to-kebab-case"},
		{name: "convert to snake case with mixed", input: "convert-to_kebab case", expected: "convert-to-kebab-case"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := ConvertToKebab(tc.input)

			if output != tc.expected {
				t.Fatalf("Output should be %s, but got %s", tc.expected, output)
			}
		})
	}
}
