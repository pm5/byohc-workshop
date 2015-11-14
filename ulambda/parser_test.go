package ulambda

import (
	"strings"
	"testing"
)

func TestReplaceParenth(t *testing.T) {
	if s := replaceParenth(`(\a b c)(d e)`); s != `( \a b c ) ( d e )` {
		t.Errorf("Replace parentheses wrong: %s", s)
	}
}

func TestLexer(t *testing.T) {
	lex := lexer(`(\a b c)(d e)`)
	if strings.Join(lex, ".") != `(.\a.b.c.).(.d.e.)` {
		t.Errorf("Lexer wrong: %s", lex)
	}
}

func TestParseExprVar(t *testing.T) {
	n, err := ParseExpr(`a`)
	if err != nil {
		t.Errorf("Parse Var wrong: %s", err)
	}
	if n.Type != "var" || n.Name != "a" {
		t.Errorf("Parse Var wrong: %s", n)
	}
}

func TestParseExprLambda(t *testing.T) {
	n1, err := ParseExpr(`\a a`)
	if err != nil {
		t.Errorf("Parse Lambda wrong: %s", err)
	}
	if n1.Type != "lambda" || n1.Argument != "a" || n1.Children[0].Type != "var" || n1.Children[0].Name != "a" {
		t.Errorf("Parse Lambda wrong: %s", n1)
	}

	n2, err := ParseExpr(`\a b a`)
	if err != nil {
		t.Errorf("Parse Lambda wrong: %s", err)
	}
	if n2.Type != "lambda" || n2.Children[0].Type != "app" || n2.Children[0].Children[0].Name != "b" || n2.Children[0].Children[1].Name != "a" {
		t.Errorf("Parse Lambda wrong: %s", n2)
	}

	n3, err := ParseExpr(`\a \b a`)
	if err != nil {
		t.Errorf("Parse Lambda wrong: %s", err)
	}
	if n3.Type != "lambda" || n3.Children[0].Type != "lambda" || n3.Children[0].Argument != "b" || n3.Children[0].Children[0].Name != "a" {
		t.Errorf("Parse Lambda wrong: %s", n3)
	}
}

func TestParseExprApp(t *testing.T) {
	n1, err := ParseExpr(`a b`)
	if err != nil {
		t.Errorf("Parse App wrong: %s", err)
	}
	if n1.Type != "app" || n1.Children[0].Type != "var" || n1.Children[0].Name != "a" || n1.Children[1].Type != "var" || n1.Children[1].Name != "b" {
		t.Errorf("Parse App wrong: %s", n1)
	}

	n2, err := ParseExpr(`a b c`)
	if err != nil {
		t.Errorf("Parse App wrong: %s", err)
	}
	if n2.Children[1].Type != "app" || n2.Children[1].Children[0].Name != "b" || n2.Children[1].Children[1].Name != "c" {
		t.Errorf("Parse App wrong: %s", n2)
	}

	n3, err := ParseExpr(`(\a b a) foo`)
	if err != nil {
		t.Errorf("Parse Lambda wrong: %s", err)
	}
	if n3.Type != "app" || n3.Children[0].Type != "lambda" || n3.Children[0].Children[0].Type != "app" {
		t.Errorf("Parse Lambda wrong: %s", n3)
	}
}

func TestParseExprComplex(t *testing.T) {
	n1, err := ParseExpr(`(\y (\x y x)(\a \b b))(\z \t t z)`)
	if err != nil {
		t.Errorf("Parse complex expression wrong: %s", err)
	}
	if n1.Type != "app" || n1.Children[0].Type != "lambda" || n1.Children[1].Type != "lambda" || n1.Children[0].Children[0].Type != "app" {
		t.Errorf("Parse complex expression wrong: %s", n1)
	}
}

func TestParseExprBooleanAnd(t *testing.T) {
	n1, err := ParseExpr(`
	(\true
	(\false
	(\and
		(and true) true
	)(\a \b (a b) false)
	)(\a \b b)
	)(\a \b a)`)
	if err != nil {
		t.Errorf("Parse true and false wrong: %s", err)
	}
	t.Errorf("%s", n1)
}
