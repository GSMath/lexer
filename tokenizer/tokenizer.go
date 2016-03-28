package tokenizer

import (
	"fmt"
	"unicode"
)

const (
	Operator      = iota
	Symbol        = iota
	RealNumeric   = iota
	ImagNumeric   = iota
	Subexpression = iota
	Empty         = -1
)

const (
	GetToken     = iota
	GetSubtokens = iota
)

type Token struct {
	_type         int
	Symbolic      string
	Numeric       float64
	Operator      rune
	Subexpression string
}

func OperatorToken(runes []rune, start int) (int, func(int) []Token) {
	var operator rune
	operator = runes[start]
	switch operator {
	case '+':
	case '*':
	case '/':
	case '^':
	case '=':
	case []rune("≤")[0]:
	case '<':
	case '>':
	case []rune("≥")[0]:
		break
	default:
		return -1, func(action int) []Token {
			var t Token
			t._type = Empty
			return []Token{t}
		}
	}
	return 1, func(action int) []Token {
		var t Token
		if action == GetToken {
			t._type = Operator
			t.Operator = operator
		} else {
			t._type = Empty
		}
		return []Token{t}
	}
}

func SymbolicToken(runes []rune, start int) (int, func(int) []Token) {
	runes = runes[start:]
	i := 0
	if unicode.IsLetter((runes[i])) == false {
		return 0, func(action int) []Token {
			var t Token
			t._type = Empty
			return []Token{t}
		}
	}
	for i = 1; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) == true {
			continue
		}
		if unicode.IsNumber(runes[i]) == true {
			continue
		}
		if runes[i] == '_' {
			continue
		}
		break
	}
	symbol := string(runes[:i])
	return i, func(action int) []Token {
		var t Token
		if action == GetToken {
			t._type = Symbol
			t.Symbolic = symbol
		} else {
			t._type = Empty
		}
		return []Token{t}
	}
}

func NumericToken(runes []rune, start int) (int, func(int) []Token) {
	runes = runes[start:]
	i := 0
	scientific := false
	decimal := false
	isComplex := false
	if unicode.IsNumber((runes[i])) == false &&
		runes[i] != '-' {
		return 0, func(action int) []Token {
			var t Token
			t._type = Empty
			return []Token{t}
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
	fmt.Sscanf(string(runes[:i]), "%f", &numeric)
	if isComplex == false {
		return i, func(action int) []Token {
			var t Token
			if action == GetToken {
				t._type = RealNumeric
				t.Numeric = numeric
			} else {
				t._type = Empty
			}
			return []Token{t}
		}
	} else {
		return i + 1, func(action int) []Token {
			var t Token
			if action == GetToken {
				t._type = ImagNumeric
				t.Numeric = numeric
			} else {
				t._type = Empty
			}
			return []Token{t}
		}
	}
}

func SubexpressionToken(runes []rune, start int) (int, func(int) []Token) {
	runes = runes[start:]
	i := 0
	var openingBracket rune
	var closingBracket rune
	level := 0
	if runes[i] != '(' && runes[i] != '[' {
		return 0, func(action int) []Token {
			var t Token
			t._type = Empty
			return []Token{t}
		}
	}
	openingBracket = runes[i]
	if runes[i] == '(' {
		closingBracket = ')'
	} else {
		closingBracket = ']'
	}
	for i = 1; i < len(runes); i++ {
		if runes[i] == closingBracket {
			if level > 0 {
				level--
				continue
			}
			i++
			return i, func(action int) []Token {
				var t Token
				var T []Token
				if action == GetToken {
					t._type = Subexpression
					t.Subexpression = string(runes[1 : i-1])
					return []Token{t}
				} else {
					generators := TokenizeString(string(runes[1 : i-1]))
					for j := 0; j < len(generators); j++ {
						k := generators[j](GetToken)
						T = append(T, k[0])
					}
					return T
				}
			}
		}
		if runes[i] == openingBracket {
			level++
		}
	}
	return 0, func(action int) []Token {
		var t Token
		t._type = Empty
		return []Token{t}
	}
}

func TokenizeString(expression string) []func(int) []Token {
	runes := []rune(expression)
	generators := [](func(int) []Token){}
	tokenizers := []func([]rune, int) (int, func(int) []Token){OperatorToken, SymbolicToken, NumericToken, SubexpressionToken}
	var end int
	var token_func func(int) []Token
	var i int
	for i = 0; i < len(runes); {
		if unicode.IsSpace(runes[i]) == true {
			i++
			continue
		}
		for j := 0; j < len(tokenizers); j++ {
			end, token_func = tokenizers[j](runes, i)
			if end > 0 {
				i += end
				generators = append(generators, token_func)
				break
			}
		}
		if end == 0 {
			return []func(int) []Token{}
		}

	}
	return generators
}
