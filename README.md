# stringcases

[![Build Status](https://travis-ci.org/parithiban/stringcases.svg?branch=master)](https://travis-ci.org/parithiban/stringcases)
[![codecov.io](https://codecov.io/github/parithiban/stringcases/coverage.svg?branch=master)](https://codecov.io/github/parithiban/stringcases?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/parithiban/stringcases)](https://goreportcard.com/report/github.com/parithiban/stringcases)


[![GitHub contributors](https://img.shields.io/github/contributors/parithiban/stringcases.svg?style=plastic&color=blue)](https://GitHub.com/parithiban/stringcases/graphs/contributors/)
![Last Commit](https://img.shields.io/github/last-commit/parithiban/stringcases.svg?style=plastic)

This is a string package written in go which is used to convert the strings to different case styles. The commonly used strategies are: camel case, pascal case, snake case, and kebab case.

## Installation

This requires Go version 1.11 or later.

> go get -u github.com/parithiban/stringcases

## Usage

The following are some of the methods that can be used from the package.

```code
ConvertToSnake  Convert a string from any format to snake_case.
ConvertToKebab  Convert a string from any format to kebab-case.
ConvertToPascal Convert a string from any format to PascalCase.
ConvertToCamel  Convert a string from any format to camelCase.
```

## Example

```code
ConvertToSnake("ConvertToSnakeCase") =>  convert_to_snake_case
ConvertToKebab("convert_to_kebab_case") => convert-to-kebab-case.
ConvertToPascal("convert-to-pascal-case") => ConvertToPascalCase.
ConvertToCamel("convert to camel case") => convertToCamelCase.
```
