package tokenizer

import (
	"fmt"
	"unicode"
)

var SupportedOperators string = DefaultOperators

func OperatorToken(runes []rune, start int) (int, Token) {
	var allOperators []rune
	var operator rune
	operator = runes[start]
	allOperators = []rune(SupportedOperators)
	count := 1
	isOperator := false
	switch operator {
	case '>':
		if runes[start+1] == '=' {
			operator = []rune("≥")[0]
			count++
		}
	case '<':
		if runes[start+1] == '=' {
			operator = []rune("≤")[0]
			count++
		}
	case '!':
		if runes[start+1] == '=' {
			operator = []rune("≠")[0]
			count++
		}
	case '-':
		if runes[start+1] == '+' {
			operator = []rune("∓")[0]
			count++
		}
	case '+':
		if runes[start+1] == '-' {
			operator = []rune("±")[0]
			count++
		}
	case '=':
		if runes[start+1] == '=' {
			operator = []rune("≡")[0]
			count++
		}
		if runes[start+1] == '!' &&
			runes[start+1] == '=' {
			operator = []rune("≢")[0]
			count += 2
		}
	case '*':
		operator = []rune("•")[0]
	case []rune("÷")[0]:
		operator = []rune("⁄")[0]
	case '/':
		operator = []rune("⁄")[0]

	}
	for i, valid := range allOperators {
		if valid == operator {
			isOperator = true
			i++
			break
		}
	}
	if isOperator != true {
		return 0, Token{TokenType: Empty}
	}
	return count, Token{
		TokenType: Operator,
		Operator:  operator,
	}
}

func SymbolicToken(runes []rune, start int) (int, Token) {
	var t Token = Token{TokenType: Empty}
	var symbol string
	runes = runes[start:]
	i := 0
	if unicode.IsLetter((runes[i])) == false {
		goto EXIT
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
	symbol = string(runes[:i])
	t = Token{
		TokenType: Symbol,
		Symbolic:  symbol,
	}
EXIT:
	return i, t
}

func NumericToken(runes []rune, start int) (int, Token) {
	var t Token = Token{TokenType: Empty}
	var length int = 0
	var numeric float64
	runes = runes[start:]
	scientific := false
	decimal := false
	isComplex := false
	if unicode.IsNumber((runes[0])) == false &&
		runes[0] != '-' {
		goto EXIT
	}
	for length = 1; length < len(runes); length++ {
		if unicode.IsNumber(runes[length]) == true {
			continue
		}
		if runes[length] == '.' && decimal == false {
			decimal = true
			continue
		}
		if unicode.ToLower(runes[length]) == 'e' && scientific == false {
			scientific = true
			continue
		}
		if runes[length] == '-' && unicode.ToLower(runes[length-1]) == 'e' {
			continue
		}
		if runes[length] == 'i' {
			isComplex = true
			break
		}
		break
	}
	fmt.Sscanf(string(runes[:length]), "%f", &numeric)
	if isComplex == false {
		t = Token{
			TokenType: RealNumeric,
			Numeric:   numeric,
		}
	} else {
		length++
		t = Token{
			TokenType: ImagNumeric,
			Numeric:   numeric,
		}
	}
EXIT:
	return length, t
}

func SubexpressionToken(runes []rune, start int) (int, Token) {
	var t Token = Token{TokenType: Empty}
	var length int = 0
	runes = runes[start:]
	i := 0
	var openingBracket rune
	var closingBracket rune
	level := 0
	if runes[i] != '(' && runes[i] != '[' {
		goto EXIT
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
			t.TokenType = Subexpression
			t.Subexpression = string(runes[1 : i-1])
			length = i
			goto EXIT
		}
		if runes[i] == openingBracket {
			level++
		}
	}
	length = 0
	t.TokenType = Empty
EXIT:
	return length, t
}

func TokenizeString(expression string) []Token {
	runes := []rune(expression)
	tokens := []Token{}
	tokenizers := []func([]rune, int) (int, Token){OperatorToken, NumericToken, SymbolicToken, SubexpressionToken}
	var end int
	var token Token
	var i int
	for i = 0; i < len(runes); {
		if unicode.IsSpace(runes[i]) == true {
			i++
			continue
		}
		for j := 0; j < len(tokenizers); j++ {
			end, token = tokenizers[j](runes, i)
			if end > 0 {
				i += end
				tokens = append(tokens, token)
				break
			}
		}
		if end == 0 {
			return []Token{}
		}
	}
	return tokens
}
