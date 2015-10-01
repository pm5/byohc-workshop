package lambda

import (
	"encoding/json"
	"fmt"
)

type Node interface{}

type Var struct {
	Name string
}

type Lambda struct {
	Argument Node
	Body     Node
}

type App struct {
	Func     Node
	Argument Node
}

func NewVar(nodes []interface{}) *Var {
	return &Var{nodes[1].(string)}
}

func (n Var) String() string {
	return fmt.Sprintf("[var %s]", n.Name)
}

func NewLambda(nodes []interface{}) *Lambda {
	return &Lambda{NewNode(nodes[1].([]interface{})), NewNode(nodes[2].([]interface{}))}
}

func (n Lambda) String() string {
	return fmt.Sprintf("[lam %s %s]", n.Argument, n.Body)
}

func NewApp(nodes []interface{}) *App {
	return &App{NewNode(nodes[1].([]interface{})), NewNode(nodes[2].([]interface{}))}
}

func (n App) String() string {
	return fmt.Sprintf("[app %s %s]", n.Func, n.Argument)
}

func NewNode(nodes []interface{}) Node {
	switch nodes[0].(string) {
	case "var":
		return NewVar(nodes)
	case "lam":
		return NewLambda(nodes)
	case "app":
		return NewApp(nodes)
	}
	return nil
}

func Test() {
	var expr []interface{}
	err := json.Unmarshal([]byte(`["app",["lam","true",["app",["lam","false",["app",["lam","and",["app",["app",["var","and"],["var","true"]],["var","true"]]],["lam","a",["lam","b",["app",["app",["var","a"],["var","b"]],["var","false"]]]]]],["lam","a",["lam","b",["var","b"]]]]],["lam","a",["lam","b",["var","a"]]]]`), &expr)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s", expr)
	n := NewNode(expr)
	fmt.Printf("%s", n)
}
