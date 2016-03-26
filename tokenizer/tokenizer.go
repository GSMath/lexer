package tokenizer

import (
	"fmt"
	"unicode"
)

type Token struct {
	_type     int
	Symbolic  func() string
	Numerical func() complex128
}

func OperatorToken(token_str string, start int) (int, int, func() rune) {
	var operator rune
	operator = []rune(token_str)[start]
	return 0, start + 1, func() rune {
		return operator
	}
}

func SymbolToken(token_str string, start int) (int, int, func() string) {
	runes := []rune(token_str)[start:]
	i := 0
	for i = 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) == true {
			continue
		}
		if unicode.IsNumber(runes[i]) == true {
			continue
		}
		if unicode.IsSymbol(runes[i]) == true {
			continue
		}
		if runes[i] == '_' {
			continue
		}
		break
	}
	symbol := string(runes[start : start+i])
	return 1, start + i, func() string {
		return symbol
	}
}

func NumericToken(token_str string, start int) (int, int, func() complex128) {
	i := 0
	var numeric complex128
	fmt.Sscanf(token_str, "%f", &numeric)
	return 2, start + i, func() complex128 {
		return numeric
	}
}
