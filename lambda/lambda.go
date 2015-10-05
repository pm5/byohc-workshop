package lambda

import "fmt"

type Node interface {
	Sub(name string, value Node) Node
}

type Var struct {
	Name string
}

type Lambda struct {
	Argument string
	Body     Node
}

type App struct {
	Func     Node
	Argument Node
}

func NewVar(name string) Var {
	return Var{name}
}

func (self Var) String() string {
	return fmt.Sprintf("%s", self.Name)
}

func (self Var) Sub(name string, value Node) Node {
	if self.Name == name {
		return value
	} else {
		return self
	}
}

func NewLambda(arg string, body Node) Lambda {
	return Lambda{arg, body}
}

func (self Lambda) String() string {
	return fmt.Sprintf("(\\%s %s)", self.Argument, self.Body)
}

func (self Lambda) Sub(name string, value Node) Node {
	if self.Argument != name {
		self.Body = self.Body.Sub(name, value)
	}
	return self
}

func (self Lambda) Eval(arg Node) Node {
	return self.Body.Sub(self.Argument, arg)
}

func NewApp(f Node, arg Node) App {
	return App{f, arg}
}

func (self App) String() string {
	return fmt.Sprintf("%s %s", self.Func, self.Argument)
}

func (self App) Sub(name string, value Node) Node {
	self.Func = self.Func.Sub(name, value)
	self.Argument = self.Argument.Sub(name, value)
	return self
}

func (self App) Eval() Node {
	return self.Func.(Lambda).Eval(self.Argument)
}

func NewNode(nodes []interface{}) Node {
	switch nodes[0].(string) {
	case "var":
		return NewVar(nodes[1].(string))
	case "lam":
		return NewLambda(nodes[1].(string), NewNode(nodes[2].([]interface{})))
	case "app":
		return NewApp(NewNode(nodes[1].([]interface{})), NewNode(nodes[2].([]interface{})))
	}
	return nil
}

func Eval(node Node) Node {
	for {
		fmt.Printf("%s\n", node)
		switch node.(type) {
		default:
			return node
		case App:
			node = node.(App).Eval()
		}
	}
}
