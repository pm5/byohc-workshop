package ulambda

import "testing"

func TestNewNode(t *testing.T) {
	expr1, err := ParseExpr(`(\a (\b a))`)
	if err != nil {
		t.Error(err)
	}
	if NewNode(expr1).String() != `(\a (\b a))` {
		t.Errorf("Node construction wrong: %s", NewNode(expr1))
	}

	expr2, err := ParseExpr(`(\true (\false (\and and true true) (\a (\b a b false))) (\a (\b b))) (\a (\b a))`)
	if err != nil {
		t.Error(err)
	}
	if n := NewNode(expr2); n.String() != `(\true (\false (\and and true true) (\a (\b a b false))) (\a (\b b))) (\a (\b a))` {
		t.Errorf("Node construction wrong: `%s`", n)
	}
}

func TestVarSub(t *testing.T) {
	foo := NewVar("foo")

	if nobar := foo.Sub("bar", NewVar("bar")); nobar.(Var).Name == "bar" {
		t.Errorf("Var.Sub() substitute variables of wrong names")
	}
	if bar := foo.Sub("foo", NewVar("bar")); bar.(Var).Name != "bar" {
		t.Errorf("Var.Sub() does not substitute variables of correct name.")
	}
}

func TestLambdaSub(t *testing.T) {
	lam := NewLambda("a", NewLambda("b", NewVar("c")))

	if r := lam.Sub("c", NewVar("foo")); r.(Lambda).Body.(Lambda).Body.(Var).Name != "foo" {
		t.Errorf("Lambda.Sub() substitutes wrong: `%s`", r)
	}
}

func TestLambdaSubScope(t *testing.T) {
	lam := NewLambda("b", NewVar("b"))
	if r := lam.Sub("b", NewVar("c")); r.(Lambda).Body.(Var).Name != "b" {
		t.Errorf("Lambda.Sub() substitutes wrong: `%s`", r)
	}
}

//func TestAlphaConversion(t *testing.T) {
//expr, err := ParseExpr(`\x y y y x`)
//if err != nil {
//t.Error(err)
//}
//n := NewNode(expr)
//if alpha := n.AlphaConv(); alpha.String() != `\0 1 1 1 0` {
//t.Errorf("Alpha conversion wrong: %s", alpha)
//}
//}

func TestWeakNormalForm(t *testing.T) {
	expr, err := ParseExpr(`(\true (\false (\and (and true) true)(\a \b (a b) false))(\a \b b))(\a \b a)`)
	if err != nil {
		t.Error(err)
	}
	n := NewNode(expr)
	if r := WeakNormalForm(n); r.String() != `(\a (\b a))` {
		t.Errorf("WeakNormalForm wrong: `%s`", r)
	}
}

//func TestNormalForm(t *testing.T) {
//var expr []interface{}
//if err := json.Unmarshal([]byte(`["app",["lam","true",["app",["lam","false",["app",["lam","and",["app",["app",["var","and"],["var","true"]],["var","true"]]],["lam","a",["lam","b",["app",["app",["var","a"],["var","b"]],["var","false"]]]]]],["lam","a",["lam","b",["var","b"]]]]],["lam","a",["lam","b",["var","a"]]]]`), &expr); err != nil {
//t.Error(err)
//}
//n := NewNode(expr)
//if r := NormalForm(n); r.String() != `(\a (\b a))` {
//t.Errorf("NormalForm() simplifies wrong: `%s`", r)
//}
//}

//func TestCapture(t *testing.T) {
//var expr []interface{}
//if err := json.Unmarshal([]byte(`["app",["lam","true",["app",["lam","false",["app",["lam","not",["app",["var","not"],["var","true"]]],["lam","p",["lam","a",["lam","b",["app",["app",["var","p"],["var","b"]],["var","a"]]]]]]],["lam","a",["lam","b",["var","b"]]]]],["lam","a",["lam","b",["var","a"]]]]`), &expr); err != nil {
//t.Error(err)
//}
//r := NormalForm(NewNode(expr))
//if rl := r.(Lambda).Body.(Lambda); rl.Argument != rl.Body.(Var).Name {
//t.Errorf("NormalForm() must not capture: `%s`", r)
//}
//}

//func dec(node Node) Node {
//return NormalForm(NewApp(NewApp(node, NewLambda("x", NewVar("x"))), NewVar("_")))
//}

