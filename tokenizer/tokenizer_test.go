package lexer

import "testing"

func TestSymbolToken(t *testing.T) {
	to_tokenize := "kappa 123"
	kind, end, symbol_func := SymbolToken(to_tokenize, 0)
	if kind != 1 {
		t.Error("Expected 0, got ", kind)
	}
	if end != len("kappa") {
		t.Error("Expected ", len("kappa"), " got ", end)
	}
	if symbol_func() != "kappa" {
		t.Error("Expected kappa got ", symbol_func())
	}
}
