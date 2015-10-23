package ulambda

import (
	"strings"
	"testing"
)

func TestParseExprVar(t *testing.T) {
	n := ParseExpr(`a`)
	if n.(Var).Name != "a" {
		t.Errorf("Parse Var wrong: `%s`", n)
	}
}

func TestReplaceParenth(t *testing.T) {
	if s := replaceParenth(`(\a b c)(d e)`); s != `( \a b c ) ( d e )` {
		t.Errorf("Replace parentheses wrong: `%s`", s)
	}
}

func TestLexer(t *testing.T) {
	lex := lexer(`(\a b c)(d e)`)
	if strings.Join(lex, ".") != `(.\a.b.c.).(.d.e.)` {
		t.Errorf("Lexer wrong: %s", lex)
	}
}

func TestParseExprLambda(t *testing.T) {
	n1 := ParseExpr(`\a a`)
	if n1.(Lambda).Argument != "a" || n1.(Lambda).Body.(Var).Name != "a" {
		t.Errorf("Parse Lam wrong: `%s`", n1)
	}

	n2 := ParseExpr(`\a b a`)
	if n2.(Lambda).Body.(App).Func.(Var).Name != "b" ||
		n2.(Lambda).Body.(App).Argument.(Var).Name != "a" {
		t.Errorf("Parse Lam wrong: `%s`", n2)
	}
}

func TestParseExprApp(t *testing.T) {
	n1 := ParseExpr(`a b`)
	if n1.(App).Func.(Var).Name != "a" || n1.(App).Argument.(Var).Name != "b" {
		t.Errorf("Parse App wrong: `%s`", n1)
	}

	n2 := ParseExpr(`a b c`)
	if n2.(App).Func.(Var).Name != "a" || n2.(App).Argument.(App).Func.(Var).Name != "b" {
		t.Errorf("Parse App wrong: `%s`", n2)
	}
}

func TestParseExprComposit(t *testing.T) {
	n := ParseExpr(`(\y (\x y x)(+ 3 2))(\z (* z z))`)
	if n.(App).Func.(Lambda).Argument != "y" {
		t.Errorf("Parse composit expression wrong: `%s`", n)
	}
}

//func TestParseExprBooleanAnd(t *testing.T) {
//n := ParseExpr(`(\true
//(\false
//(\and
//(and true) true
//)(\a \b (a b) false)
//)(\a \b b)
//)(\a \b a)`)
//}