//func TestNaturalNumber(t *testing.T) {
//var expr []interface{}
//if err := json.Unmarshal([]byte(`["app",["lam","+1",["app",["lam","0",["app",["lam","1",["app",["lam","2",["app",["lam","3",["app",["lam","4",["app",["lam","5",["app",["lam","6",["app",["lam","7",["app",["lam","8",["app",["lam","9",["app",["lam","-1",["app",["var","-1"],["app",["var","-1"],["app",["var","-1"],["app",["var","+1"],["app",["var","+1"],["app",["var","+1"],["app",["var","-1"],["app",["var","-1"],["app",["var","-1"],["app",["var","-1"],["var","9"]]]]]]]]]]]],["lam","n",["app",["app",["var","n"],["lam","n-",["var","n-"]]],["var","0"]]]]],["app",["var","+1"],["var","8"]]]],["app",["var","+1"],["var","7"]]]],["app",["var","+1"],["var","6"]]]],["app",["var","+1"],["var","5"]]]],["app",["var","+1"],["var","4"]]]],["app",["var","+1"],["var","3"]]]],["app",["var","+1"],["var","2"]]]],["app",["var","+1"],["var","1"]]]],["app",["var","+1"],["var","0"]]]],["lam","s",["lam","z",["var","z"]]]]],["lam","n",["lam","s",["lam","z",["app",["var","s"],["var","n"]]]]]]`), &expr); err != nil {
//t.Error(err)
//}
//r := NormalForm(NewNode(expr))
//zero := dec(dec(dec(dec(dec(r)))))
//if zero := zero.(Lambda).Body.(Lambda); zero.Argument != zero.Body.(Var).Name {
//t.Errorf("Natural number evaluates wrong: `%s`\n", r)
//}
//}

/*
func TestInfiniteSequencel(t *testing.T) {
	var expr []interface{}
	err := json.Unmarshal([]byte(`["app",["lam","+1",["app",["lam","0",["app",["lam","1",["app",["lam","2",["app",["lam","3",["app",["lam","4",["app",["lam","5",["app",["lam","6",["app",["lam","7",["app",["lam","8",["app",["lam","9",["app",["lam","nil",["app",["lam","cons",["app",["lam","Y",["app",["lam","take",["app",["lam","map",["app",["lam","0-1-2-",["app",["app",["var","take"],["var","5"]],["var","0-1-2-"]]],["app",["var","Y"],["lam","0-1-2-",["app",["app",["var","cons"],["var","0"]],["app",["app",["var","map"],["var","+1"]],["var","0-1-2-"]]]]]]],["lam","f",["app",["var","Y"],["lam","go",["lam","ls",["app",["app",["var","ls"],["lam","a",["lam","as",["app",["app",["var","cons"],["app",["var","f"],["var","a"]]],["app",["var","go"],["var","as"]]]]]],["var","nil"]]]]]]]],["app",["var","Y"],["lam","take",["lam","n",["lam","ls",["app",["app",["var","n"],["lam","n-",["app",["app",["var","ls"],["lam","a",["lam","as",["app",["app",["var","cons"],["var","a"]],["app",["app",["var","take"],["var","n-"]],["var","as"]]]]]],["var","nil"]]]],["var","nil"]]]]]]]],["lam","f",["app",["lam","x",["app",["var","f"],["app",["var","x"],["var","x"]]]],["lam","x",["app",["var","f"],["app",["var","x"],["var","x"]]]]]]]],["lam","a",["lam","as",["lam","is-cons",["lam","is-nil",["app",["app",["var","is-cons"],["var","a"]],["var","as"]]]]]]]],["lam","is-cons",["lam","is-nil",["var","is-nil"]]]]],["app",["var","+1"],["var","8"]]]],["app",["var","+1"],["var","7"]]]],["app",["var","+1"],["var","6"]]]],["app",["var","+1"],["var","5"]]]],["app",["var","+1"],["var","4"]]]],["app",["var","+1"],["var","3"]]]],["app",["var","+1"],["var","2"]]]],["app",["var","+1"],["var","1"]]]],["app",["var","+1"],["var","0"]]]],["lam","s",["lam","z",["var","z"]]]]],["lam","n",["lam","s",["lam","z",["app",["var","s"],["var","n"]]]]]]`), &expr)
	if err != nil {
		t.Error(err)
	}
	_ = NormalForm(NewNode(expr))
}
*/
