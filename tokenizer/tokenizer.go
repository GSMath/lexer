package tokenizer

import (
	"fmt"
	"unicode"
)

type Token struct {
	_type    int
	Symbolic string
	Numeric  float64
	Operator rune
}

func OperatorToken(token_str string, start int) (int, func() Token) {
	var operator rune
	operator = rune(token_str[start])
	switch operator {
	case '+':
	case '*':
	case '/':
	case '^':
	case '=':
	case rune("≤"[0]):
	case '<':
	case '>':
	case rune("≥"[0]):
		break
	default:
		return -1, func() Token {
			var t Token
			t._type = -1
			return t
		}
	}
	return start + 1, func() Token {
		var t Token
		t._type = 0
		t.Operator = operator
		return t
	}
}

func SymbolicToken(token_str string, start int) (int, func() Token) {
	runes := []rune(token_str)[start:]
	i := 0
	if unicode.IsLetter((runes[i])) == false {
		return 0, func() Token {
			var t Token
			t._type = -1
			return t
		}
	}
	for i = 1; i < len(runes); i++ {
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
	return start + i, func() Token {
		var t Token
		t._type = 1
		t.Symbolic = symbol
		return t
	}
}

func NumericToken(token_str string, start int) (int, func() Token) {
	runes := []rune(token_str)[start:]
	i := 0
	scientific := false
	decimal := false
	isComplex := false
	if unicode.IsNumber((runes[i])) == false &&
		runes[i] != '-' {
		return 0, func() Token {
			var t Token
			t._type = -1
			return t
		}
	}
	for i = 1; i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) == true {
			continue
		}
		if runes[i] == '.' && decimal == false {
			decimal = true
			continue
		}
		if unicode.ToLower(runes[i]) == 'e' && scientific == false {
			scientific = true
			continue
		}
		if runes[i] == '-' && unicode.ToLower(runes[i-1]) == 'e' {
			continue
		}
		if runes[i] == 'i' {
			isComplex = true
			break
		}
		break
	}
	var numeric float64
	fmt.Sscanf(string(token_str[start:i]), "%f", &numeric)
	if isComplex == false {
		return start + i, func() Token {
			var t Token
			t._type = 2
			t.Numeric = numeric
			return t
		}
	} else {
		return start + i, func() Token {
			var t Token
			t._type = 3
			t.Numeric = numeric
			return t
		}
	}

}
