package tokenizer

import "testing"

func TestSymbolicToken(t *testing.T) {
	to_tokenize := "kappa 123"
	end, symbol_func := SymbolicToken(to_tokenize, 0)
	if end != len("kappa") {
		t.Error("Expected ", len("kappa"), " got ", end)
	}
	if symbol_func()._type != 1 {
		t.Error("Expected 1, got ", symbol_func()._type)
	}
	if symbol_func().Symbolic != "kappa" {
		t.Error("Expected kappa got ", symbol_func())
	}
}

func TestOperatorToken(t *testing.T) {
	to_tokenize := "≥5"
	end, symbol_func := OperatorToken(to_tokenize, 0)
	if end != 1 {
		t.Error("Expected 1, got ", end)
	}
	if symbol_func()._type != 0 {
		t.Error("Expected 0, got ", symbol_func()._type)
	}
	if symbol_func().Operator != rune("≥"[0]) {
		t.Error("Expected '+', got ", symbol_func().Operator)
	}
}

func TestNumericToken(t *testing.T) {
	to_tokenize := "-1e-34ie5"
	end, numeric_func := NumericToken(to_tokenize, 0)
	if end != 6 {
		t.Error("Expected 6, got ", end)
	}
	if numeric_func()._type != 3 {
		t.Error("Expected 3, got ", numeric_func()._type)
	}
	if numeric_func().Numeric != -1e-34 {
		t.Error("Expected '-1e-34', got ", numeric_func().Numeric)
	}
}
