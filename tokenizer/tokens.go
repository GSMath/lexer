package tokenizer

const (
	Wildcard      = iota
	Numeric       = iota
	Expression    = iota
	Symbol        = iota
	RealNumeric   = iota
	ImagNumeric   = iota
	Subexpression = iota
	Operator      = iota
	Empty         = -1
)

const DefaultOperators = "-+*/^=≠<>≤≥∏∑∫√≈⁄÷∓±≡≢•∂"

type Token struct {
	TokenType     int
	Symbolic      string
	Numeric       float64
	Operator      rune
	Subexpression string
}

func (lhs Token) Equivalent(rhs Token) bool {
	if lhs.TokenType == Wildcard ||
		rhs.TokenType == Wildcard {
		return true
	}
	if lhs.TokenType == Numeric {
		switch rhs.TokenType {
		case RealNumeric:
			return true
		case ImagNumeric:
			return true
		}
	}
	if rhs.TokenType == Numeric {
		switch lhs.TokenType {
		case RealNumeric:
			return true
		case ImagNumeric:
			return true
		}
	}
	if lhs.TokenType == Expression &&
		rhs.TokenType != Operator {
		return true
	}
	if rhs.TokenType == Expression &&
		lhs.TokenType != Operator {
		return true
	}
	if lhs.TokenType != rhs.TokenType {
		return false
	}
	if lhs.TokenType != Operator {
		return true
	}
	if lhs.Operator == rhs.Operator {
		return true
	}
	return false
}
