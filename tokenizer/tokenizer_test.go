package tokenizer

import "testing"

func TestSymbolicToken(t *testing.T) {
	to_tokenize := []rune("kappa 123")
	end, symbol := SymbolicToken(to_tokenize, 0)
	if end != len("kappa") {
		t.Error("Expected ", len("kappa"), " got ", end)
	}
	if symbol.TokenType != Symbol {
		t.Error("Expected ", Symbol, ", got ", symbol.TokenType)
	}
	if symbol.Symbolic != "kappa" {
		t.Error("Expected kappa got ", symbol.Symbolic)
	}
}

func TestOperatorToken(t *testing.T) {
	to_tokenize := []rune(">=5")
	end, operator := OperatorToken(to_tokenize, 0)
	if end != 2 {
		t.Error("Expected 2, got ", end)
	}
	if operator.TokenType != Operator {
		t.Error("Expected Operator (", Operator, "), got ", operator.TokenType)
	}
	if operator.Operator != []rune("≥")[0] {
		t.Error("Expected '≥', got ", string(operator.Operator))
	}
}

func TestNumericToken(t *testing.T) {
	to_tokenize := []rune("-1e-34ie5")
	end, numeric := NumericToken(to_tokenize, 0)
	if end != 7 {
		t.Error("Expected 7, got ", end)
	}
	if numeric.TokenType != ImagNumeric {
		t.Error("Expected ", ImagNumeric, ", got ", numeric.TokenType)
	}
	if numeric.Numeric != -1e-34 {
		t.Error("Expected '-1e-34', got ", numeric.Numeric)
	}
}

func TestSubexpressionToken(t *testing.T) {
	to_tokenize := []rune("(3 + (-1e-34i+4))")
	end, subexpression := SubexpressionToken(to_tokenize, 0)
	if end != 17 {
		t.Error("Expected 17, got ", end)
	}
	if subexpression.TokenType != Subexpression {
		t.Error("Expected ", Subexpression, ", got ", subexpression.TokenType)
	}
	if subexpression.Subexpression != "3 + (-1e-34i+4)" {
		t.Error("Expected \"3 + (-1e-34i+4)\", got \"", subexpression.Subexpression, "\"")
	}
}

func TestTokenizeString(t *testing.T) {
	to_tokenize := "∫1e-34ikappa ∂kappa (3 + (-1e-34i+4))"
	tokens := TokenizeString(to_tokenize)
	if len(tokens) != 6 {
		t.Error("Expected 6, got ", len(tokens))
	}
	if tokens[0].TokenType < Operator ||
		tokens[1].TokenType != ImagNumeric ||
		tokens[2].TokenType != Symbol ||
		tokens[3].TokenType < Operator ||
		tokens[4].TokenType != Symbol ||
		tokens[5].TokenType != Subexpression {
		t.Error("Expected {Operator, ImaginaryNumber, Symbol, Operator, Symbol, Subexpression}")
	}
}
