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
	return fmt.Sprintf("%s", self.Name)
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

func (self App) FreeVar() []string {
	return append(self.Func.FreeVar(), self.Argument.FreeVar()...)
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
