package ulambda

import "fmt"

type Node interface {
	Sub(name string, value Node) Node
	FreeVar() []string
	String() string
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
	return self.Name
}

func (self Var) Sub(name string, value Node) Node {
	if self.Name == name {
		return value
	} else {
		return self
	}
}

func (self Var) FreeVar() []string {
	return []string{self.Name}
}

func NewLambda(arg string, body Node) Lambda {
	return Lambda{arg, body}
}

func (self Lambda) String() string {
	return fmt.Sprintf("(\\%s %s)", self.Argument, self.Body)
}

func (self Lambda) Sub(name string, value Node) Node {
	if self.Argument != name {
		freeVar := value.FreeVar()
		varMap := make(map[string]bool)
		for i := 0; i < len(freeVar); i++ {
			varMap[freeVar[i]] = true
		}
		if _, exists := varMap[self.Argument]; exists {
			freeVar = append(freeVar, self.Body.FreeVar()...)
			for i := 0; i < len(freeVar); i++ {
				varMap[freeVar[i]] = true
			}
			m := "m"
			for {
				if _, exists := varMap[m]; exists || m == name {
					m = m + "_"
				} else {
					break
				}
			}
			self.Body = self.Body.Sub(self.Argument, NewVar(m))
			self.Argument = m
		}
		self.Body = self.Body.Sub(name, value)
	}
	return self
}

func (self Lambda) FreeVar() []string {
	return append(self.Body.FreeVar(), self.Argument)
}

func NewApp(fn Node, arg Node) App {
	return App{fn, arg}
}

func (self App) String() string {
	return fmt.Sprintf("%s %s", self.Func, self.Argument)
}

func (self App) Sub(name string, value Node) Node {
	self.Func = self.Func.Sub(name, value)
	self.Argument = self.Argument.Sub(name, value)
	return self
}

func (self App) FreeVar() []string {
	return append(self.Func.FreeVar(), self.Argument.FreeVar()...)
}

func NewNode(expr *ASTNode) Node {
	switch expr.Type {
	case "var":
		return NewVar(expr.Name)
	case "lambda":
		return NewLambda(expr.Argument, NewNode(expr.Children[0]))
	case "app":
		return NewApp(NewNode(expr.Children[0]), NewNode(expr.Children[1]))
	}
	return nil
}

func WeakNormalForm(node Node) Node {
	var r Node
	switch node := node.(type) {
	default:
		return nil
	case Var:
		return node
	case Lambda:
		return node
	case App:
		r = WeakNormalForm(node.Func)
	}
	switch r := r.(type) {
	default:
		return node
	case Lambda:
		return WeakNormalForm(r.Body.Sub(r.Argument, node.(App).Argument))
	}
}

func NormalForm(node Node) Node {
	var r Node
	switch node := node.(type) {
	default:
		return nil
	case Var:
		return node
	case Lambda:
		return NewLambda(node.Argument, NormalForm(node.Body))
	case App:
		r = WeakNormalForm(node.Func)
	}
	switch r := r.(type) {
	default:
		return NewApp(NormalForm(node.(App).Func), NormalForm(node.(App).Argument))
	case Lambda:
		return NormalForm(r.Body.Sub(r.Argument, node.(App).Argument))
	}
}

func AlphaConv(node Node, from string, to string) Node {
	switch node := node.(type) {
	default:
		return nil
	case Var:
		if node.Name == from {
			node.Name = to
		}
		return node
	case Lambda:
		if node.Argument == from {
			node.Argument = to
		}
		node.Body = AlphaConv(node.Body, from, to)
		return node
	case App:
		node.Func = AlphaConv(node.Func, from, to)
		node.Argument = AlphaConv(node.Argument, from, to)
		return node
	}
}
