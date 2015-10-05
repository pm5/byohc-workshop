package lambda

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	var expr []interface{}
	err := json.Unmarshal([]byte(`["app",["lam","true",["app",["lam","false",["app",["lam","and",["app",["app",["var","and"],["var","true"]],["var","true"]]],["lam","a",["lam","b",["app",["app",["var","a"],["var","b"]],["var","false"]]]]]],["lam","a",["lam","b",["var","b"]]]]],["lam","a",["lam","b",["var","a"]]]]`), &expr)
	if err != nil {
		t.Error(err)
	}
	n := NewNode(expr)
	if s := fmt.Sprintf("%s", n); s != `(\true (\false (\and and true true) (\a (\b a b false))) (\a (\b b))) (\a (\b a))` {
		t.Errorf("Correct nodes constructed: `%s`", s)
	}
}

func TestVarSub(t *testing.T) {
	foo := NewVar("foo")
	if nosub := foo.Sub("bar", NewVar("bar")); nosub.(Var).Name == "bar" {
		t.Errorf("Var Sub() should only substitute variables of the same name")
	}

	if sub := foo.Sub("foo", NewVar("bar")); sub.(Var).Name != "bar" {
		t.Errorf("Var Sub() should substitute variables of the same name.")
	}
}

func TestLambdaSub(t *testing.T) {
	lam := NewLambda("a", NewLambda("b", NewVar("c")))
	if r := lam.Sub("c", NewVar("foo")); r.(Lambda).Body.(Lambda).Body.(Var).Name != "foo" {
		t.Errorf("Incorrect substitute to Lambda `%s`", r)
	}
}

func TestLambdaSubScope(t *testing.T) {
	lam := NewLambda("b", NewVar("b"))
	if r := lam.Sub("b", NewVar("c")); r.(Lambda).Body.(Var).Name != "b" {
		t.Errorf("Incorrect substitute of Lambda `%s`", r)
	}
}

func TestLambdaEval(t *testing.T) {
	lam := NewLambda("a", NewVar("a"))
	if r := lam.Eval(NewVar("c")); r.(Var).Name != "c" {
		t.Errorf("Incorrect evaluation of Lambda `%s`", r)
	}
}

func TestAppEval(t *testing.T) {
	app := NewApp(NewLambda("a", NewVar("a")), NewVar("c"))
	if r := app.Eval(); r.(Var).Name != "c" {
		t.Errorf("Incorrect evaluation of App `%s`", r)
	}
}

func TestWeakNormalForm(t *testing.T) {
	var expr []interface{}
	err := json.Unmarshal([]byte(`["app",["lam","true",["app",["lam","false",["app",["lam","and",["app",["app",["var","and"],["var","true"]],["var","true"]]],["lam","a",["lam","b",["app",["app",["var","a"],["var","b"]],["var","false"]]]]]],["lam","a",["lam","b",["var","b"]]]]],["lam","a",["lam","b",["var","a"]]]]`), &expr)
	if err != nil {
		t.Error(err)
	}
	n := NewNode(expr)
	if r := WeakNormalForm(n); fmt.Sprintf("%s", r) != `(\a (\b a))` {
		t.Errorf("Incorrect evaluation `%s`", r)
	}
}

func TestNormalForm(t *testing.T) {
	var expr []interface{}
	err := json.Unmarshal([]byte(`["app",["lam","true",["app",["lam","false",["app",["lam","and",["app",["app",["var","and"],["var","true"]],["var","true"]]],["lam","a",["lam","b",["app",["app",["var","a"],["var","b"]],["var","false"]]]]]],["lam","a",["lam","b",["var","b"]]]]],["lam","a",["lam","b",["var","a"]]]]`), &expr)
	if err != nil {
		t.Error(err)
	}
	n := NewNode(expr)
	if r := NormalForm(n); fmt.Sprintf("%s", r) != `(\a (\b a))` {
		t.Errorf("Incorrect evaluation `%s`", r)
	}
}
