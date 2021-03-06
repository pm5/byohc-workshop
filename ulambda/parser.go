package ulambda

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Error codes returned by failures to parse an expression.
var (
	ErrUnknown   = errors.New("ulambda: unknown error")
	ErrUnmatched = errors.New("ulambda: unmatched parentheses")
	ErrEmptyExpr = errors.New("ulambda: empty expression")
	ErrSyntax    = errors.New("ulambda: syntax error")
)

// ASTNode represents a node in the Abstract Syntax Tree.
type ASTNode struct {
	Type     string
	Name     string // for var type
	Argument string // for lambda type
	Children []*ASTNode
}

// String returns the string representation of the AST.
func (self ASTNode) String() string {
	switch self.Type {
	case "var":
		return fmt.Sprintf("ASTNode{var, %s}", self.Name)
	case "lambda":
		return fmt.Sprintf("ASTNode{lambda, %s, %s}", self.Argument, self.Children)
	case "app":
		return fmt.Sprintf("ASTNode{app, %s, %s}", self.Children[0], self.Children[1:])
	}
	return ""
}

// NewVarASTNode creates an AST node for a variable.
func NewVarASTNode(name string) *ASTNode {
	return &ASTNode{"var", name, "", nil}
}

// NewLambdaASTNode creates an AST node for a function.
func NewLambdaASTNode(arg string, body *ASTNode) *ASTNode {
	return &ASTNode{"lambda", "", arg, []*ASTNode{body}}
}

// NewAppASTNode creates an AST node for an application expression.
func NewAppASTNode(fun *ASTNode, arg *ASTNode) *ASTNode {
	return &ASTNode{"app", "", "", []*ASTNode{fun, arg}}
}

func replaceParenth(s string) string {
	pr := regexp.MustCompile(`([\(\)])`)
	sp := regexp.MustCompile(`  +`)
	ht := regexp.MustCompile(`^ +| +$`)
	return ht.ReplaceAllString(sp.ReplaceAllString(pr.ReplaceAllString(s, ` $1 `), ` `), ``)
}

func replaceWhitespace(s string) string {
	ws := regexp.MustCompile(`[\s\n]+`)
	return ws.ReplaceAllString(s, ` `)
}

func lexer(s string) []string {
	return strings.Split(replaceParenth(replaceWhitespace(s)), " ")
}

// ParseLex parses a sequence of string tokens into AST.
func ParseLex(lex []string) (*ASTNode, error) {
	if len(lex) == 0 {
		return nil, ErrEmptyExpr
	}
	if lex[0] == "(" {
		i := 1
		count := 1
		for ; i < len(lex) && count > 0; i++ {
			if lex[i] == "(" {
				count++
			} else if lex[i] == ")" {
				count--
			}
		}
		if count > 0 {
			return nil, ErrUnmatched
		} else {
			i--
		}
		body, err := ParseLex(lex[2:i])
		if err != nil {
			return nil, err
		}
		if i == len(lex)-1 {
			return NewLambdaASTNode(lex[1][1:], body), nil
		} else {
			arg, err := ParseLex(lex[i+1:])
			if err != nil {
				return nil, err
			}
			return NewAppASTNode(NewLambdaASTNode(lex[1][1:], body), arg), nil
		}
	}
	if len(lex) == 1 {
		return NewVarASTNode(lex[0]), nil
	} else if lex[0][0] == '\\' {
		body, err := ParseLex(lex[1:])
		if err != nil {
			return nil, err
		}
		return NewLambdaASTNode(lex[0][1:], body), nil
	} else {
		fun, err := ParseLex(lex[0:1])
		if err != nil {
			return nil, err
		}
		arg, err := ParseLex(lex[1:])
		if err != nil {
			return nil, err
		}
		return NewAppASTNode(fun, arg), nil
	}
	return nil, ErrUnknown
}

// ParseExpr parses a string of text into AST.
func ParseExpr(expr string) (*ASTNode, error) {
	return ParseLex(lexer(expr))
}
