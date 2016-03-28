package tokenizer

import "testing"

func TestSymbolicToken(t *testing.T) {
	to_tokenize := []rune("kappa 123")
	end, symbol_func := SymbolicToken(to_tokenize, 0)
	if end != len("kappa") {
		t.Error("Expected ", len("kappa"), " got ", end)
	}
	if len(symbol_func(GetToken)) != 1 {
		t.Error("Expected 1, got ", len(symbol_func(GetToken)))
	}
	if symbol_func(GetToken)[0]._type != 1 {
		t.Error("Expected 1, got ", symbol_func(GetToken)[0]._type)
	}
	if symbol_func(GetToken)[0].Symbolic != "kappa" {
		t.Error("Expected kappa got ", symbol_func(GetToken)[0].Symbolic)
	}
}

func TestOperatorToken(t *testing.T) {
	to_tokenize := []rune("≥5")
	end, symbol_func := OperatorToken(to_tokenize, 0)
	if end != 1 {
		t.Error("Expected 1, got ", end)
	}
	if len(symbol_func(GetToken)) != 1 {
		t.Error("Expected 1, got ", len(symbol_func(GetToken)))
	}
	if symbol_func(GetToken)[0]._type != 0 {
		t.Error("Expected 0, got ", symbol_func(GetToken)[0]._type)
	}
	if symbol_func(GetToken)[0].Operator != []rune("≥")[0] {
		t.Error("Expected '≥', got ", string(symbol_func(GetToken)[0].Operator))
	}
}

func TestNumericToken(t *testing.T) {
	to_tokenize := []rune("-1e-34ie5")
	end, numeric_func := NumericToken(to_tokenize, 0)
	if end != 7 {
		t.Error("Expected 6, got ", end)
	}
	if len(numeric_func(GetToken)) != 1 {
		t.Error("Expected 1, got ", len(numeric_func(GetToken)))
	}
	if numeric_func(GetToken)[0]._type != 3 {
		t.Error("Expected 3, got ", numeric_func(GetToken)[0]._type)
	}
	if numeric_func(GetToken)[0].Numeric != -1e-34 {
		t.Error("Expected '-1e-34', got ", numeric_func(GetToken)[0].Numeric)
	}
}

func TestSubexpressionToken(t *testing.T) {
	to_tokenize := []rune("(3 + (-1e-34i+4))")
	end, subexpression_func := SubexpressionToken(to_tokenize, 0)
	if end != 17 {
		t.Error("Expected 17, got ", end)
	}
	if subexpression_func(GetToken)[0]._type != Subexpression {
		t.Error("Expected ", Subexpression, ", got ", subexpression_func(GetToken)[0]._type)
	}
	if subexpression_func(GetToken)[0].Subexpression != "3 + (-1e-34i+4)" {
		t.Error("Expected \"3 + (-1e-34i+4)\", got \"", subexpression_func(GetToken)[0].Subexpression, "\"")
	}
	if len(subexpression_func(GetSubtokens)) != 3 {
		t.Error("Expected 3, got ", len(subexpression_func(GetSubtokens)))
	}
}

func TestTokenizeString(t *testing.T) {
	to_tokenize := "-1e-34i⁄kappa (3 + (-1e-34i+4))"
	generators := TokenizeString(to_tokenize)
	if len(generators) != 4 {
		t.Error("Expected 4, got ", len(generators))
	}
	if generators[0](GetToken)[0]._type != ImagNumeric ||
		generators[1](GetToken)[0]._type != Operator ||
		generators[2](GetToken)[0]._type != Symbol ||
		generators[3](GetToken)[0]._type != Subexpression {
		t.Error("Expected {ImaginaryNumber, Operator, Symbol, Subexpression}")
	}
}
