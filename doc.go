// Copyright 2019 The parithiban. All rights reserved.
// license that can be found in the LICENSE file.

/*
Package stringcases used to convert a string to different case styles.

Usage:

	import "github.com/parithiban/stringcases"

Call the respective string convertion methods to convert the cases. For example:

	func main() {
		camelCase  := stringcases.ConvertToCamel("convert to camel case")   //output: convertToCamelCase
		snake_case := stringcases.ConvertToSnake("convert to snake case")   //output: convert_to_snake_case
		kebab-case := stringcases.ConvertToKebab("convert to kebab case")   //output: convert-to-kebab-case
		PascalCase := stringcases.ConvertToPascal("convert to pascal case") //output: ConvertToPascalCase
		lowerfirst := stringcases.FirstLetterLower("Convert To Lower Case") //output: convert To Lower Case
	}
*/
package stringcases
